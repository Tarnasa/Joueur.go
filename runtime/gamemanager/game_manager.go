package gamemanager

import (
	"errors"
	"joueur/base"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"

	"net/url"
)

// GameManager is a factory and manager for the game.
// It sits between the client, AI, and Game to facilitate interactions
// between all of them.
type GameManager struct {
	GameNamespace games.GameNamespace
	Game          base.DeltaMergeableGame
	AI            base.AI
	Player        base.Player

	aiImpl          *base.AIImpl
	gameObjects     map[string]base.DeltaMergeableGameObject
	serverConstants client.ServerConstants
	deltaMerge      base.DeltaMerge
}

// New creates a new instance of a GameManager for a given namespace.
// This should be the only factory/way to create GameManagers.
func New(
	gameNamespace games.GameNamespace,
	aiSettings string,
	serverConstants client.ServerConstants,
) *GameManager {
	gameManager := GameManager{}

	gameManager.GameNamespace = gameNamespace
	gameManager.Game = gameNamespace.CreateGame()
	gameManager.AI, gameManager.aiImpl = gameNamespace.CreateAI()
	gameManager.gameObjects = make(map[string]base.DeltaMergeableGameObject)

	if serverConstants.DeltaListLengthKey == "" || serverConstants.DeltaRemoved == "" {
		errorhandler.HandleError(
			errorhandler.DeltaMergeFailure,
			errors.New("invalid server constants, delta array length key is: '"+
				serverConstants.DeltaListLengthKey+
				"' and delta removed is: '"+
				serverConstants.DeltaRemoved+"'."),
		)
	}

	gameManager.serverConstants = serverConstants
	gameManager.deltaMerge = gameManager.GameNamespace.CreateDeltaMerge(&base.DeltaMergeImpl{
		Game:              gameManager.Game,
		DeltaRemovedValue: gameManager.serverConstants.DeltaRemoved,
		DeltaLengthKey:    gameManager.serverConstants.DeltaListLengthKey,
	})

	client.EventDeltaHandler = func(delta map[string]interface{}) {
		gameManager.applyDeltaState(delta)
	}

	client.EventOrderHandler = func(order client.EventOrderData) {
		gameManager.handleOrder(order)
	}

	base.RunOnServerCallback = func(
		caller base.GameObject,
		functionName string,
		args map[string]interface{},
	) interface{} {
		return gameManager.RunOnServer(caller, functionName, args)
	}

	return &gameManager
}

// parseAISettings parses an AI Settings string and sets the AI's setting
// map accordingly.
func (gameManager *GameManager) parseAISettings(aiSettings string) {
	settings := make(map[string]([]string))
	parsedSettings, parseErr := url.ParseQuery(aiSettings)
	if parseErr != nil {
		errorhandler.HandleError(
			errorhandler.InvalidArgs,
			parseErr,
			"Error occurred while parsing AI Settings query string. Ensure the format is correct",
		)
	}

	for key, value := range parsedSettings {
		settings[key] = value
	}

	// hack-y, look into a cleaner way?
	base.AISettings = settings
}

// Start should be invoked when the ame starts and our playerID is known
func (gameManager *GameManager) Start(playerID string) {
	playerGameObject, foundGameObjectWithID := gameManager.Game.GetGameObject(playerID)
	if !foundGameObjectWithID {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("could not find GameObject with id #"+playerID+" for AI's Player"),
		)
	}
	player, isPlayer := playerGameObject.(base.Player)
	if !isPlayer {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Game Object #"+playerID+" is not a Player when it's supposed to be our player"),
		)
	}
	gameManager.Player = player
	base.InjectIntoAI(gameManager.aiImpl, gameManager.Game, gameManager.Player)

	gameManager.AI.GameUpdated()
	gameManager.AI.Start()

	// game should now be started
}

// RunOnServer should be invoked by GameObjects when they want some function
// and args to be ran on the game server on their behalf.
func (gameManager *GameManager) RunOnServer(
	caller base.GameObject,
	functionName string,
	args map[string]interface{},
) interface{} {
	serializedArgs := gameManager.serialize(args)
	serializedArgsMap, isMap := serializedArgs.(map[string]interface{})
	if !isMap {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Serialized args for "+functionName+" and did not get expected map"),
		)
	}
	client.SendEventRun(client.EventRunData{
		Caller:       client.GameObjectReference{Id: caller.ID()},
		FunctionName: functionName,
		Args:         serializedArgsMap,
	})

	returned := client.WaitForEventRan()

	return gameManager.deSerialize(returned)
}

// handlerOrder is automatically invoked when an  event comes from the server.
func (gameManager *GameManager) handleOrder(order client.EventOrderData) {
	args := gameManager.deSerialize(order.Args)
	argsList, isList := args.([]interface{})
	if !isList {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot handle order "+order.Name+" because the args are not a slice"),
		)
	}
	returned, err := gameManager.GameNamespace.OrderAI(gameManager.AI, order.Name, argsList)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			err,
			"GameManager could not handle order "+order.Name,
		)
	}

	// now that we've finished the order, tell the server
	client.SendEventFinished(client.EventFinishedData{
		OrderIndex: order.Index,
		Returned:   returned,
	})
}

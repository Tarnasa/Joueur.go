package gamemanager

import (
	"errors"
	"fmt"
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
	GameNamespace   games.GameNamespace
	ServerConstants client.ServerConstants
	Game            base.DeltaMergeableGame
	AI              base.AI
	Player          base.Player

	started     bool
	backOrders  []client.EventOrderData
	aiImpl      *base.AIImpl
	gameObjects map[string]base.DeltaMergeableGameObject
	deltaMerge  base.DeltaMerge
}

// New creates a new instance of a GameManager for a given namespace.
// This should be the only factory/way to create GameManagers.
func New(gameNamespace games.GameNamespace, aiSettings string) *GameManager {
	gameManager := GameManager{}

	gameManager.GameNamespace = gameNamespace
	gameManager.Game = gameNamespace.CreateGame()
	gameManager.AI, gameManager.aiImpl = gameNamespace.CreateAI()
	gameManager.aiImpl.Game = gameManager.Game
	gameManager.gameObjects = make(map[string]base.DeltaMergeableGameObject)
	gameManager.deltaMerge = gameNamespace.CreateDeltaMerge(base.DeltaMergeImpl{
		Game:              gameManager.Game,
		DeltaRemovedValue: gameManager.ServerConstants.DeltaRemoved,
		DeltaLengthKey:    gameManager.ServerConstants.DeltaListLengthKey,
	})

	gameManager.started = false // normal default but we want to be clear
	gameManager.backOrders = make([]client.EventOrderData, 0)

	client.RegisterEventDeltaHandler(func(delta map[string]interface{}) {
		fmt.Println("registered delta thing do a thing", delta)
		gameManager.applyDeltaState(delta)
	})

	client.EventOverHandler = func(order client.EventOrderData) {
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

func (gameManager GameManager) parseAISettings(aiSettings string) {
	settings := make(map[string]([]string))
	parsedSettings, parseErr := url.ParseQuery(aiSettings)
	if parseErr != nil {
		errorhandler.HandleError(
			errorhandler.InvalidArgs,
			parseErr,
			"Error occured while parsing AI Settings query string. Ensure the format is correct",
		)
	}

	for key, value := range parsedSettings {
		settings[key] = value
	}

	// hack-y, look into a cleaner way?
	base.AISettings = settings
}

// Start should be invoked when the ame starts and our playerID is known
func (gameManager GameManager) Start(playerID string) {
	gameManager.started = true
	// TODO: set player in ai by this ID
	if playerGameObject, ok := gameManager.Game.GetGameObject(playerID); ok {
		player, isPlayer := playerGameObject.(base.Player)
		if !isPlayer {
			errorhandler.HandleError(
				errorhandler.ReflectionFailed,
				errors.New("Game Object #"+playerID+" is not a Player when it's supposed to be our player"),
			)
		}
		gameManager.aiImpl.Player = player
		gameManager.Player = player
	} else {
		// handle error
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not find Player with id #"+playerID),
		)
	}

	gameManager.AI.GameUpdated()
	// do back orders
	for _, order := range gameManager.backOrders {
		gameManager.handleOrder(order)
	}

	// game should now be started
}

// RunOnServer should be invoked by GameObjects when they want some function
// and args to be ran on the game server on their behalf.
func (gameManager GameManager) RunOnServer(
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
func (gameManager GameManager) handleOrder(order client.EventOrderData) {
	if !gameManager.started {
		gameManager.backOrders = append(gameManager.backOrders, order)
		return
	}

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

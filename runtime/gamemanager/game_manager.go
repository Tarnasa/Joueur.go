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

type GameManager struct {
	GameNamespace   games.GameNamespace
	ServerConstants client.ServerConstants
	Game            base.BaseGame
	AI              base.BaseAI
	Player          base.BasePlayer

	started         bool
	backOrders      []client.EventOrderData
	gameImpl        *base.BaseDeltaMergeableImpl
	gameObjectImpls map[string]*base.BaseDeltaMergeableImpl
	AIImpl          *base.BaseAIImpl
}

func New(gameNamespace games.GameNamespace, aiSettings string) *GameManager {
	gameManager := GameManager{}

	gameManager.GameNamespace = gameNamespace
	gameManager.Game, gameManager.gameImpl = gameNamespace.CreateGame()
	gameManager.AI, gameManager.AIImpl = gameNamespace.CreateAI()
	gameManager.AIImpl.Game = gameManager.Game

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
		caller base.BaseGameObject,
		functionName string,
		args map[string]interface{},
	) interface{} {
		return gameManager.RunOnServer(caller, functionName, args)
	}

	return &gameManager
}

func (this GameManager) parseAISettings(aiSettings string) {
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

	// hack-y, dont' like
	base.AISettings = settings
}

func (this GameManager) Start(playerID string) {
	this.started = true
	// TODO: set player in ai by this ID
	if playerGameObject, ok := this.Game.GetGameObject(playerID); ok {
		player, isPlayer := playerGameObject.(base.BasePlayer)
		if !isPlayer {
			errorhandler.HandleError(
				errorhandler.ReflectionFailed,
				errors.New("Game Object #"+playerID+" is not a Player when it's supposed to be our player"),
			)
		}
		this.AIImpl.Player = player
		this.Player = player
	} else {
		// handle error
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not find Player with id #"+playerID),
		)
	}

	this.AI.GameUpdated()
	// do back orders
	for _, order := range this.backOrders {
		this.handleOrder(order)
	}

	// game should now be started
}

func (this GameManager) RunOnServer(
	caller base.BaseGameObject,
	functionName string,
	args map[string]interface{},
) interface{} {
	serializedArgs := this.serialize(args)
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

	return this.deSerialize(returned)
}

func (this GameManager) handleOrder(order client.EventOrderData) {
	if !this.started {
		this.backOrders = append(this.backOrders, order)
		return
	}

	args := this.deSerialize(order.Args)
	argsList, isList := args.([]interface{})
	if !isList {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot handle order "+order.Name+" because the args are not a slice"),
		)
	}
	returned, err := this.GameNamespace.OrderAI(this.AI, order.Name, argsList)
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

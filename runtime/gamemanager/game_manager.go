package gamemanager

import (
	"fmt"
	"joueur/base"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"

	"net/url"
)

type GameManager struct {
	GameNamespace    games.GameNamespace
	ServerConstants  client.ServerConstants
	Game             base.BaseGame
	AI               base.BaseAI

	started bool
	backOrders      []client.EventOrderData
}

func New(gameNamespace *GameNamespace, aiSettings string) *GameManager {
	gameManager := GameManager{
		GameNamespace: gameNamespace,
		Game: (*gameManager.GameNamespace).CreateGame(),
		AI: (*gameManager.GameNamespace).CreateAI(),
	}

	gameManager.AI.Game = gameManager.Game

	client.RegisterEventDeltaHandler(func(delta map[string]interface{}) {
		fmt.Println("registered delta thing do a thing", delta)
		gameManager.applyDeltaState(delta)
	})

	client.EventOverHandler = func (order client.EventOrderData) {
		gameManager.handleOrder(order)
	}

	base.RunOnServerCallback = func GameManagerRunOnServer(
		caller *base.BaseGameObject,
		functionName string,
		args map[string]interface{},
	) interface{} {
		return gameManager.RunOnServer(caller, functionName, args)
	}

	return gameManager
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
	if playerGameObject, ok := (*this.Game).GameObjects[playerID]; ok {
		this.AI.Player = playerGameObject.(base.BasePlayer)
	} else {
		// handle error
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not find Player with id #" + playerID)
		)
	}

	this.AI.GameUpdated()
	// do back orders
	for _, order := range this.backOrders {
		this.handleOrder(order)
	}

	// game should now be started
}

func (this GameManager) RunOnServer(functionName string, args map[string]interface{}) interface{} {
	client.SendEventRun(client.EventRunData{
		FunctionName: functionName,
		Args: this.serialize(args),
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
	returned, err := this.GameNamespace.OrderAI(this.AI, order.Name, args)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			err,
			"GameManager could not handle order " + order.Name,
		)
	}

	// now that we've finished the order, tell the server
	client.SendEventFinished(client.EventFinishedData{
		OrderIndex: order.Index,
		Returned: returned,
	})
}

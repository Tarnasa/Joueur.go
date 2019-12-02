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
	ServerConstants client.ServerConstants
	GameNamespace   *games.GameNamespace
	InterfaceAI     *base.InterfaceAI
	Game            *base.BaseGame
	AI              *base.BaseAI
}

func New(gameManager *GameManager, aiSettings string) *GameManager {
	game := (*gameManager.GameNamespace).CreateGame()
	gameManager.Game = game

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

	ai := (*gameManager.GameNamespace).CreateAI()

	gameManager.AI = ai
	// hack-y, dont' like
	base.AISettings = settings

	(*ai).Game = game

	client.RegisterEventDeltaHandler(func(delta map[string]interface{}) {
		fmt.Println("registered delta thing do a thing", delta)
		gameManager.applyDeltaState(delta)
	})

	return gameManager
}

func (this GameManager) Start(playerID string) {
	// TODO: set player in ai by this ID
	if playerGameObject, ok := (*this.Game).GameObjects[playerID]; ok {
		(*this.AI).Player = playerGameObject
	} else {
		// handle error
	}
}

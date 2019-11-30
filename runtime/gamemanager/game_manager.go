package gamemanager

import (
	"joueur/base"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"

	"errors"
	"reflect"
)

type GameManager struct {
	ServerConstants client.ServerConstants
	GameNamespace *games.GameNamespace
	InterfaceAI *base.InterfaceAI
	ReflectAI *reflect.Value

	reflectGame *reflect.Value
}

func New(gameManager *GameManager) *GameManager {
	reflectGame := reflect.New((*gameManager.GameNamespace).GameType)
	gameManager.reflectGame = &reflectGame

	if !reflectGame.IsValid() {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not create Game instance for "+(*(gameManager.GameNamespace)).Name),
		)
	}

	return gameManager
}

func (gameManager GameManager) Start(playerID string) {
	// TODO: set player in ai by this ID
}

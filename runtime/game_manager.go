package runtime

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

func (gameManager GameManager) init() {
	reflectGame := reflect.New((*gameManager.GameNamespace).GameType)
	gameManager.reflectGame = &reflectGame

	if !reflectGame.IsValid() {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not create Game instance for "+(*(gameManager.GameNamespace)).Name),
		)
	}
}

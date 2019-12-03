package base

import (
	"errors"
	"joueur/runtime/errorhandler"
)

var RunOnServerCallback *func(BaseGameObject, string, map[string]interface{}) interface{}

type BaseGameObject interface {
	/*
		// A unique ID (unique to the game instance) of the game object.
		// Will never change, and IDs are never re-used.
		Id() string

		GameObjectName() string

		Game() BaseGame
	*/
}

type BaseGameObjectImpl struct {
	BaseDeltaMergeableImpl

	game BaseGame
}

func (this BaseGameObjectImpl) RunOnServer(functionName string, args map[string]interface{}) interface{} {
	if RunOnServerCallback == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot invoke function "+functionName+"on server without callback set!"),
		)
	}

	return (*RunOnServerCallback)(&this, functionName, args)
}

/*
func (this BaseGameObjectImpl) Id() string {
	return this.Data["id"].(string)
}

func (this BaseGameObjectImpl) GameObjectName() string {
	return this.Data["gameObjectName"].(string)
}

func (this BaseGameObjectImpl) Game() BaseGame {
	return this.Game
}
*/

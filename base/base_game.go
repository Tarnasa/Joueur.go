package base

import (
	"errors"
	"joueur/runtime/errorhandler"
)

type BaseGameImpl struct {
	BaseDeltaMergeableImpl
}

type BaseGame interface {
	GetGameObject(string) (BaseGameObject, bool)
}

func (this BaseGameImpl) GetGameObject(id string) (BaseGameObject, bool) {
	gameObjectsRaw, found := (this.InternalDataMap)["gameObjects"]
	gameObjectsMap, isMap := gameObjectsRaw.(map[string]BaseGameObject)
	if !found || gameObjectsMap == nil || !isMap {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("gameObjects not found as a map of strings to BaseGameObjects in Game!"),
		)
	}

	gameObject, found := gameObjectsMap[id]
	return gameObject, found
}

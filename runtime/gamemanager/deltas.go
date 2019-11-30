package gamemanager

import (
	"errors"
	"joueur/runtime/errorhandler"
	"reflect"
)

func (gameManager GameManager) applyDeltaState(delta map[string]interface{}) {
	gameObjects, ok := delta["gameObjects"]
	if ok {
		gameManager.initGameObjects(gameObjects.(map[string](map[string]interface{})))
	}
}

func (gameManager GameManager) initGameObjects(gameObjectDeltas map[string](map[string]interface{})) {
	reflectedGameObjects := (*gameManager.reflectGame).FieldByName("GameObjects")
	for key, gameObjectDelta := range gameObjectDeltas {
		rawGameObjectID, idOK := gameObjectDelta["id"]
		if !idOK {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot find id for "+key),
			)
		}

		id := rawGameObjectID.(string)
		if id != key {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot create new game object for mismatched ids: " + key + " and " + id),
			)
		}

		rawGameObjectName, gameObjectNameOK := gameObjectDelta["gameObjectName"]
		gameObjectName := rawGameObjectName.(string)
		if gameObjectName == "" || !gameObjectNameOK {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot get game object name from game object #" + id),
			)
		}

		gameObjectType, gameObjectTypeOK := (*gameManager.GameNamespace).GameObjectTypes[gameObjectName]
		if !gameObjectTypeOK {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot get type for gameObjectName " + gameObjectName),
			)
		}

		reflectedGameObject := reflect.New(gameObjectType)
		if !reflectedGameObject.IsValid() {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Could not create valid instance for " + gameObjectName + " #" + id),
			)
		}
		reflectedGameObjects.SetMapIndex(reflect.ValueOf(id), reflectedGameObject.Addr())
	}
}

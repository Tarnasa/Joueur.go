package gamemanager

import (
	"errors"
	"fmt"
	"joueur/runtime/errorhandler"
	"reflect"
)

func (this GameManager) applyDeltaState(delta map[string]interface{}) {
	fmt.Println(">>len of base delta", len(delta))
	gameObjects, ok := delta["gameObjects"]
	for key, value := range delta {
		fmt.Println(">>apply delta, what's in it?", key, value)
	}
	if ok {
		fmt.Println(">>going to attempt to merge gameObjects", gameObjects)
		this.initGameObjects(gameObjects.(map[string]interface{}))
	}

	// TODO: now delta merge
}

func (gameManager GameManager) initGameObjects(gameObjectDeltas map[string]interface{}) {
	reflectedGameObjects := (*gameManager.reflectGame).Elem().FieldByName("GameObjects")
	for key, gameObjectDelta := range gameObjectDeltas {
		fmt.Println("Attemping to initialize game object", key, gameObjectDelta)
		godAsMap := gameObjectDelta.(map[string]interface{})
		rawGameObjectID, idOK := godAsMap["id"]
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
				errors.New("Cannot create new game object for mismatched ids: "+key+" and "+id),
			)
		}

		rawGameObjectName, gameObjectNameOK := godAsMap["gameObjectName"]
		gameObjectName := rawGameObjectName.(string)
		if gameObjectName == "" || !gameObjectNameOK {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot get game object name from game object #"+id),
			)
		}

		gameObjectType, gameObjectTypeOK := (*gameManager.GameNamespace).GameObjectTypes[gameObjectName]
		if !gameObjectTypeOK {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot get type for gameObjectName "+gameObjectName),
			)
		}

		reflectedGameObject := reflect.New(gameObjectType)
		if !reflectedGameObject.IsValid() {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Could not create valid instance for "+gameObjectName+" #"+id),
			)
		}
		fmt.Println(">>Oh boy this shit is gonna be bananas", reflectedGameObjects.Kind(), reflectedGameObjects.Type())
		reflectedGameObjects.SetMapIndex(reflect.ValueOf(id), reflectedGameObject)
		fmt.Println(">>hohoho")
		reflectedGameObject.FieldByName("Game").Set(*gameManager.reflectGame)
	}
}

func (this GameManager) isDeltaPrimitive(delta interface{}) bool {
	if delta == this.ServerConstants.DeltaRemoved {
		return false
	}

	_, isBool := delta.(bool)
	_, isInt := delta.(int64)
	_, isFloat := delta.(float64)
	_, isString := delta.(string)

	return isBool || isInt || isFloat || isString
}

func (this GameManager) deltaMerge(delta interface{}) interface{} {
	if isDeltaPrimitive(delta) {
		return delta
	}
}

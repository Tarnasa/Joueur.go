package gamemanager

import (
	"errors"
	"fmt"
	"joueur/runtime/errorhandler"
	"strconv"
)

// applyDeltaState take the root (game) state and delta updates all the
// structs within the game it is managing.
func (gameManager *GameManager) applyDeltaState(delta map[string]interface{}) {
	fmt.Println(">> Merging delta start", delta)
	gameObjectsRaw, gameObjectsExist := delta["gameObjects"]
	gameObjectsDeltaRaw, gameObjectsAreMapRaw := gameObjectsRaw.(map[string]interface{})
	if !gameObjectsAreMapRaw {
		errorhandler.HandleError(
			errorhandler.DeltaMergeFailure,
			errors.New("cannot merge delta when 'gameObjects' property is not a map: "+fmt.Sprintf("%v", gameObjectsRaw)),
		)
	}

	gameObjectsDeltas := make(map[string]map[string]interface{})
	for id, rawGameObjectDelta := range gameObjectsDeltaRaw {
		gameObjectsDeltas[id] = gameManager.deltaMerge.ToDeltaMap(rawGameObjectDelta)
	}

	if gameObjectsExist {
		fmt.Println(">> init game objects", gameObjectsDeltas)
		gameManager.initGameObjects(gameObjectsDeltas)
	}
	fmt.Println(">> game objects should be init'd")

	// now all new game objects should be initialize so we can delta merge as normal
	if gameObjectsExist {
		for id, gameObjectDelta := range gameObjectsDeltas {
			gameObject, gameObjectExists := gameManager.gameObjects[id]
			if !gameObjectExists {
				errorhandler.HandleError(
					errorhandler.DeltaMergeFailure,
					errors.New("cannot merge delta state of game object #"+id+" with no game object for given id"),
				)
			}
			fmt.Println(">> delta merge game object", gameObjectDelta)
			for gameObjectAttribute, gameObjectAttributeDelta := range gameObjectDelta {
				fmt.Println("->-> DeltaMerging", gameObjectAttribute, gameObjectAttributeDelta)
				gameObject.DeltaMerge(gameManager.deltaMerge, gameObjectAttribute, gameObjectAttributeDelta)
			}
		}
	}
	// now all game objects should be delta merged, only thing remaining is the game itself's delta state
	for gameAttribute, gameAttributeDelta := range delta {
		if gameAttribute == "gameObjects" {
			continue // we already updated gameObject above
		}

		fmt.Println(">> delta merge into game", gameAttribute, gameAttributeDelta)
		gameManager.Game.DeltaMerge(gameManager.deltaMerge, gameAttribute, gameAttributeDelta)
	}

	fmt.Println("++++++++++++++++++ DONE DELTA MERGE? +++++++++++++")
}

func (gameManager *GameManager) initGameObjects(gameObjectsDeltas map[string]map[string]interface{}) {
	for id, gameObjectDelta := range gameObjectsDeltas {
		_, gameObjectAlreadyExists := gameManager.gameObjects[id]
		if gameObjectAlreadyExists {
			continue
		}

		fmt.Println("!!!yo yo yo", gameObjectDelta)
		rawGameObjectName, rawNameOk := gameObjectDelta["gameObjectName"]
		gameObjectName, nameIsString := rawGameObjectName.(string)
		if !rawNameOk || !nameIsString || gameObjectName == "" {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("field 'gameObjectName' not a string on new game object #"+id),
			)
		}

		newGameObject, creationError := gameManager.GameNamespace.CreateGameObject(gameObjectName)
		if creationError != nil {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("creation error on new game object "+gameObjectName+" #"+id),
			)
		}

		fmt.Println(">>NEW NEW NEW<<", newGameObject)
		gameManager.gameObjects[id] = newGameObject
		gameManager.Game.AddGameObject(id, newGameObject)
	}
}

func (gameManager *GameManager) isDeltaPrimitive(delta interface{}) bool {
	if delta == gameManager.ServerConstants.DeltaRemoved {
		return false
	}

	_, isBool := delta.(bool)
	_, isInt := delta.(int64)
	_, isFloat := delta.(float64)
	_, isString := delta.(string)

	return isBool || isInt || isFloat || isString
}

func (gameManager *GameManager) mergeDelta(state interface{}, delta interface{}) interface{} {
	if gameManager.isDeltaPrimitive(delta) {
		return delta
	}

	gameObject := gameManager.getIfGameObjectReference(delta)
	if gameObject != nil {
		return gameObject
	}

	deltaMap, isDeltaMap := delta.(map[string]interface{})

	if !isDeltaMap {
		errorhandler.HandleError(
			errorhandler.DeltaMergeFailure,
			errors.New("cannot merge non primitive and non map delta"),
		)
	}
	deltaLengthValue, hasDeltaLength := deltaMap[gameManager.ServerConstants.DeltaListLengthKey]

	if hasDeltaLength {
		// Then gameManager part in the state is an array
		deltaLength, deltaLenggameManagerInt := deltaLengthValue.(int64)
		// We don't want to copy gameManager key/value over to the state, it was just to signify the delta is an array
		delete(deltaMap, gameManager.ServerConstants.DeltaListLengthKey)

		if !deltaLenggameManagerInt {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("delta list length key present without being a number"),
			)
		}

		if state == nil {
			state = make([]interface{}, deltaLength)
		}

		stateList, isList := state.([]interface{})
		if !isList {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("delta merging is not a slice! Cannot resize"),
			)
		}
		stateList = stateList[:deltaLength]
	}

	if state == nil {
		state = make(map[string]interface{})
	}

	stateList, isList := state.([]interface{})
	stateMap, isMap := state.(map[string]interface{})

	for key, deltaValue := range deltaMap {
		keyAsIndex := 0
		if isList {
			keyAsIndex, err := strconv.Atoi(key)
			if err != nil || len(stateList) >= keyAsIndex || keyAsIndex < 0 {
				if err == nil {
					err = errors.New("key index " + key + "out out of range")
				}
				errorhandler.HandleError(
					errorhandler.DeltaMergeFailure,
					err,
					"Cannot merge into list with key index "+key,
				)
			}
		}

		if deltaValue == gameManager.ServerConstants.DeltaRemoved && !isList {
			delete(stateMap, key)
		} else {
			if isList {
				stateList[keyAsIndex] = gameManager.mergeDelta(stateList[keyAsIndex], deltaMap[key])
			} else if isMap {
				stateMap[key] = gameManager.mergeDelta(stateMap[key], deltaMap[key])
			}
		}
	}
	return state
}

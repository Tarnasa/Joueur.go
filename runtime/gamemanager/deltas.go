package gamemanager

import (
	"errors"
	"fmt"
	"joueur/runtime/errorhandler"
	"strconv"
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

func (this GameManager) initGameObjects(gameObjectDeltas map[string]interface{}) {
	for id, gameObjectDelta := range gameObjectDeltas {
		gameObjectDeltaAsMap, isMap := gameObjectDelta.(map[string]interface{})
		if !isMap {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot initialize new game object with id #"+id),
			)
		}

		rawGameObjectName, rawNameOk := gameObjectDeltaAsMap["gameObjectName"]
		gameObjectName, nameIsString := rawGameObjectName.(string)
		if !rawNameOk || !nameIsString || gameObjectName == "" {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("field 'gameObjectName' not a string on new game object #"+id),
			)
		}

		newGameObject, newGameObjectImpl, creationError := this.GameNamespace.CreateGameObject(gameObjectName)
		if creationError != nil {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Creation error on new game object "+gameObjectName+" #"+id),
			)
		}

		this.gameObjectImpls[id] = newGameObjectImpl

		internalGameObjectsRaw, ok := (*this.gameImpl).InternalDataMap["gameObjects"]
		internalGameObjectsMap, okConversion := internalGameObjectsRaw.(map[string]interface{})
		if !ok || !okConversion {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("Cannot find internal 'gameObjects' map on GameImpls's InternalDataMap"),
			)
		}
		internalGameObjectsMap[id] = newGameObject
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

func (this GameManager) mergeDelta(state interface{}, delta interface{}) interface{} {
	if this.isDeltaPrimitive(delta) {
		return delta
	}

	gameObject := this.getIfGameObjectReference(delta)
	if gameObject != nil {
		return gameObject
	}

	deltaMap, isDeltaMap := delta.(map[string]interface{})

	if !isDeltaMap {
		errorhandler.HandleError(
			errorhandler.DeltaMergeFailure,
			errors.New("Cannot merge non primitive and non map delta!"),
		)
	}
	deltaLengthValue, hasDeltaLength := deltaMap[this.ServerConstants.DeltaListLengthKey]

	if hasDeltaLength {
		// Then this part in the state is an array
		deltaLength, deltaLengthIsInt := deltaLengthValue.(int64)
		// We don't want to copy this key/value over to the state, it was just to signify the delta is an array
		delete(deltaMap, this.ServerConstants.DeltaListLengthKey)

		if !deltaLengthIsInt {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("DeltaListLength key present without being a number!"),
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

		if deltaValue == this.ServerConstants.DeltaRemoved && !isList {
			delete(stateMap, key)
		} else {
			if isList {
				stateList[keyAsIndex] = this.mergeDelta(stateList[keyAsIndex], deltaMap[key])
			} else if isMap {
				stateMap[key] = this.mergeDelta(stateMap[key], deltaMap[key])
			}
		}
	}
	return state
}

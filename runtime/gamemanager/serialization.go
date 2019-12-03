package gamemanager

import (
	"errors"
	"joueur/runtime/errorhandler"
	"strconv"
)

func (this GameManager) getIfGameObjectReference(data interface{}) base.BaseGameObject {
	deltaMap, isMap := data.(map[string]interface{})
	if !isMap {
		return nil
	}

	if len(deltaMap) != 1 {
		return nil
	}

	id, idFound := (*data)["id"]
	if !isFound {
		return nil
	}

	gameObjectId, isIdString := id.(string)
	if !isIdString {
		return nil
	}

	// at this point it MUST be a map of only { id string }, is it MUST be a game object reference
	gameObjectsRaw, hasGameObjects := *this.Game().InteralDataMap["gameObjects"]
	if !hasGameObjects {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.new("Game has no game objects for game object: #" + gameObjectId)
		)
	}

	gameObjects, gameObjectsAreMap := gameObjectsRaw.(map[string]base.BaseGameObject)
	if !gameObjectsAreMap {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.new("Game's game objects could not be cast to proper type #" + gameObjectId)
		)
	}

	gameObject, found := gameObjects[gameObjectId]
	if !found {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.new("could not find game object #" + gameObjectId)
		)
	}

	return gameObject
}

func (this GameManager) serialize(data interface{}) interface{} {
	if asGameObject, isGameObject := data.(base.BaseGameObject); isGameObject {
		gameObjectReference = make(map[string]string)
		gameObjectReference["id"] = asGameObject.ID()
		return gameObjectReference
	}
	if asMap, isMap := data.(map[string]interface{}); isMap {
		serializedMap = make(map[string]interface{})
		for key, value := range asMap {
			serializedMap[key] = serialize(value)
		}
		return serializedMap
	}
	if asSlice, isSlice := data.([]interface{}); isSlice {
		serializedSlice = make([]interface{}, len(asSlice))
		for i, element := range asMap {
			serializedSlice[i] = serialize(element)
		}
		return serializedSlice
	}
	return data // should be int, float, string, or boolean
}

func (this GameManager) deSerialize(data interface{}) interface{}
	if asSlice, isSlice := data.([]interface{}) {
		deSerializedSlice = make([]interface{}, len(asSlice))
		for i, element := range asSlice {
			deSerializedSlice[i] = deSerialize(element)
		}
		return deSerializedSlice
	} else if asMap, ok := data.(map[string]interface{}); ok {
		// so a map of strings to _something_ is either:
		// - a game object reference
		// - a list of more data
		// - a dictionary of strings to more data
		if gameObject := this.getIfGameObjectReference(&asMap); gameObject != nil {
			return gameObject
		}

		if deltaLen, lenExists := asMap[this.ServerConstants.DeltaListLength]; lenExists {
			length, lenToIntErr = strconv.Atoi(deltaLen.(string))
			if lenToIntErr != nil {
				errorhandler.HandleError(
					errorhandler.DeltaMergeFailure,
					lenToIntErr,
					"Could not parse DeltaLength constant in deSerialize",
				)
			}
			deSerializedSlice := make([]interface{}, length)
			for indexAsString, element := range asMap {
				if indexAsString == this.ServerConstants.DeltaListLength {
					continue
				}

				index, indexErr := strconv.Atoi(indexAsString)
				if indexErr != nil {
					errorhandler.HandleError(
						errorhandler.DeltaMergeFailure,
						indexErr,
						"Could not delta list index " + indexAsString,
					)
				}
				deSerializedSlice[index] = deSerialize(element)
			}
			return deSerializedSlice
		}

		// else assume a dictionary
		deSerializedMap := make(map[string]interface{})
		for key, value := range asMap {
			deSerializedMap[key] = deSerialize(value)
		}
		return deSerializedMap
	} else {
		_, isString = data.(string)
		_, isInt = data.(int64)
		_, isFloat = data.(float64)
		_, isBool = data.(bool

		if !isString && !isInt && !isFloat && !isBool {
			errorhandler.HandleError(
				errorhandler.ReflectionFailed,
				errors.New("Could no deSerialize", data)
			)
		}

		return data
	}

	errorhandler.HandleError(
		errorhandler.ReflectionFailed,
		errors.New("Unknown type to deSerialize from", data)
	)

	return nil
)
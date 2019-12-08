package gamemanager

import (
	"errors"
	"joueur/base"
	"joueur/runtime/errorhandler"
	"strconv"
)

func (gameManager *GameManager) getIfGameObjectReference(data interface{}) base.GameObject {
	deltaMap, isMap := data.(map[string]interface{})
	if !isMap {
		return nil
	}

	if len(deltaMap) != 1 {
		return nil
	}

	id, idFound := deltaMap["id"]
	if !idFound {
		return nil
	}

	gameObjectId, isIdString := id.(string)
	if !isIdString {
		return nil
	}

	gameObject, found := gameManager.gameImpl.GetGameObject(gameObjectId)
	if !found {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("could not find game object #"+gameObjectId),
		)
	}

	return gameObject
}

func (gameManager *GameManager) serialize(data interface{}) interface{} {
	if asGameObject, isGameObject := data.(base.GameObject); isGameObject {
		gameObjectReference := make(map[string]string)
		gameObjectReference["id"] = asGameObject.ID()
		return gameObjectReference
	}
	if asMap, isMap := data.(map[string]interface{}); isMap {
		serializedMap := make(map[string]interface{})
		for key, value := range asMap {
			serializedMap[key] = gameManager.serialize(value)
		}
		return serializedMap
	}
	if asSlice, isSlice := data.([]interface{}); isSlice {
		serializedSlice := make([]interface{}, len(asSlice))
		for i, element := range asSlice {
			serializedSlice[i] = gameManager.serialize(element)
		}
		return serializedSlice
	}
	return data // should be int, float, string, or boolean
}

func (gameManager *GameManager) deSerialize(data interface{}) interface{} {
	if asSlice, isSlice := data.([]interface{}); isSlice {
		deSerializedSlice := make([]interface{}, len(asSlice))
		for i, element := range asSlice {
			deSerializedSlice[i] = gameManager.deSerialize(element)
		}
		return deSerializedSlice
	} else if asMap, ok := data.(map[string]interface{}); ok {
		// so a map of strings to _something_ is either:
		// - a game object reference
		// - a list of more data
		// - a dictionary of strings to more data
		if gameObject := gameManager.getIfGameObjectReference(&asMap); gameObject != nil {
			return gameObject
		}

		if deltaLen, lenExists := asMap[gameManager.ServerConstants.DeltaListLengthKey]; lenExists {
			length, lenToIntErr := strconv.Atoi(deltaLen.(string))
			if lenToIntErr != nil {
				errorhandler.HandleError(
					errorhandler.DeltaMergeFailure,
					lenToIntErr,
					"Could not parse DeltaLength constant in deSerialize",
				)
			}
			deSerializedSlice := make([]interface{}, length)
			for indexAsString, element := range asMap {
				if indexAsString == gameManager.ServerConstants.DeltaListLengthKey {
					continue
				}

				index, indexErr := strconv.Atoi(indexAsString)
				if indexErr != nil {
					errorhandler.HandleError(
						errorhandler.DeltaMergeFailure,
						indexErr,
						"Could not delta list index "+indexAsString,
					)
				}
				deSerializedSlice[index] = gameManager.deSerialize(element)
			}
			return deSerializedSlice
		}

		// else assume a dictionary
		deSerializedMap := make(map[string]interface{})
		for key, value := range asMap {
			deSerializedMap[key] = gameManager.deSerialize(value)
		}
		return deSerializedMap
	} else {
		_, isString := data.(string)
		_, isInt := data.(int64)
		_, isFloat := data.(float64)
		_, isBool := data.(bool)

		if !isString && !isInt && !isFloat && !isBool {
			errorhandler.HandleError(
				errorhandler.ReflectionFailed,
				errors.New("Could no deSerialize data"),
			)
		}

		return data
	}

	errorhandler.HandleError(
		errorhandler.ReflectionFailed,
		errors.New("Unknown type to deSerialize from data"),
	)

	return nil
}

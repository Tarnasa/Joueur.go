package base

import (
	"errors"
	"fmt"
	"joueur/runtime/errorhandler"
	"strconv"
)

// DeltaMerge is a container of functions to facilitate type safe delta merging
type DeltaMerge interface {
	CannotConvertDeltaTo(string, interface{})

	String(interface{}) string
	Int(int64, interface{})
	Float(float64, interface{})
	Boolean(bool, interface{})

	BaseGameObject(interface{}) GameObject
	ToDeltaMap(interface{}) map[string]interface{}
	ToDeltaArray(interface{}) (map[int]interface{}, int)

	IsDeltaRemoved(interface{}) bool
}

func (deltaMergeImpl DeltaMergeImpl) CannotConvertDeltaTo(strName string, delta interface{}) {
	errorhandler.HandleError(
		errorhandler.ReflectionFailed,
		errors.New("cannot convert delta to "+strName+": "+fmt.Sprintf("%v", delta)),
	)
}

// DeltaMergeImpl is the base logic for primitive merging
type DeltaMergeImpl struct {
	getGameObject     func(string) GameObject
	deltaLengthKey    string
	deltaRemovedValue string
}

func (deltaMergeImpl DeltaMergeImpl) String(delta interface{}) string {
	asString, isString := delta.(string)

	if !isString {
		deltaMergeImpl.CannotConvertDeltaTo("string", delta)
	}

	return asString
}

func (deltaMergeImpl DeltaMergeImpl) Int(delta interface{}) int64 {
	asFloat, isFloat := delta.(float64)

	if !isFloat {
		deltaMergeImpl.CannotConvertDeltaTo("int64", delta)
	}

	return int64(asFloat)
}

func (deltaMergeImpl DeltaMergeImpl) Float(delta interface{}) float64 {
	asFloat, isFloat := delta.(float64)

	if !isFloat {
		deltaMergeImpl.CannotConvertDeltaTo("float64", delta)
	}

	return asFloat
}

func (deltaMergeImpl DeltaMergeImpl) Boolean(delta interface{}) bool {
	asBool, isBool := delta.(bool)

	if !isBool {
		deltaMergeImpl.CannotConvertDeltaTo("bool", delta)
	}

	return asBool
}

func (deltaMergeImpl DeltaMergeImpl) BaseGameObject(delta interface{}) GameObject {
	if delta == nil {
		return nil // nil pointer is valid for all game objects
	}

	deltaMap := deltaMergeImpl.ToDeltaMap(delta)
	rawID, hasID := deltaMap["id"]
	id, idIsString := rawID.(string)

	if !hasID || !idIsString || id == "" {
		deltaMergeImpl.CannotConvertDeltaTo("base.GameObject", delta)
	}

	gameObject := deltaMergeImpl.getGameObject(id)
	if gameObject == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot find game object #"+id+" to convert Delta reference to"),
		)
	}

	return gameObject
}

func (deltaMergeImpl DeltaMergeImpl) ToDeltaMap(delta interface{}) map[string]interface{} {
	asMap, isMap := delta.(map[string]interface{})

	if !isMap {
		deltaMergeImpl.CannotConvertDeltaTo("map[string]interface{}", delta)
	}

	return asMap
}

func (deltaMergeImpl DeltaMergeImpl) ToDeltaArray(delta interface{}) (map[int]interface{}, int) {
	deltaMap := deltaMergeImpl.ToDeltaMap(delta)
	deltaLengthString := deltaMergeImpl.String(deltaMap[deltaMergeImpl.deltaLengthKey])
	deltaLength, atoiErr := strconv.Atoi(deltaLengthString)
	if atoiErr != nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			atoiErr,
			"Cannot convert DeltaLength key to int: "+deltaLengthString,
		)
	}

	intMap := make(map[int]interface{})
	for deltaKey, deltaValue := range deltaMap {
		if deltaKey == deltaMergeImpl.deltaLengthKey || deltaValue == deltaMergeImpl.deltaRemovedValue {
			continue // we don't care about these entries
		}
		index, indexErr := strconv.Atoi(deltaKey)
		if indexErr != nil {
			errorhandler.HandleError(
				errorhandler.ReflectionFailed,
				indexErr,
				"Cannot convert array delta key to int: "+deltaKey,
			)
		}
		intMap[index] = deltaValue
	}

	return intMap, deltaLength
}

func (deltaMergeImpl DeltaMergeImpl) IsDeltaRemoved(delta interface{}) bool {
	return delta == deltaMergeImpl.deltaRemovedValue
}

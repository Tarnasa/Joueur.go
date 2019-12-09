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
	Int(interface{}) int64
	Float(interface{}) float64
	Boolean(interface{}) bool

	BaseGameObject(interface{}) GameObject
	ToDeltaMap(interface{}) map[string]interface{}
	ToDeltaArray(interface{}) (map[int]interface{}, int)

	IsDeltaRemoved(interface{}) bool
}

// DeltaMergeImpl is the base logic for primitive merging
type DeltaMergeImpl struct {
	Game              Game
	DeltaRemovedValue string
	DeltaLengthKey    string
}

// CannotConvertDeltaTo takes a string about why and a delta to explain why
// it's going to crash the program because a delta could not be merged.
func (deltaMergeImpl *DeltaMergeImpl) CannotConvertDeltaTo(strName string, delta interface{}) {
	errorhandler.HandleError(
		errorhandler.ReflectionFailed,
		errors.New("cannot convert delta to "+strName+": "+fmt.Sprintf("%v", delta)),
	)
}

// String attempts to extract a string from a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) String(delta interface{}) string {
	asString, isString := delta.(string)

	if !isString {
		deltaMergeImpl.CannotConvertDeltaTo("string", delta)
	}

	return asString
}

// Int attempts to extract an int from a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) Int(delta interface{}) int64 {
	asFloat, isFloat := delta.(float64)

	if !isFloat {
		deltaMergeImpl.CannotConvertDeltaTo("int64", delta)
	}

	return int64(asFloat)
}

// Float attempts to extract a float from a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) Float(delta interface{}) float64 {
	asFloat, isFloat := delta.(float64)

	if !isFloat {
		deltaMergeImpl.CannotConvertDeltaTo("float64", delta)
	}

	return asFloat
}

// Boolean attempts to extract a boolean from a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) Boolean(delta interface{}) bool {
	asBool, isBool := delta.(bool)

	if !isBool {
		deltaMergeImpl.CannotConvertDeltaTo("bool", delta)
	}

	return asBool
}

// BaseGameObject attempts to extract a base.GameObject from a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) BaseGameObject(delta interface{}) GameObject {
	if delta == nil {
		return nil // nil pointer is valid for all game objects
	}

	deltaMap := deltaMergeImpl.ToDeltaMap(delta)
	rawID, hasID := deltaMap["id"]
	id, idIsString := rawID.(string)

	if !hasID || !idIsString || id == "" {
		deltaMergeImpl.CannotConvertDeltaTo("base.GameObject", delta)
	}

	gameObject, found := deltaMergeImpl.Game.GetGameObject(id)
	if gameObject == nil || !found {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot find game object #"+id+" to convert Delta reference to"),
		)
	}

	return gameObject
}

// ToDeltaMap attempts to extract a map of string to more generics from
// a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) ToDeltaMap(delta interface{}) map[string]interface{} {
	asMap, isMap := delta.(map[string]interface{})

	if !isMap {
		deltaMergeImpl.CannotConvertDeltaTo("map[string]interface{}", delta)
	}

	return asMap
}

// ToDeltaArray attempts to extract a data for merging a list delta from
// a generic delta.
func (deltaMergeImpl *DeltaMergeImpl) ToDeltaArray(delta interface{}) (map[int]interface{}, int) {
	deltaMap := deltaMergeImpl.ToDeltaMap(delta)
	deltaLength := int(deltaMergeImpl.Int(deltaMap[deltaMergeImpl.DeltaLengthKey]))

	intMap := make(map[int]interface{})
	for deltaKey, deltaValue := range deltaMap {
		if deltaKey == deltaMergeImpl.DeltaLengthKey || deltaValue == deltaMergeImpl.DeltaRemovedValue {
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

// IsDeltaRemoved checks if a given delta is the special DeltaRemoved token.
func (deltaMergeImpl *DeltaMergeImpl) IsDeltaRemoved(delta interface{}) bool {
	return delta == deltaMergeImpl.DeltaRemovedValue
}

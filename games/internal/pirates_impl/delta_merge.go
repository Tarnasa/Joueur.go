package impl

import (
	"joueur/base"
	"joueur/games/pirates"
)

// DeltaMerge is the set of functions that can delta merge a
// Pirates game.
type DeltaMerge interface {
	base.DeltaMerge

	GameObject(interface{}) pirates.GameObject
	Player(interface{}) pirates.Player
	Port(interface{}) pirates.Port
	Tile(interface{}) pirates.Tile
	Unit(interface{}) pirates.Unit

	ArrayOfPlayer(*[]pirates.Player, interface{}) []pirates.Player
	ArrayOfPort(*[]pirates.Port, interface{}) []pirates.Port
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]pirates.Tile, interface{}) []pirates.Tile
	ArrayOfUnit(*[]pirates.Unit, interface{}) []pirates.Unit
	MapOfStringToGameObject(*map[string]pirates.GameObject, interface{}) map[string]pirates.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Pirates Game Object Deltas -- \\

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) pirates.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(pirates.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("pirates.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) pirates.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(pirates.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("pirates.Player", delta)
	}

	return asPlayer
}

// Port attempts to return the instance of Port
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Port(delta interface{}) pirates.Port {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPort, isPort := baseGameObject.(pirates.Port)
	if !isPort {
		(*deltaMergeImpl).CannotConvertDeltaTo("pirates.Port", delta)
	}

	return asPort
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) pirates.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(pirates.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("pirates.Tile", delta)
	}

	return asTile
}

// Unit attempts to return the instance of Unit
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Unit(delta interface{}) pirates.Unit {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnit, isUnit := baseGameObject.(pirates.Unit)
	if !isUnit {
		(*deltaMergeImpl).CannotConvertDeltaTo("pirates.Unit", delta)
	}

	return asUnit
}

// -- Deep Deltas -- \\

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]pirates.Player, delta interface{}) []pirates.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]pirates.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfPort delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPort(state *[]pirates.Port, delta interface{}) []pirates.Port {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]pirates.Port, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Port(deltaValue)
	}
	return newArray
}

// ArrayOfString delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfString(state *[]string, delta interface{}) []string {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]string, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.String(deltaValue)
	}
	return newArray
}

// ArrayOfTile delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]pirates.Tile, delta interface{}) []pirates.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]pirates.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// ArrayOfUnit delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnit(state *[]pirates.Unit, delta interface{}) []pirates.Unit {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]pirates.Unit, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Unit(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]pirates.GameObject, delta interface{}) map[string]pirates.GameObject {
	deltaMap := (*deltaMergeImpl).ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if (*deltaMergeImpl).IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = deltaMergeImpl.GameObject(deltaValue)
		}
	}
	return *state
}

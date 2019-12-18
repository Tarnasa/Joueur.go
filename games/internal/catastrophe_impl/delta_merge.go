package impl

import (
	"joueur/base"
	"joueur/games/catastrophe"
)

// DeltaMerge is the set of functions that can delta merge a
// Catastrophe game.
type DeltaMerge interface {
	base.DeltaMerge

	GameObject(interface{}) catastrophe.GameObject
	Job(interface{}) catastrophe.Job
	Player(interface{}) catastrophe.Player
	Structure(interface{}) catastrophe.Structure
	Tile(interface{}) catastrophe.Tile
	Unit(interface{}) catastrophe.Unit

	ArrayOfJob(*[]catastrophe.Job, interface{}) []catastrophe.Job
	ArrayOfPlayer(*[]catastrophe.Player, interface{}) []catastrophe.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfStructure(*[]catastrophe.Structure, interface{}) []catastrophe.Structure
	ArrayOfTile(*[]catastrophe.Tile, interface{}) []catastrophe.Tile
	ArrayOfUnit(*[]catastrophe.Unit, interface{}) []catastrophe.Unit
	MapOfStringToGameObject(*map[string]catastrophe.GameObject, interface{}) map[string]catastrophe.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Catastrophe Game Object Deltas -- \\

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) catastrophe.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(catastrophe.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.GameObject", delta)
	}

	return asGameObject
}

// Job attempts to return the instance of Job
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Job(delta interface{}) catastrophe.Job {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asJob, isJob := baseGameObject.(catastrophe.Job)
	if !isJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.Job", delta)
	}

	return asJob
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) catastrophe.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(catastrophe.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.Player", delta)
	}

	return asPlayer
}

// Structure attempts to return the instance of Structure
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Structure(delta interface{}) catastrophe.Structure {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asStructure, isStructure := baseGameObject.(catastrophe.Structure)
	if !isStructure {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.Structure", delta)
	}

	return asStructure
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) catastrophe.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(catastrophe.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.Tile", delta)
	}

	return asTile
}

// Unit attempts to return the instance of Unit
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Unit(delta interface{}) catastrophe.Unit {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnit, isUnit := baseGameObject.(catastrophe.Unit)
	if !isUnit {
		(*deltaMergeImpl).CannotConvertDeltaTo("catastrophe.Unit", delta)
	}

	return asUnit
}

// -- Deep Deltas -- \\

// ArrayOfJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfJob(state *[]catastrophe.Job, delta interface{}) []catastrophe.Job {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]catastrophe.Job, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Job(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]catastrophe.Player, delta interface{}) []catastrophe.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]catastrophe.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
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

// ArrayOfStructure delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfStructure(state *[]catastrophe.Structure, delta interface{}) []catastrophe.Structure {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]catastrophe.Structure, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Structure(deltaValue)
	}
	return newArray
}

// ArrayOfTile delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]catastrophe.Tile, delta interface{}) []catastrophe.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]catastrophe.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// ArrayOfUnit delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnit(state *[]catastrophe.Unit, delta interface{}) []catastrophe.Unit {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]catastrophe.Unit, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Unit(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]catastrophe.GameObject, delta interface{}) map[string]catastrophe.GameObject {
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

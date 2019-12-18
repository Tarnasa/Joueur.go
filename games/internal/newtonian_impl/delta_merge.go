package impl

import (
	"joueur/base"
	"joueur/games/newtonian"
)

// DeltaMerge is the set of functions that can delta merge a
// Newtonian game.
type DeltaMerge interface {
	base.DeltaMerge

	GameObject(interface{}) newtonian.GameObject
	Job(interface{}) newtonian.Job
	Machine(interface{}) newtonian.Machine
	Player(interface{}) newtonian.Player
	Tile(interface{}) newtonian.Tile
	Unit(interface{}) newtonian.Unit

	ArrayOfJob(*[]newtonian.Job, interface{}) []newtonian.Job
	ArrayOfMachine(*[]newtonian.Machine, interface{}) []newtonian.Machine
	ArrayOfPlayer(*[]newtonian.Player, interface{}) []newtonian.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]newtonian.Tile, interface{}) []newtonian.Tile
	ArrayOfUnit(*[]newtonian.Unit, interface{}) []newtonian.Unit
	MapOfStringToGameObject(*map[string]newtonian.GameObject, interface{}) map[string]newtonian.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Newtonian Game Object Deltas -- \\

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) newtonian.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(newtonian.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.GameObject", delta)
	}

	return asGameObject
}

// Job attempts to return the instance of Job
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Job(delta interface{}) newtonian.Job {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asJob, isJob := baseGameObject.(newtonian.Job)
	if !isJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.Job", delta)
	}

	return asJob
}

// Machine attempts to return the instance of Machine
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Machine(delta interface{}) newtonian.Machine {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asMachine, isMachine := baseGameObject.(newtonian.Machine)
	if !isMachine {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.Machine", delta)
	}

	return asMachine
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) newtonian.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(newtonian.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.Player", delta)
	}

	return asPlayer
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) newtonian.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(newtonian.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.Tile", delta)
	}

	return asTile
}

// Unit attempts to return the instance of Unit
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Unit(delta interface{}) newtonian.Unit {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnit, isUnit := baseGameObject.(newtonian.Unit)
	if !isUnit {
		(*deltaMergeImpl).CannotConvertDeltaTo("newtonian.Unit", delta)
	}

	return asUnit
}

// -- Deep Deltas -- \\

// ArrayOfJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfJob(state *[]newtonian.Job, delta interface{}) []newtonian.Job {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]newtonian.Job, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Job(deltaValue)
	}
	return newArray
}

// ArrayOfMachine delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfMachine(state *[]newtonian.Machine, delta interface{}) []newtonian.Machine {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]newtonian.Machine, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Machine(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]newtonian.Player, delta interface{}) []newtonian.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]newtonian.Player, listLength) // resize array with new copy
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

// ArrayOfTile delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]newtonian.Tile, delta interface{}) []newtonian.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]newtonian.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// ArrayOfUnit delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnit(state *[]newtonian.Unit, delta interface{}) []newtonian.Unit {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]newtonian.Unit, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Unit(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]newtonian.GameObject, delta interface{}) map[string]newtonian.GameObject {
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

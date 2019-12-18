package impl

import (
	"joueur/base"
	"joueur/games/necrowar"
)

// DeltaMerge is the set of functions that can delta merge a
// Necrowar game.
type DeltaMerge interface {
	base.DeltaMerge

	GameObject(interface{}) necrowar.GameObject
	Player(interface{}) necrowar.Player
	Tile(interface{}) necrowar.Tile
	Tower(interface{}) necrowar.Tower
	TowerJob(interface{}) necrowar.TowerJob
	Unit(interface{}) necrowar.Unit
	UnitJob(interface{}) necrowar.UnitJob

	ArrayOfPlayer(*[]necrowar.Player, interface{}) []necrowar.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]necrowar.Tile, interface{}) []necrowar.Tile
	ArrayOfTower(*[]necrowar.Tower, interface{}) []necrowar.Tower
	ArrayOfTowerJob(*[]necrowar.TowerJob, interface{}) []necrowar.TowerJob
	ArrayOfUnit(*[]necrowar.Unit, interface{}) []necrowar.Unit
	ArrayOfUnitJob(*[]necrowar.UnitJob, interface{}) []necrowar.UnitJob
	MapOfStringToGameObject(*map[string]necrowar.GameObject, interface{}) map[string]necrowar.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Necrowar Game Object Deltas -- \\

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) necrowar.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(necrowar.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) necrowar.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(necrowar.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.Player", delta)
	}

	return asPlayer
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) necrowar.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(necrowar.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.Tile", delta)
	}

	return asTile
}

// Tower attempts to return the instance of Tower
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tower(delta interface{}) necrowar.Tower {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTower, isTower := baseGameObject.(necrowar.Tower)
	if !isTower {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.Tower", delta)
	}

	return asTower
}

// TowerJob attempts to return the instance of TowerJob
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) TowerJob(delta interface{}) necrowar.TowerJob {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTowerJob, isTowerJob := baseGameObject.(necrowar.TowerJob)
	if !isTowerJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.TowerJob", delta)
	}

	return asTowerJob
}

// Unit attempts to return the instance of Unit
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Unit(delta interface{}) necrowar.Unit {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnit, isUnit := baseGameObject.(necrowar.Unit)
	if !isUnit {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.Unit", delta)
	}

	return asUnit
}

// UnitJob attempts to return the instance of UnitJob
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) UnitJob(delta interface{}) necrowar.UnitJob {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnitJob, isUnitJob := baseGameObject.(necrowar.UnitJob)
	if !isUnitJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("necrowar.UnitJob", delta)
	}

	return asUnitJob
}

// -- Deep Deltas -- \\

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]necrowar.Player, delta interface{}) []necrowar.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.Player, listLength) // resize array with new copy
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
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]necrowar.Tile, delta interface{}) []necrowar.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// ArrayOfTower delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTower(state *[]necrowar.Tower, delta interface{}) []necrowar.Tower {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.Tower, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tower(deltaValue)
	}
	return newArray
}

// ArrayOfTowerJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTowerJob(state *[]necrowar.TowerJob, delta interface{}) []necrowar.TowerJob {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.TowerJob, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.TowerJob(deltaValue)
	}
	return newArray
}

// ArrayOfUnit delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnit(state *[]necrowar.Unit, delta interface{}) []necrowar.Unit {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.Unit, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Unit(deltaValue)
	}
	return newArray
}

// ArrayOfUnitJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnitJob(state *[]necrowar.UnitJob, delta interface{}) []necrowar.UnitJob {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]necrowar.UnitJob, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.UnitJob(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]necrowar.GameObject, delta interface{}) map[string]necrowar.GameObject {
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

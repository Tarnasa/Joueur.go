package impl

import (
	"joueur/base"
	"joueur/games/galapagos"
)

// DeltaMerge is the set of functions that can delta merge a
// Galapagos game.
type DeltaMerge interface {
	base.DeltaMerge

	Creature(interface{}) galapagos.Creature
	GameObject(interface{}) galapagos.GameObject
	Plant(interface{}) galapagos.Plant
	Player(interface{}) galapagos.Player
	Tile(interface{}) galapagos.Tile

	ArrayOfCreature(*[]galapagos.Creature, interface{}) []galapagos.Creature
	ArrayOfPlant(*[]galapagos.Plant, interface{}) []galapagos.Plant
	ArrayOfPlayer(*[]galapagos.Player, interface{}) []galapagos.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]galapagos.Tile, interface{}) []galapagos.Tile
	MapOfStringToGameObject(*map[string]galapagos.GameObject, interface{}) map[string]galapagos.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Galapagos Game Object Deltas -- \\

// Creature attempts to return the instance of Creature
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Creature(delta interface{}) galapagos.Creature {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asCreature, isCreature := baseGameObject.(galapagos.Creature)
	if !isCreature {
		(*deltaMergeImpl).CannotConvertDeltaTo("galapagos.Creature", delta)
	}

	return asCreature
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) galapagos.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(galapagos.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("galapagos.GameObject", delta)
	}

	return asGameObject
}

// Plant attempts to return the instance of Plant
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Plant(delta interface{}) galapagos.Plant {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlant, isPlant := baseGameObject.(galapagos.Plant)
	if !isPlant {
		(*deltaMergeImpl).CannotConvertDeltaTo("galapagos.Plant", delta)
	}

	return asPlant
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) galapagos.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(galapagos.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("galapagos.Player", delta)
	}

	return asPlayer
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) galapagos.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(galapagos.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("galapagos.Tile", delta)
	}

	return asTile
}

// -- Deep Deltas -- \\

// ArrayOfCreature delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfCreature(state *[]galapagos.Creature, delta interface{}) []galapagos.Creature {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]galapagos.Creature, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Creature(deltaValue)
	}
	return newArray
}

// ArrayOfPlant delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlant(state *[]galapagos.Plant, delta interface{}) []galapagos.Plant {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]galapagos.Plant, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Plant(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]galapagos.Player, delta interface{}) []galapagos.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]galapagos.Player, listLength) // resize array with new copy
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
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]galapagos.Tile, delta interface{}) []galapagos.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]galapagos.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]galapagos.GameObject, delta interface{}) map[string]galapagos.GameObject {
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

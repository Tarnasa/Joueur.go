package impl

import (
	"joueur/base"
	"joueur/games/saloon"
)

type DeltaMerge interface {
	base.DeltaMerge

	Bottle(interface{}) saloon.Bottle
	Cowboy(interface{}) saloon.Cowboy
	Furnishing(interface{}) saloon.Furnishing
	GameObject(interface{}) saloon.GameObject
	Player(interface{}) saloon.Player
	Tile(interface{}) saloon.Tile
	YoungGun(interface{}) saloon.YoungGun

	ArrayOfBottle(*[]saloon.Bottle, interface{}) []saloon.Bottle
	ArrayOfCowboy(*[]saloon.Cowboy, interface{}) []saloon.Cowboy
	ArrayOfFurnishing(*[]saloon.Furnishing, interface{}) []saloon.Furnishing
	ArrayOfPlayer(*[]saloon.Player, interface{}) []saloon.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]saloon.Tile, interface{}) []saloon.Tile
	MapOfStringToGameObject(*map[string]saloon.GameObject, interface{}) map[string]saloon.GameObject
}

type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Saloon Game Object Deltas -- \\

// Bottle attempts to return the instance of Bottle
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Bottle(delta interface{}) saloon.Bottle {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBottle, isBottle := baseGameObject.(saloon.Bottle)
	if !isBottle {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.Bottle", delta)
	}

	return asBottle
}

// Cowboy attempts to return the instance of Cowboy
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Cowboy(delta interface{}) saloon.Cowboy {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asCowboy, isCowboy := baseGameObject.(saloon.Cowboy)
	if !isCowboy {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.Cowboy", delta)
	}

	return asCowboy
}

// Furnishing attempts to return the instance of Furnishing
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Furnishing(delta interface{}) saloon.Furnishing {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asFurnishing, isFurnishing := baseGameObject.(saloon.Furnishing)
	if !isFurnishing {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.Furnishing", delta)
	}

	return asFurnishing
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) saloon.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(saloon.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) saloon.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(saloon.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.Player", delta)
	}

	return asPlayer
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) saloon.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(saloon.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.Tile", delta)
	}

	return asTile
}

// YoungGun attempts to return the instance of YoungGun
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) YoungGun(delta interface{}) saloon.YoungGun {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asYoungGun, isYoungGun := baseGameObject.(saloon.YoungGun)
	if !isYoungGun {
		(*deltaMergeImpl).CannotConvertDeltaTo("saloon.YoungGun", delta)
	}

	return asYoungGun
}

// -- Deep Deltas -- \\

// ArrayOfBottle delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfBottle(state *[]saloon.Bottle, delta interface{}) []saloon.Bottle {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]saloon.Bottle, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Bottle(deltaValue)
	}
	return newArray
}

// ArrayOfCowboy delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfCowboy(state *[]saloon.Cowboy, delta interface{}) []saloon.Cowboy {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]saloon.Cowboy, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Cowboy(deltaValue)
	}
	return newArray
}

// ArrayOfFurnishing delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfFurnishing(state *[]saloon.Furnishing, delta interface{}) []saloon.Furnishing {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]saloon.Furnishing, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Furnishing(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]saloon.Player, delta interface{}) []saloon.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]saloon.Player, listLength) // resize array with new copy
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
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]saloon.Tile, delta interface{}) []saloon.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]saloon.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]saloon.GameObject, delta interface{}) map[string]saloon.GameObject {
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

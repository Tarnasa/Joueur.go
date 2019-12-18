package impl

import (
	"joueur/base"
	"joueur/games/checkers"
)

// DeltaMerge is the set of functions that can delta merge a
// Checkers game.
type DeltaMerge interface {
	base.DeltaMerge

	Checker(interface{}) checkers.Checker
	GameObject(interface{}) checkers.GameObject
	Player(interface{}) checkers.Player

	ArrayOfChecker(*[]checkers.Checker, interface{}) []checkers.Checker
	ArrayOfPlayer(*[]checkers.Player, interface{}) []checkers.Player
	ArrayOfString(*[]string, interface{}) []string
	MapOfStringToGameObject(*map[string]checkers.GameObject, interface{}) map[string]checkers.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Checkers Game Object Deltas -- \\

// Checker attempts to return the instance of Checker
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Checker(delta interface{}) checkers.Checker {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asChecker, isChecker := baseGameObject.(checkers.Checker)
	if !isChecker {
		(*deltaMergeImpl).CannotConvertDeltaTo("checkers.Checker", delta)
	}

	return asChecker
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) checkers.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(checkers.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("checkers.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) checkers.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(checkers.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("checkers.Player", delta)
	}

	return asPlayer
}

// -- Deep Deltas -- \\

// ArrayOfChecker delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfChecker(state *[]checkers.Checker, delta interface{}) []checkers.Checker {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]checkers.Checker, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Checker(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]checkers.Player, delta interface{}) []checkers.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]checkers.Player, listLength) // resize array with new copy
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

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]checkers.GameObject, delta interface{}) map[string]checkers.GameObject {
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

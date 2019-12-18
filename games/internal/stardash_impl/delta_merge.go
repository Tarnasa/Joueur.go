package impl

import (
	"joueur/base"
	"joueur/games/stardash"
)

// DeltaMerge is the set of functions that can delta merge a
// Stardash game.
type DeltaMerge interface {
	base.DeltaMerge

	Body(interface{}) stardash.Body
	GameObject(interface{}) stardash.GameObject
	Job(interface{}) stardash.Job
	Player(interface{}) stardash.Player
	Projectile(interface{}) stardash.Projectile
	Unit(interface{}) stardash.Unit

	ArrayOfBody(*[]stardash.Body, interface{}) []stardash.Body
	ArrayOfJob(*[]stardash.Job, interface{}) []stardash.Job
	ArrayOfPlayer(*[]stardash.Player, interface{}) []stardash.Player
	ArrayOfProjectile(*[]stardash.Projectile, interface{}) []stardash.Projectile
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfUnit(*[]stardash.Unit, interface{}) []stardash.Unit
	MapOfStringToGameObject(*map[string]stardash.GameObject, interface{}) map[string]stardash.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Stardash Game Object Deltas -- \\

// Body attempts to return the instance of Body
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Body(delta interface{}) stardash.Body {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBody, isBody := baseGameObject.(stardash.Body)
	if !isBody {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.Body", delta)
	}

	return asBody
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) stardash.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(stardash.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.GameObject", delta)
	}

	return asGameObject
}

// Job attempts to return the instance of Job
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Job(delta interface{}) stardash.Job {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asJob, isJob := baseGameObject.(stardash.Job)
	if !isJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.Job", delta)
	}

	return asJob
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) stardash.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(stardash.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.Player", delta)
	}

	return asPlayer
}

// Projectile attempts to return the instance of Projectile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Projectile(delta interface{}) stardash.Projectile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asProjectile, isProjectile := baseGameObject.(stardash.Projectile)
	if !isProjectile {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.Projectile", delta)
	}

	return asProjectile
}

// Unit attempts to return the instance of Unit
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Unit(delta interface{}) stardash.Unit {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUnit, isUnit := baseGameObject.(stardash.Unit)
	if !isUnit {
		(*deltaMergeImpl).CannotConvertDeltaTo("stardash.Unit", delta)
	}

	return asUnit
}

// -- Deep Deltas -- \\

// ArrayOfBody delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfBody(state *[]stardash.Body, delta interface{}) []stardash.Body {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stardash.Body, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Body(deltaValue)
	}
	return newArray
}

// ArrayOfJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfJob(state *[]stardash.Job, delta interface{}) []stardash.Job {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stardash.Job, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Job(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]stardash.Player, delta interface{}) []stardash.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stardash.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfProjectile delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfProjectile(state *[]stardash.Projectile, delta interface{}) []stardash.Projectile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stardash.Projectile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Projectile(deltaValue)
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

// ArrayOfUnit delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUnit(state *[]stardash.Unit, delta interface{}) []stardash.Unit {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stardash.Unit, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Unit(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]stardash.GameObject, delta interface{}) map[string]stardash.GameObject {
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

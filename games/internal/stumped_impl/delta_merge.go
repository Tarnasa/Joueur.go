package impl

import (
	"joueur/base"
	"joueur/games/stumped"
)

// DeltaMerge is the set of functions that can delta merge a
// Stumped game.
type DeltaMerge interface {
	base.DeltaMerge

	Beaver(interface{}) stumped.Beaver
	GameObject(interface{}) stumped.GameObject
	Job(interface{}) stumped.Job
	Player(interface{}) stumped.Player
	Spawner(interface{}) stumped.Spawner
	Tile(interface{}) stumped.Tile

	ArrayOfBeaver(*[]stumped.Beaver, interface{}) []stumped.Beaver
	ArrayOfJob(*[]stumped.Job, interface{}) []stumped.Job
	ArrayOfPlayer(*[]stumped.Player, interface{}) []stumped.Player
	ArrayOfSpawner(*[]stumped.Spawner, interface{}) []stumped.Spawner
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]stumped.Tile, interface{}) []stumped.Tile
	MapOfStringToGameObject(*map[string]stumped.GameObject, interface{}) map[string]stumped.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Stumped Game Object Deltas -- \\

// Beaver attempts to return the instance of Beaver
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Beaver(delta interface{}) stumped.Beaver {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBeaver, isBeaver := baseGameObject.(stumped.Beaver)
	if !isBeaver {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.Beaver", delta)
	}

	return asBeaver
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) stumped.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(stumped.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.GameObject", delta)
	}

	return asGameObject
}

// Job attempts to return the instance of Job
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Job(delta interface{}) stumped.Job {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asJob, isJob := baseGameObject.(stumped.Job)
	if !isJob {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.Job", delta)
	}

	return asJob
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) stumped.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(stumped.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.Player", delta)
	}

	return asPlayer
}

// Spawner attempts to return the instance of Spawner
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Spawner(delta interface{}) stumped.Spawner {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asSpawner, isSpawner := baseGameObject.(stumped.Spawner)
	if !isSpawner {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.Spawner", delta)
	}

	return asSpawner
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) stumped.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(stumped.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("stumped.Tile", delta)
	}

	return asTile
}

// -- Deep Deltas -- \\

// ArrayOfBeaver delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfBeaver(state *[]stumped.Beaver, delta interface{}) []stumped.Beaver {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stumped.Beaver, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Beaver(deltaValue)
	}
	return newArray
}

// ArrayOfJob delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfJob(state *[]stumped.Job, delta interface{}) []stumped.Job {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stumped.Job, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Job(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]stumped.Player, delta interface{}) []stumped.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stumped.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfSpawner delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfSpawner(state *[]stumped.Spawner, delta interface{}) []stumped.Spawner {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stumped.Spawner, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Spawner(deltaValue)
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
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]stumped.Tile, delta interface{}) []stumped.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]stumped.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]stumped.GameObject, delta interface{}) map[string]stumped.GameObject {
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

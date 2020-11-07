package impl

import (
	"joueur/base"
	"joueur/games/coreminer"
)

// DeltaMerge is the set of functions that can delta merge a
// Coreminer game.
type DeltaMerge interface {
	base.DeltaMerge

	Bomb(interface{}) coreminer.Bomb
	GameObject(interface{}) coreminer.GameObject
	Miner(interface{}) coreminer.Miner
	Player(interface{}) coreminer.Player
	Tile(interface{}) coreminer.Tile
	Upgrade(interface{}) coreminer.Upgrade

	ArrayOfBomb(*[]coreminer.Bomb, interface{}) []coreminer.Bomb
	ArrayOfMiner(*[]coreminer.Miner, interface{}) []coreminer.Miner
	ArrayOfPlayer(*[]coreminer.Player, interface{}) []coreminer.Player
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfTile(*[]coreminer.Tile, interface{}) []coreminer.Tile
	ArrayOfUpgrade(*[]coreminer.Upgrade, interface{}) []coreminer.Upgrade
	MapOfStringToGameObject(*map[string]coreminer.GameObject, interface{}) map[string]coreminer.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Coreminer Game Object Deltas -- \\

// Bomb attempts to return the instance of Bomb
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Bomb(delta interface{}) coreminer.Bomb {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBomb, isBomb := baseGameObject.(coreminer.Bomb)
	if !isBomb {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.Bomb", delta)
	}

	return asBomb
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) coreminer.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(coreminer.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.GameObject", delta)
	}

	return asGameObject
}

// Miner attempts to return the instance of Miner
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Miner(delta interface{}) coreminer.Miner {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asMiner, isMiner := baseGameObject.(coreminer.Miner)
	if !isMiner {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.Miner", delta)
	}

	return asMiner
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) coreminer.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(coreminer.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.Player", delta)
	}

	return asPlayer
}

// Tile attempts to return the instance of Tile
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Tile(delta interface{}) coreminer.Tile {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asTile, isTile := baseGameObject.(coreminer.Tile)
	if !isTile {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.Tile", delta)
	}

	return asTile
}

// Upgrade attempts to return the instance of Upgrade
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Upgrade(delta interface{}) coreminer.Upgrade {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asUpgrade, isUpgrade := baseGameObject.(coreminer.Upgrade)
	if !isUpgrade {
		(*deltaMergeImpl).CannotConvertDeltaTo("coreminer.Upgrade", delta)
	}

	return asUpgrade
}

// -- Deep Deltas -- \\

// ArrayOfBomb delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfBomb(state *[]coreminer.Bomb, delta interface{}) []coreminer.Bomb {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]coreminer.Bomb, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Bomb(deltaValue)
	}
	return newArray
}

// ArrayOfMiner delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfMiner(state *[]coreminer.Miner, delta interface{}) []coreminer.Miner {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]coreminer.Miner, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Miner(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]coreminer.Player, delta interface{}) []coreminer.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]coreminer.Player, listLength) // resize array with new copy
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
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfTile(state *[]coreminer.Tile, delta interface{}) []coreminer.Tile {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]coreminer.Tile, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Tile(deltaValue)
	}
	return newArray
}

// ArrayOfUpgrade delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfUpgrade(state *[]coreminer.Upgrade, delta interface{}) []coreminer.Upgrade {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]coreminer.Upgrade, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Upgrade(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]coreminer.GameObject, delta interface{}) map[string]coreminer.GameObject {
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

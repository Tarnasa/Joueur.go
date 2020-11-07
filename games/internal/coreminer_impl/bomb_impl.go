package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// BombImpl is the struct that implements the container for Bomb instances
// in Coreminer.
type BombImpl struct {
	GameObjectImpl

	tileImpl  coreminer.Tile
	timerImpl int64
}

// Tile returns the Tile this Miner is on.
//
// Value can be returned as a nil pointer.
func (bombImpl *BombImpl) Tile() coreminer.Tile {
	return bombImpl.tileImpl
}

// Timer returns the number of turns before this Bomb explodes. Zero means
// it will explode after the current turn.
func (bombImpl *BombImpl) Timer() int64 {
	return bombImpl.timerImpl
}

// InitImplDefaults initializes safe defaults for all fields in Bomb.
func (bombImpl *BombImpl) InitImplDefaults() {
	bombImpl.GameObjectImpl.InitImplDefaults()

	bombImpl.tileImpl = nil
	bombImpl.timerImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Bomb.
func (bombImpl *BombImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := bombImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	coreminerDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'coreminer.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "tile":
		bombImpl.tileImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "timer":
		bombImpl.timerImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

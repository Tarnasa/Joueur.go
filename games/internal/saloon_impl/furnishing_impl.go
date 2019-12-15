package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// FurnishingImpl is the struct that implements the container for
// Furnishing instances in Saloon.
type FurnishingImpl struct {
	GameObjectImpl

	healthImpl      int64
	isDestroyedImpl bool
	isPianoImpl     bool
	isPlayingImpl   bool
	tileImpl        saloon.Tile
}

// Health returns how much health this Furnishing currently has.
func (furnishingImpl *FurnishingImpl) Health() int64 {
	return furnishingImpl.healthImpl
}

// IsDestroyed returns if this Furnishing has been destroyed, and has been
// removed from the game.
func (furnishingImpl *FurnishingImpl) IsDestroyed() bool {
	return furnishingImpl.isDestroyedImpl
}

// IsPiano returns true if this Furnishing is a piano and can be played,
// False otherwise.
func (furnishingImpl *FurnishingImpl) IsPiano() bool {
	return furnishingImpl.isPianoImpl
}

// IsPlaying returns if this is a piano and a Cowboy is playing it this
// turn.
func (furnishingImpl *FurnishingImpl) IsPlaying() bool {
	return furnishingImpl.isPlayingImpl
}

// Tile returns the Tile that this Furnishing is located on.
func (furnishingImpl *FurnishingImpl) Tile() saloon.Tile {
	return furnishingImpl.tileImpl
}

// InitImplDefaults initializes safe defaults for all fields in Furnishing.
func (furnishingImpl *FurnishingImpl) InitImplDefaults() {
	furnishingImpl.GameObjectImpl.InitImplDefaults()

	furnishingImpl.healthImpl = 0
	furnishingImpl.isDestroyedImpl = true
	furnishingImpl.isPianoImpl = true
	furnishingImpl.isPlayingImpl = true
	furnishingImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Furnishing.
func (furnishingImpl *FurnishingImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := furnishingImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	saloonDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: "+
			"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "health":
		furnishingImpl.healthImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "isDestroyed":
		furnishingImpl.isDestroyedImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "isPiano":
		furnishingImpl.isPianoImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "isPlaying":
		furnishingImpl.isPlayingImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "tile":
		furnishingImpl.tileImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// BottleImpl is the struct that implements the container for Bottle
// instances in Saloon.
type BottleImpl struct {
	GameObjectImpl

	directionImpl      string
	drunkDirectionImpl string
	isDestroyedImpl    bool
	tileImpl           saloon.Tile
}

// Direction returns the Direction this Bottle is flying and will move to
// between turns, can be 'North', 'East', 'South', or 'West'.
//
// Literal Values: "North", "East", "South", or "West"
func (bottleImpl *BottleImpl) Direction() string {
	return bottleImpl.directionImpl
}

// DrunkDirection returns the direction any Cowboys hit by this will move,
// can be 'North', 'East', 'South', or 'West'.
//
// Literal Values: "North", "East", "South", or "West"
func (bottleImpl *BottleImpl) DrunkDirection() string {
	return bottleImpl.drunkDirectionImpl
}

// IsDestroyed returns true if this Bottle has impacted and has been
// destroyed (removed from the Game). False if still in the game flying
// through the saloon.
func (bottleImpl *BottleImpl) IsDestroyed() bool {
	return bottleImpl.isDestroyedImpl
}

// Tile returns the Tile this bottle is currently flying over.
//
// Value can be returned as a nil pointer.
func (bottleImpl *BottleImpl) Tile() saloon.Tile {
	return bottleImpl.tileImpl
}

// InitImplDefaults initializes safe defaults for all fields in Bottle.
func (bottleImpl *BottleImpl) InitImplDefaults() {
	bottleImpl.GameObjectImpl.InitImplDefaults()

	bottleImpl.directionImpl = ""
	bottleImpl.drunkDirectionImpl = ""
	bottleImpl.isDestroyedImpl = true
	bottleImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Bottle.
func (bottleImpl *BottleImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := bottleImpl.GameObjectImpl.DeltaMerge(
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
			"deltaMerge is not the expected type of: " +
				"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "direction":
		bottleImpl.directionImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "drunkDirection":
		bottleImpl.drunkDirectionImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "isDestroyed":
		bottleImpl.isDestroyedImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "tile":
		bottleImpl.tileImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

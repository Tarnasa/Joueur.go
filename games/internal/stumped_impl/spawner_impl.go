package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// SpawnerImpl is the struct that implements the container for Spawner
// instances in Stumped.
type SpawnerImpl struct {
	GameObjectImpl

	hasBeenHarvestedImpl bool
	healthImpl           int64
	tileImpl             stumped.Tile
	typeImpl             string
}

// HasBeenHarvested returns true if this Spawner has been harvested this
// turn, and it will not heal at the end of the turn, false otherwise.
func (spawnerImpl *SpawnerImpl) HasBeenHarvested() bool {
	return spawnerImpl.hasBeenHarvestedImpl
}

// Health returns how much health this Spawner has, which is used to
// calculate how much of its resource can be harvested.
func (spawnerImpl *SpawnerImpl) Health() int64 {
	return spawnerImpl.healthImpl
}

// Tile returns the Tile this Spawner is on.
func (spawnerImpl *SpawnerImpl) Tile() stumped.Tile {
	return spawnerImpl.tileImpl
}

// Type returns what type of resource this is ('food' or 'branches').
//
// Literal Values: "food" or "branches"
func (spawnerImpl *SpawnerImpl) Type() string {
	return spawnerImpl.typeImpl
}

// InitImplDefaults initializes safe defaults for all fields in Spawner.
func (spawnerImpl *SpawnerImpl) InitImplDefaults() {
	spawnerImpl.GameObjectImpl.InitImplDefaults()

	spawnerImpl.hasBeenHarvestedImpl = true
	spawnerImpl.healthImpl = 0
	spawnerImpl.tileImpl = nil
	spawnerImpl.typeImpl = ""
}

// DeltaMerge merges the delta for a given attribute in Spawner.
func (spawnerImpl *SpawnerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := spawnerImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	stumpedDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stumped.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "hasBeenHarvested":
		spawnerImpl.hasBeenHarvestedImpl = stumpedDeltaMerge.Boolean(delta)
		return true, nil
	case "health":
		spawnerImpl.healthImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "tile":
		spawnerImpl.tileImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "type":
		spawnerImpl.typeImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

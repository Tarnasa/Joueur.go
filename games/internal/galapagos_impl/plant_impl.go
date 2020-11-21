package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// PlantImpl is the struct that implements the container for Plant
// instances in Galapagos.
type PlantImpl struct {
	GameObjectImpl

	growthRateImpl       int64
	sizeImpl             int64
	tileImpl             galapagos.Tile
	turnsUntilGrowthImpl int64
}

// GrowthRate returns the total number of turns it takes this plant to grow
// in size.
func (plantImpl *PlantImpl) GrowthRate() int64 {
	return plantImpl.growthRateImpl
}

// Size returns the size of the plant.
func (plantImpl *PlantImpl) Size() int64 {
	return plantImpl.sizeImpl
}

// Tile returns the Tile this Plant occupies.
//
// Value can be returned as a nil pointer.
func (plantImpl *PlantImpl) Tile() galapagos.Tile {
	return plantImpl.tileImpl
}

// TurnsUntilGrowth returns the number of turns left until this plant will
// grow again.
func (plantImpl *PlantImpl) TurnsUntilGrowth() int64 {
	return plantImpl.turnsUntilGrowthImpl
}

// InitImplDefaults initializes safe defaults for all fields in Plant.
func (plantImpl *PlantImpl) InitImplDefaults() {
	plantImpl.GameObjectImpl.InitImplDefaults()

	plantImpl.growthRateImpl = 0
	plantImpl.sizeImpl = 0
	plantImpl.tileImpl = nil
	plantImpl.turnsUntilGrowthImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Plant.
func (plantImpl *PlantImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := plantImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	galapagosDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'galapagos.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "growthRate":
		plantImpl.growthRateImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "size":
		plantImpl.sizeImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "tile":
		plantImpl.tileImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	case "turnsUntilGrowth":
		plantImpl.turnsUntilGrowthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

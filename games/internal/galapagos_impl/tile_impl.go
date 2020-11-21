package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// TileImpl is the struct that implements the container for Tile instances
// in Galapagos.
type TileImpl struct {
	GameObjectImpl

	creatureImpl  galapagos.Creature
	eggImpl       galapagos.Creature
	plantImpl     galapagos.Plant
	tileEastImpl  galapagos.Tile
	tileNorthImpl galapagos.Tile
	tileSouthImpl galapagos.Tile
	tileWestImpl  galapagos.Tile
	xImpl         int64
	yImpl         int64
}

// Creature returns the Creature on this Tile or nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Creature() galapagos.Creature {
	return tileImpl.creatureImpl
}

// Egg returns the unhatched Creature on this Tile or nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Egg() galapagos.Creature {
	return tileImpl.eggImpl
}

// Plant returns the Plant on this Tile or nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Plant() galapagos.Plant {
	return tileImpl.plantImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() galapagos.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() galapagos.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() galapagos.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() galapagos.Tile {
	return tileImpl.tileWestImpl
}

// X returns the x (horizontal) position of this Tile.
func (tileImpl *TileImpl) X() int64 {
	return tileImpl.xImpl
}

// Y returns the y (vertical) position of this Tile.
func (tileImpl *TileImpl) Y() int64 {
	return tileImpl.yImpl
}

// InitImplDefaults initializes safe defaults for all fields in Tile.
func (tileImpl *TileImpl) InitImplDefaults() {
	tileImpl.GameObjectImpl.InitImplDefaults()

	tileImpl.creatureImpl = nil
	tileImpl.eggImpl = nil
	tileImpl.plantImpl = nil
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.xImpl = 0
	tileImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Tile.
func (tileImpl *TileImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := tileImpl.GameObjectImpl.DeltaMerge(
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
	case "creature":
		tileImpl.creatureImpl = galapagosDeltaMerge.Creature(delta)
		return true, nil
	case "egg":
		tileImpl.eggImpl = galapagosDeltaMerge.Creature(delta)
		return true, nil
	case "plant":
		tileImpl.plantImpl = galapagosDeltaMerge.Plant(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []galapagos.Tile {
	neighbors := []galapagos.Tile{}

	if tileImpl.tileNorthImpl != nil {
		neighbors = append(neighbors, tileImpl.tileNorthImpl)
	}

	if tileImpl.tileEastImpl != nil {
		neighbors = append(neighbors, tileImpl.tileEastImpl)
	}

	if tileImpl.tileSouthImpl != nil {
		neighbors = append(neighbors, tileImpl.tileSouthImpl)
	}

	if tileImpl.tileWestImpl != nil {
		neighbors = append(neighbors, tileImpl.tileWestImpl)
	}

	return neighbors
}

// IsPathable returns if the Tile is pathable for FindPath
func (tileImpl *TileImpl) IsPathable() bool {
	return tileImpl.creature == nil && tileImpl.plant == nil
}

// HasNeighbor checks if this Tile has a specific neighboring Tile.
func (tileImpl *TileImpl) HasNeighbor(tile galapagos.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

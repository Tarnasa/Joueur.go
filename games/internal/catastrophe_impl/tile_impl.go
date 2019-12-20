package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/catastrophe"
)

// TileImpl is the struct that implements the container for Tile instances
// in Catastrophe.
type TileImpl struct {
	GameObjectImpl

	foodImpl           int64
	harvestRateImpl    int64
	materialsImpl      int64
	structureImpl      catastrophe.Structure
	tileEastImpl       catastrophe.Tile
	tileNorthImpl      catastrophe.Tile
	tileSouthImpl      catastrophe.Tile
	tileWestImpl       catastrophe.Tile
	turnsToHarvestImpl int64
	unitImpl           catastrophe.Unit
	xImpl              int64
	yImpl              int64
}

// Food returns the number of food dropped on this Tile.
func (tileImpl *TileImpl) Food() int64 {
	return tileImpl.foodImpl
}

// HarvestRate returns the amount of food that can be harvested from this
// Tile per turn.
func (tileImpl *TileImpl) HarvestRate() int64 {
	return tileImpl.harvestRateImpl
}

// Materials returns the number of materials dropped on this Tile.
func (tileImpl *TileImpl) Materials() int64 {
	return tileImpl.materialsImpl
}

// Structure returns the Structure on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Structure() catastrophe.Structure {
	return tileImpl.structureImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() catastrophe.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() catastrophe.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() catastrophe.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() catastrophe.Tile {
	return tileImpl.tileWestImpl
}

// TurnsToHarvest returns the amount of turns before this resource can be
// harvested.
func (tileImpl *TileImpl) TurnsToHarvest() int64 {
	return tileImpl.turnsToHarvestImpl
}

// Unit returns the Unit on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Unit() catastrophe.Unit {
	return tileImpl.unitImpl
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

	tileImpl.foodImpl = 0
	tileImpl.harvestRateImpl = 0
	tileImpl.materialsImpl = 0
	tileImpl.structureImpl = nil
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.turnsToHarvestImpl = 0
	tileImpl.unitImpl = nil
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

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "food":
		tileImpl.foodImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "harvestRate":
		tileImpl.harvestRateImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "materials":
		tileImpl.materialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "structure":
		tileImpl.structureImpl = catastropheDeltaMerge.Structure(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "turnsToHarvest":
		tileImpl.turnsToHarvestImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "unit":
		tileImpl.unitImpl = catastropheDeltaMerge.Unit(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []catastrophe.Tile {
	neighbors := []catastrophe.Tile{}

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
	// <<-- Creer-Merge: is-pathable -->>
	return false // TODO: developer add game logic here!
	// <<-- /Creer-Merge: is-pathable -->>
}

// HasNeighbor checks if this Tile has a specific neighboring Tile.
func (tileImpl *TileImpl) HasNeighbor(tile catastrophe.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

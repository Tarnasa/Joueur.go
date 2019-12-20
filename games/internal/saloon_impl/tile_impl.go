package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// TileImpl is the struct that implements the container for Tile instances
// in Saloon.
type TileImpl struct {
	GameObjectImpl

	bottleImpl     saloon.Bottle
	cowboyImpl     saloon.Cowboy
	furnishingImpl saloon.Furnishing
	hasHazardImpl  bool
	isBalconyImpl  bool
	tileEastImpl   saloon.Tile
	tileNorthImpl  saloon.Tile
	tileSouthImpl  saloon.Tile
	tileWestImpl   saloon.Tile
	xImpl          int64
	yImpl          int64
	youngGunImpl   saloon.YoungGun
}

// Bottle returns the beer Bottle currently flying over this Tile, nil
// otherwise.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Bottle() saloon.Bottle {
	return tileImpl.bottleImpl
}

// Cowboy returns the Cowboy that is on this Tile, nil otherwise.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Cowboy() saloon.Cowboy {
	return tileImpl.cowboyImpl
}

// Furnishing returns the furnishing that is on this Tile, nil otherwise.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Furnishing() saloon.Furnishing {
	return tileImpl.furnishingImpl
}

// HasHazard returns if this Tile is pathable, but has a hazard that
// damages Cowboys that path through it.
func (tileImpl *TileImpl) HasHazard() bool {
	return tileImpl.hasHazardImpl
}

// IsBalcony returns if this Tile is a balcony of the Saloon that YoungGuns
// walk around on, and can never be pathed through by Cowboys.
func (tileImpl *TileImpl) IsBalcony() bool {
	return tileImpl.isBalconyImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() saloon.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() saloon.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() saloon.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() saloon.Tile {
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

// YoungGun returns the YoungGun on this tile, nil otherwise.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) YoungGun() saloon.YoungGun {
	return tileImpl.youngGunImpl
}

// InitImplDefaults initializes safe defaults for all fields in Tile.
func (tileImpl *TileImpl) InitImplDefaults() {
	tileImpl.GameObjectImpl.InitImplDefaults()

	tileImpl.bottleImpl = nil
	tileImpl.cowboyImpl = nil
	tileImpl.furnishingImpl = nil
	tileImpl.hasHazardImpl = true
	tileImpl.isBalconyImpl = true
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.xImpl = 0
	tileImpl.yImpl = 0
	tileImpl.youngGunImpl = nil
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

	saloonDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bottle":
		tileImpl.bottleImpl = saloonDeltaMerge.Bottle(delta)
		return true, nil
	case "cowboy":
		tileImpl.cowboyImpl = saloonDeltaMerge.Cowboy(delta)
		return true, nil
	case "furnishing":
		tileImpl.furnishingImpl = saloonDeltaMerge.Furnishing(delta)
		return true, nil
	case "hasHazard":
		tileImpl.hasHazardImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "isBalcony":
		tileImpl.isBalconyImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "youngGun":
		tileImpl.youngGunImpl = saloonDeltaMerge.YoungGun(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []saloon.Tile {
	neighbors := []saloon.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile saloon.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

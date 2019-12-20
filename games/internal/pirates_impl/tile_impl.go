package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// TileImpl is the struct that implements the container for Tile instances
// in Pirates.
type TileImpl struct {
	GameObjectImpl

	decorationImpl bool
	goldImpl       int64
	portImpl       pirates.Port
	tileEastImpl   pirates.Tile
	tileNorthImpl  pirates.Tile
	tileSouthImpl  pirates.Tile
	tileWestImpl   pirates.Tile
	typeImpl       string
	unitImpl       pirates.Unit
	xImpl          int64
	yImpl          int64
}

// Decoration returns (Visualizer only) Whether this tile is deep sea or
// grassy. This has no effect on gameplay, but feel free to use it if you
// want.
func (tileImpl *TileImpl) Decoration() bool {
	return tileImpl.decorationImpl
}

// Gold returns the amount of gold buried on this tile.
func (tileImpl *TileImpl) Gold() int64 {
	return tileImpl.goldImpl
}

// Port returns the Port on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Port() pirates.Port {
	return tileImpl.portImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() pirates.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() pirates.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() pirates.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() pirates.Tile {
	return tileImpl.tileWestImpl
}

// Type returns the type of Tile this is ('water' or 'land').
//
// Literal Values: "water" or "land"
func (tileImpl *TileImpl) Type() string {
	return tileImpl.typeImpl
}

// Unit returns the Unit on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Unit() pirates.Unit {
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

	tileImpl.decorationImpl = true
	tileImpl.goldImpl = 0
	tileImpl.portImpl = nil
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.typeImpl = ""
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

	piratesDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'pirates.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "decoration":
		tileImpl.decorationImpl = piratesDeltaMerge.Boolean(delta)
		return true, nil
	case "gold":
		tileImpl.goldImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "port":
		tileImpl.portImpl = piratesDeltaMerge.Port(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	case "type":
		tileImpl.typeImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "unit":
		tileImpl.unitImpl = piratesDeltaMerge.Unit(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []pirates.Tile {
	neighbors := []pirates.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile pirates.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

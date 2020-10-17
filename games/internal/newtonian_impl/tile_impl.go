package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// TileImpl is the struct that implements the container for Tile instances
// in Newtonian.
type TileImpl struct {
	GameObjectImpl

	blueiumImpl    int64
	blueiumOreImpl int64
	decorationImpl int64
	directionImpl  string
	isWallImpl     bool
	machineImpl    newtonian.Machine
	ownerImpl      newtonian.Player
	rediumImpl     int64
	rediumOreImpl  int64
	tileEastImpl   newtonian.Tile
	tileNorthImpl  newtonian.Tile
	tileSouthImpl  newtonian.Tile
	tileWestImpl   newtonian.Tile
	typeImpl       string
	unitImpl       newtonian.Unit
	xImpl          int64
	yImpl          int64
}

// Blueium returns the amount of blueium on this tile.
func (tileImpl *TileImpl) Blueium() int64 {
	return tileImpl.blueiumImpl
}

// BlueiumOre returns the amount of blueium ore on this tile.
func (tileImpl *TileImpl) BlueiumOre() int64 {
	return tileImpl.blueiumOreImpl
}

// Decoration returns (Visualizer only) Different tile types, cracked,
// slightly dirty, etc. This has no effect on gameplay, but feel free to
// use it if you want.
func (tileImpl *TileImpl) Decoration() int64 {
	return tileImpl.decorationImpl
}

// Direction returns the direction of a conveyor belt ('blank', 'north',
// 'east', 'south', or 'west'). Blank means conveyor doesn't move.
//
// Literal Values: "blank", "north", "east", "south", or "west"
func (tileImpl *TileImpl) Direction() string {
	return tileImpl.directionImpl
}

// IsWall returns whether or not the tile is a wall.
func (tileImpl *TileImpl) IsWall() bool {
	return tileImpl.isWallImpl
}

// Machine returns the Machine on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Machine() newtonian.Machine {
	return tileImpl.machineImpl
}

// Owner returns the owner of this Tile, or nil if owned by no-one. Only
// for generators and spawn areas.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Owner() newtonian.Player {
	return tileImpl.ownerImpl
}

// Redium returns the amount of redium on this tile.
func (tileImpl *TileImpl) Redium() int64 {
	return tileImpl.rediumImpl
}

// RediumOre returns the amount of redium ore on this tile.
func (tileImpl *TileImpl) RediumOre() int64 {
	return tileImpl.rediumOreImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() newtonian.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() newtonian.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() newtonian.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() newtonian.Tile {
	return tileImpl.tileWestImpl
}

// Type returns the type of Tile this is ('normal', 'generator',
// 'conveyor', or 'spawn').
//
// Literal Values: "normal", "generator", "conveyor", or "spawn"
func (tileImpl *TileImpl) Type() string {
	return tileImpl.typeImpl
}

// Unit returns the Unit on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Unit() newtonian.Unit {
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

	tileImpl.blueiumImpl = 0
	tileImpl.blueiumOreImpl = 0
	tileImpl.decorationImpl = 0
	tileImpl.directionImpl = ""
	tileImpl.isWallImpl = true
	tileImpl.machineImpl = nil
	tileImpl.ownerImpl = nil
	tileImpl.rediumImpl = 0
	tileImpl.rediumOreImpl = 0
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

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "blueium":
		tileImpl.blueiumImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "blueiumOre":
		tileImpl.blueiumOreImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "decoration":
		tileImpl.decorationImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "direction":
		tileImpl.directionImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "isWall":
		tileImpl.isWallImpl = newtonianDeltaMerge.Boolean(delta)
		return true, nil
	case "machine":
		tileImpl.machineImpl = newtonianDeltaMerge.Machine(delta)
		return true, nil
	case "owner":
		tileImpl.ownerImpl = newtonianDeltaMerge.Player(delta)
		return true, nil
	case "redium":
		tileImpl.rediumImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "rediumOre":
		tileImpl.rediumOreImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	case "type":
		tileImpl.typeImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "unit":
		tileImpl.unitImpl = newtonianDeltaMerge.Unit(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []newtonian.Tile {
	neighbors := []newtonian.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile newtonian.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// TileImpl is the struct that implements the container for Tile instances
// in Coreminer.
type TileImpl struct {
	GameObjectImpl

	bombsImpl     []coreminer.Bomb
	dirtImpl      int64
	isBaseImpl    bool
	isFallingImpl bool
	isHopperImpl  bool
	isLadderImpl  bool
	isSupportImpl bool
	minersImpl    []coreminer.Miner
	oreImpl       int64
	ownerImpl     coreminer.Player
	shieldingImpl int64
	tileEastImpl  coreminer.Tile
	tileNorthImpl coreminer.Tile
	tileSouthImpl coreminer.Tile
	tileWestImpl  coreminer.Tile
	xImpl         int64
	yImpl         int64
}

// Bombs returns an array of Bombs on this Tile.
func (tileImpl *TileImpl) Bombs() []coreminer.Bomb {
	return tileImpl.bombsImpl
}

// Dirt returns the amount of dirt on this Tile.
func (tileImpl *TileImpl) Dirt() int64 {
	return tileImpl.dirtImpl
}

// IsBase returns whether or not the Tile is a base Tile.
func (tileImpl *TileImpl) IsBase() bool {
	return tileImpl.isBaseImpl
}

// IsFalling returns whether or not this Tile is about to fall after this
// turn.
func (tileImpl *TileImpl) IsFalling() bool {
	return tileImpl.isFallingImpl
}

// IsHopper returns whether or not a hopper is on this Tile.
func (tileImpl *TileImpl) IsHopper() bool {
	return tileImpl.isHopperImpl
}

// IsLadder returns whether or not a ladder is built on this Tile.
func (tileImpl *TileImpl) IsLadder() bool {
	return tileImpl.isLadderImpl
}

// IsSupport returns whether or not a support is built on this Tile.
func (tileImpl *TileImpl) IsSupport() bool {
	return tileImpl.isSupportImpl
}

// Miners returns an array of the Miners on this Tile.
func (tileImpl *TileImpl) Miners() []coreminer.Miner {
	return tileImpl.minersImpl
}

// Ore returns the amount of ore on this Tile.
func (tileImpl *TileImpl) Ore() int64 {
	return tileImpl.oreImpl
}

// Owner returns the owner of this Tile, or undefined if owned by no-one.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Owner() coreminer.Player {
	return tileImpl.ownerImpl
}

// Shielding returns the amount of shielding on this Tile.
func (tileImpl *TileImpl) Shielding() int64 {
	return tileImpl.shieldingImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() coreminer.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() coreminer.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() coreminer.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() coreminer.Tile {
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

	tileImpl.bombsImpl = []coreminer.Bomb{}
	tileImpl.dirtImpl = 0
	tileImpl.isBaseImpl = true
	tileImpl.isFallingImpl = true
	tileImpl.isHopperImpl = true
	tileImpl.isLadderImpl = true
	tileImpl.isSupportImpl = true
	tileImpl.minersImpl = []coreminer.Miner{}
	tileImpl.oreImpl = 0
	tileImpl.ownerImpl = nil
	tileImpl.shieldingImpl = 0
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

	coreminerDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'coreminer.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bombs":
		tileImpl.bombsImpl = coreminerDeltaMerge.ArrayOfBomb(&tileImpl.bombsImpl, delta)
		return true, nil
	case "dirt":
		tileImpl.dirtImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "isBase":
		tileImpl.isBaseImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "isFalling":
		tileImpl.isFallingImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "isHopper":
		tileImpl.isHopperImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "isLadder":
		tileImpl.isLadderImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "isSupport":
		tileImpl.isSupportImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "miners":
		tileImpl.minersImpl = coreminerDeltaMerge.ArrayOfMiner(&tileImpl.minersImpl, delta)
		return true, nil
	case "ore":
		tileImpl.oreImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		tileImpl.ownerImpl = coreminerDeltaMerge.Player(delta)
		return true, nil
	case "shielding":
		tileImpl.shieldingImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []coreminer.Tile {
	neighbors := []coreminer.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile coreminer.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

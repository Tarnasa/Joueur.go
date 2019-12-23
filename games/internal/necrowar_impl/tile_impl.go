package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// TileImpl is the struct that implements the container for Tile instances
// in Necrowar.
type TileImpl struct {
	GameObjectImpl

	corpsesImpl          int64
	isCastleImpl         bool
	isGoldMineImpl       bool
	isGrassImpl          bool
	isIslandGoldMineImpl bool
	isPathImpl           bool
	isRiverImpl          bool
	isTowerImpl          bool
	isUnitSpawnImpl      bool
	isWallImpl           bool
	isWorkerSpawnImpl    bool
	numGhoulsImpl        int64
	numHoundsImpl        int64
	numZombiesImpl       int64
	ownerImpl            necrowar.Player
	tileEastImpl         necrowar.Tile
	tileNorthImpl        necrowar.Tile
	tileSouthImpl        necrowar.Tile
	tileWestImpl         necrowar.Tile
	towerImpl            necrowar.Tower
	unitImpl             necrowar.Unit
	xImpl                int64
	yImpl                int64
}

// Corpses returns the amount of corpses on this tile.
func (tileImpl *TileImpl) Corpses() int64 {
	return tileImpl.corpsesImpl
}

// IsCastle returns whether or not the tile is a castle tile.
func (tileImpl *TileImpl) IsCastle() bool {
	return tileImpl.isCastleImpl
}

// IsGoldMine returns whether or not the tile is considered to be a gold
// mine or not.
func (tileImpl *TileImpl) IsGoldMine() bool {
	return tileImpl.isGoldMineImpl
}

// IsGrass returns whether or not the tile is considered grass or not
// (Workers can walk on grass).
func (tileImpl *TileImpl) IsGrass() bool {
	return tileImpl.isGrassImpl
}

// IsIslandGoldMine returns whether or not the tile is considered to be the
// island gold mine or not.
func (tileImpl *TileImpl) IsIslandGoldMine() bool {
	return tileImpl.isIslandGoldMineImpl
}

// IsPath returns whether or not the tile is considered a path or not
// (Units can walk on paths).
func (tileImpl *TileImpl) IsPath() bool {
	return tileImpl.isPathImpl
}

// IsRiver returns whether or not the tile is considered a river or not.
func (tileImpl *TileImpl) IsRiver() bool {
	return tileImpl.isRiverImpl
}

// IsTower returns whether or not the tile is considered a tower or not.
func (tileImpl *TileImpl) IsTower() bool {
	return tileImpl.isTowerImpl
}

// IsUnitSpawn returns whether or not the tile is the unit spawn.
func (tileImpl *TileImpl) IsUnitSpawn() bool {
	return tileImpl.isUnitSpawnImpl
}

// IsWall returns whether or not the tile can be moved on by workers.
func (tileImpl *TileImpl) IsWall() bool {
	return tileImpl.isWallImpl
}

// IsWorkerSpawn returns whether or not the tile is the worker spawn.
func (tileImpl *TileImpl) IsWorkerSpawn() bool {
	return tileImpl.isWorkerSpawnImpl
}

// NumGhouls returns the amount of Ghouls on this tile.
func (tileImpl *TileImpl) NumGhouls() int64 {
	return tileImpl.numGhoulsImpl
}

// NumHounds returns the amount of Hounds on this tile.
func (tileImpl *TileImpl) NumHounds() int64 {
	return tileImpl.numHoundsImpl
}

// NumZombies returns the amount of Zombies on this tile.
func (tileImpl *TileImpl) NumZombies() int64 {
	return tileImpl.numZombiesImpl
}

// Owner returns which player owns this tile, only applies to grass tiles
// for workers, NULL otherwise.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Owner() necrowar.Player {
	return tileImpl.ownerImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() necrowar.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() necrowar.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() necrowar.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() necrowar.Tile {
	return tileImpl.tileWestImpl
}

// Tower returns the Tower on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Tower() necrowar.Tower {
	return tileImpl.towerImpl
}

// Unit returns the Unit on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Unit() necrowar.Unit {
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

// Res runs logic that resurrect the corpses on this tile into Zombies.
func (tileImpl *TileImpl) Res(num int64) bool {
	return tileImpl.RunOnServer("res", map[string]interface{}{
		"num": num,
	}).(bool)
}

// SpawnUnit runs logic that spawns a fighting unit on the correct tile.
func (tileImpl *TileImpl) SpawnUnit(title string) bool {
	return tileImpl.RunOnServer("spawnUnit", map[string]interface{}{
		"title": title,
	}).(bool)
}

// SpawnWorker runs logic that spawns a worker on the correct tile.
func (tileImpl *TileImpl) SpawnWorker() bool {
	return tileImpl.RunOnServer("spawnWorker", map[string]interface{}{}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Tile.
func (tileImpl *TileImpl) InitImplDefaults() {
	tileImpl.GameObjectImpl.InitImplDefaults()

	tileImpl.corpsesImpl = 0
	tileImpl.isCastleImpl = true
	tileImpl.isGoldMineImpl = true
	tileImpl.isGrassImpl = true
	tileImpl.isIslandGoldMineImpl = true
	tileImpl.isPathImpl = true
	tileImpl.isRiverImpl = true
	tileImpl.isTowerImpl = true
	tileImpl.isUnitSpawnImpl = true
	tileImpl.isWallImpl = true
	tileImpl.isWorkerSpawnImpl = true
	tileImpl.numGhoulsImpl = 0
	tileImpl.numHoundsImpl = 0
	tileImpl.numZombiesImpl = 0
	tileImpl.ownerImpl = nil
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.towerImpl = nil
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

	necrowarDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'necrowar.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "corpses":
		tileImpl.corpsesImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "isCastle":
		tileImpl.isCastleImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isGoldMine":
		tileImpl.isGoldMineImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isGrass":
		tileImpl.isGrassImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isIslandGoldMine":
		tileImpl.isIslandGoldMineImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isPath":
		tileImpl.isPathImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isRiver":
		tileImpl.isRiverImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isTower":
		tileImpl.isTowerImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isUnitSpawn":
		tileImpl.isUnitSpawnImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isWall":
		tileImpl.isWallImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "isWorkerSpawn":
		tileImpl.isWorkerSpawnImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "numGhouls":
		tileImpl.numGhoulsImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "numHounds":
		tileImpl.numHoundsImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "numZombies":
		tileImpl.numZombiesImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		tileImpl.ownerImpl = necrowarDeltaMerge.Player(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	case "tower":
		tileImpl.towerImpl = necrowarDeltaMerge.Tower(delta)
		return true, nil
	case "unit":
		tileImpl.unitImpl = necrowarDeltaMerge.Unit(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []necrowar.Tile {
	neighbors := []necrowar.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile necrowar.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

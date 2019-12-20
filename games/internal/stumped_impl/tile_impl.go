package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// TileImpl is the struct that implements the container for Tile instances
// in Stumped.
type TileImpl struct {
	GameObjectImpl

	beaverImpl        stumped.Beaver
	branchesImpl      int64
	flowDirectionImpl string
	foodImpl          int64
	lodgeOwnerImpl    stumped.Player
	spawnerImpl       stumped.Spawner
	tileEastImpl      stumped.Tile
	tileNorthImpl     stumped.Tile
	tileSouthImpl     stumped.Tile
	tileWestImpl      stumped.Tile
	typeImpl          string
	xImpl             int64
	yImpl             int64
}

// Beaver returns the Beaver on this Tile if present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Beaver() stumped.Beaver {
	return tileImpl.beaverImpl
}

// Branches returns the number of branches dropped on this Tile.
func (tileImpl *TileImpl) Branches() int64 {
	return tileImpl.branchesImpl
}

// FlowDirection returns the cardinal direction water is flowing on this
// Tile ('North', 'East', 'South', 'West').
//
// Literal Values: "North", "East", "South", "West", or ""
func (tileImpl *TileImpl) FlowDirection() string {
	return tileImpl.flowDirectionImpl
}

// Food returns the number of food dropped on this Tile.
func (tileImpl *TileImpl) Food() int64 {
	return tileImpl.foodImpl
}

// LodgeOwner returns the owner of the Beaver lodge on this Tile, if
// present, otherwise nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) LodgeOwner() stumped.Player {
	return tileImpl.lodgeOwnerImpl
}

// Spawner returns the resource Spawner on this Tile if present, otherwise
// nil.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) Spawner() stumped.Spawner {
	return tileImpl.spawnerImpl
}

// TileEast returns the Tile to the 'East' of this one (x+1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileEast() stumped.Tile {
	return tileImpl.tileEastImpl
}

// TileNorth returns the Tile to the 'North' of this one (x, y-1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileNorth() stumped.Tile {
	return tileImpl.tileNorthImpl
}

// TileSouth returns the Tile to the 'South' of this one (x, y+1). Nil if
// out of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileSouth() stumped.Tile {
	return tileImpl.tileSouthImpl
}

// TileWest returns the Tile to the 'West' of this one (x-1, y). Nil if out
// of bounds of the map.
//
// Value can be returned as a nil pointer.
func (tileImpl *TileImpl) TileWest() stumped.Tile {
	return tileImpl.tileWestImpl
}

// Type returns what type of Tile this is, either 'water' or 'land'.
//
// Literal Values: "land" or "water"
func (tileImpl *TileImpl) Type() string {
	return tileImpl.typeImpl
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

	tileImpl.beaverImpl = nil
	tileImpl.branchesImpl = 0
	tileImpl.flowDirectionImpl = ""
	tileImpl.foodImpl = 0
	tileImpl.lodgeOwnerImpl = nil
	tileImpl.spawnerImpl = nil
	tileImpl.tileEastImpl = nil
	tileImpl.tileNorthImpl = nil
	tileImpl.tileSouthImpl = nil
	tileImpl.tileWestImpl = nil
	tileImpl.typeImpl = ""
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

	stumpedDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'stumped.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "beaver":
		tileImpl.beaverImpl = stumpedDeltaMerge.Beaver(delta)
		return true, nil
	case "branches":
		tileImpl.branchesImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "flowDirection":
		tileImpl.flowDirectionImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "food":
		tileImpl.foodImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "lodgeOwner":
		tileImpl.lodgeOwnerImpl = stumpedDeltaMerge.Player(delta)
		return true, nil
	case "spawner":
		tileImpl.spawnerImpl = stumpedDeltaMerge.Spawner(delta)
		return true, nil
	case "tileEast":
		tileImpl.tileEastImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "tileNorth":
		tileImpl.tileNorthImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "tileSouth":
		tileImpl.tileSouthImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "tileWest":
		tileImpl.tileWestImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "type":
		tileImpl.typeImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "x":
		tileImpl.xImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "y":
		tileImpl.yImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetNeighbors returns an array of the neighbors of this Tile.
func (tileImpl *TileImpl) GetNeighbors() []stumped.Tile {
	neighbors := []stumped.Tile{}

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
func (tileImpl *TileImpl) HasNeighbor(tile stumped.Tile) bool {
	return tile != nil &&
		(tileImpl.tileNorthImpl == tile ||
			tileImpl.tileEastImpl == tile ||
			tileImpl.tileSouthImpl == tile ||
			tileImpl.tileWestImpl == tile)
}

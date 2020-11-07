package coreminer

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Bombs is an array of Bombs on this Tile.
	Bombs() []Bomb

	// Dirt is the amount of dirt on this Tile.
	Dirt() int64

	// IsBase is whether or not the Tile is a base Tile.
	IsBase() bool

	// IsFalling is whether or not this Tile is about to fall after
	// this turn.
	IsFalling() bool

	// IsHopper is whether or not a hopper is on this Tile.
	IsHopper() bool

	// IsLadder is whether or not a ladder is built on this Tile.
	IsLadder() bool

	// IsSupport is whether or not a support is built on this Tile.
	IsSupport() bool

	// Miners is an array of the Miners on this Tile.
	Miners() []Miner

	// Ore is the amount of ore on this Tile.
	Ore() int64

	// Owner is the owner of this Tile, or undefined if owned by no-
	// one.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Shielding is the amount of shielding on this Tile.
	Shielding() int64

	// TileEast is the Tile to the 'East' of this one (x+1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileEast() Tile

	// TileNorth is the Tile to the 'North' of this one (x, y-1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileNorth() Tile

	// TileSouth is the Tile to the 'South' of this one (x, y+1). Nil
	// if out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileSouth() Tile

	// TileWest is the Tile to the 'West' of this one (x-1, y). Nil if
	// out of bounds of the map.
	//
	// Value can be returned as a nil pointer.
	TileWest() Tile

	// X is the x (horizontal) position of this Tile.
	X() int64

	// Y is the y (vertical) position of this Tile.
	Y() int64

	// -- Tiled Game Utils -- \\

	// GetNeighbors returns an array of the neighbors of this Tile.
	GetNeighbors() []Tile

	// IsPathable returns if the Tile is pathable for FindPath
	IsPathable() bool

	// HasNeighbor checks if this Tile has a specific neighboring Tile.
	HasNeighbor(Tile) bool
}

// TileDirections are all the direction strings that Tile's neighbors can be
// in.
var TileDirections = [...]string{"North", "South", "East", "West"}

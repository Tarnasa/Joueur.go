package newtonian

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Blueium is the amount of blueium on this tile.
	Blueium() int64

	// BlueiumOre is the amount of blueium ore on this tile.
	BlueiumOre() int64

	// Decoration is (Visualizer only) Different tile types, cracked,
	// slightly dirty, etc. This has no effect on gameplay, but feel
	// free to use it if you want.
	Decoration() int64

	// Direction is the direction of a conveyor belt ('blank',
	// 'north', 'east', 'south', or 'west'). blank means conveyor
	// doesn't move.
	//
	// Literal Values: "blank", "north", "east", "south", or "west"
	Direction() string

	// IsWall is whether or not the tile is a wall.
	IsWall() bool

	// Machine is the Machine on this Tile if present, otherwise nil.
	//
	// Value can be returned as a nil pointer.
	Machine() Machine

	// Owner is the owner of this Tile, or nil if owned by no-one.
	// Only for generators and spawn areas.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Redium is the amount of redium on this tile.
	Redium() int64

	// RediumOre is the amount of redium ore on this tile.
	RediumOre() int64

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

	// Type is the type of Tile this is ('normal', 'generator',
	// 'conveyor', or 'spawn').
	//
	// Literal Values: "normal", "generator", "conveyor", or "spawn"
	Type() string

	// Unit is the Unit on this Tile if present, otherwise nil.
	//
	// Value can be returned as a nil pointer.
	Unit() Unit

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

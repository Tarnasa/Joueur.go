package stumped

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Beaver is the Beaver on this Tile if present, otherwise nil.
//
	// Value can be returned as a nil pointer.
	Beaver() Beaver

	// Branches is the number of branches dropped on this Tile.
	Branches() int64

	// FlowDirection is the cardinal direction water is flowing on
	// this Tile ('North', 'East', 'South', 'West').
//
	// Literal Values: "North", "East", "South", "West", or ""
	FlowDirection() string

	// Food is the number of food dropped on this Tile.
	Food() int64

	// LodgeOwner is the owner of the Beaver lodge on this Tile, if
	// present, otherwise nil.
//
	// Value can be returned as a nil pointer.
	LodgeOwner() Player

	// Spawner is the resource Spawner on this Tile if present,
	// otherwise nil.
//
	// Value can be returned as a nil pointer.
	Spawner() Spawner

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

	// Type is what type of Tile this is, either 'water' or 'land'.
//
	// Literal Values: "land" or "water"
	Type() string

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

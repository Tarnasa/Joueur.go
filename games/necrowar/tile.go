package necrowar

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Corpses is the amount of corpses on this tile.
	Corpses() int64

	// IsCastle is whether or not the tile is a castle tile.
	IsCastle() bool

	// IsGoldMine is whether or not the tile is considered to be a
	// gold mine or not.
	IsGoldMine() bool

	// IsGrass is whether or not the tile is considered grass or not
	// (Workers can walk on grass).
	IsGrass() bool

	// IsIslandGoldMine is whether or not the tile is considered to be
	// the island gold mine or not.
	IsIslandGoldMine() bool

	// IsPath is whether or not the tile is considered a path or not
	// (Units can walk on paths).
	IsPath() bool

	// IsRiver is whether or not the tile is considered a river or
	// not.
	IsRiver() bool

	// IsTower is whether or not the tile is considered a tower or
	// not.
	IsTower() bool

	// IsUnitSpawn is whether or not the tile is the unit spawn.
	IsUnitSpawn() bool

	// IsWall is whether or not the tile can be moved on by workers.
	IsWall() bool

	// IsWorkerSpawn is whether or not the tile is the worker spawn.
	IsWorkerSpawn() bool

	// NumGhouls is the amount of Ghouls on this tile.
	NumGhouls() int64

	// NumHounds is the amount of Hounds on this tile.
	NumHounds() int64

	// NumZombies is the amount of Zombies on this tile.
	NumZombies() int64

	// Owner is which player owns this tile, only applies to grass
	// tiles for workers, NULL otherwise.
//
	// Value can be returned as a nil pointer.
	Owner() Player

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

	// Tower is the Tower on this Tile if present, otherwise nil.
//
	// Value can be returned as a nil pointer.
	Tower() Tower

	// Unit is the Unit on this Tile if present, otherwise nil.
//
	// Value can be returned as a nil pointer.
	Unit() Unit

	// X is the x (horizontal) position of this Tile.
	X() int64

	// Y is the y (vertical) position of this Tile.
	Y() int64

	// -- Methods -- \\

	// Res resurrect the corpses on this tile into Zombies.
	Res(int64) bool

	// SpawnUnit spawns a fighting unit on the correct tile.
	SpawnUnit(string) bool

	// SpawnWorker spawns a worker on the correct tile.
	SpawnWorker() bool

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

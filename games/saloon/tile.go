package saloon

// Tile is a Tile in the game that makes up the 2D map grid.
type Tile interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Bottle is the beer Bottle currently flying over this Tile, nil
	// otherwise.
	Bottle() Bottle

	// Cowboy is the Cowboy that is on this Tile, nil otherwise.
	Cowboy() Cowboy

	// Furnishing is the furnishing that is on this Tile, nil
	// otherwise.
	Furnishing() Furnishing

	// HasHazard is if this Tile is pathable, but has a hazard that
	// damages Cowboys that path through it.
	HasHazard() bool

	// IsBalcony is if this Tile is a balcony of the Saloon that
	// YoungGuns walk around on, and can never be pathed through by
	// Cowboys.
	IsBalcony() bool

	// TileEast is the Tile to the 'East' of this one (x+1, y). Nil if
	// out of bounds of the map.
	TileEast() Tile

	// TileNorth is the Tile to the 'North' of this one (x, y-1). Nil
	// if out of bounds of the map.
	TileNorth() Tile

	// TileSouth is the Tile to the 'South' of this one (x, y+1). Nil
	// if out of bounds of the map.
	TileSouth() Tile

	// TileWest is the Tile to the 'West' of this one (x-1, y). Nil if
	// out of bounds of the map.
	TileWest() Tile

	// X is the x (horizontal) position of this Tile.
	X() int64

	// Y is the y (vertical) position of this Tile.
	Y() int64

	// YoungGun is the YoungGun on this tile, nil otherwise.
	YoungGun() YoungGun
}

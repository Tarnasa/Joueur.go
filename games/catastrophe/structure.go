package catastrophe

// Structure is a structure on a Tile.
type Structure interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// EffectRadius is the range of this Structure's effect. For
	// example, a radius of 1 means this Structure affects a 3x3
	// square centered on this Structure.
	EffectRadius() int64

	// Materials is the number of materials in this Structure. Once
	// this number reaches 0, this Structure is destroyed.
	Materials() int64

	// Owner is the owner of this Structure if any, otherwise nil.
//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Tile is the Tile this Structure is on.
//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// Type is the type of Structure this is ('shelter', 'monument',
	// 'wall', 'road', 'neutral').
//
	// Literal Values: "neutral", "shelter", "monument", "wall", or
	// "road"
	Type() string
}

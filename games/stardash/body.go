package stardash

// Body is a celestial body located within the game.
type Body interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// Amount is the amount of material the object has, or energy if
	// it is a planet.
	Amount() int64

	// BodyType is the type of celestial body it is. Either 'planet',
	// 'asteroid', or 'sun'.
	//
	// Literal Values: "planet", "asteroid", or "sun"
	BodyType() string

	// MaterialType is the type of material the celestial body has.
	// Either 'none', 'genarium', 'rarium', 'legendarium', or
	// 'mythicite'.
	//
	// Literal Values: "none", "genarium", "rarium", "legendarium", or
	// "mythicite"
	MaterialType() string

	// Owner is the Player that owns and can control this Body.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Radius is the radius of the circle that this body takes up.
	Radius() float64

	// X is the x value this celestial body is on.
	X() float64

	// Y is the y value this celestial body is on.
	Y() float64

	// -- Methods -- \\

	// NextX the x value of this body a number of turns from now.
	// (0-how many you want).
	NextX(int64) int64

	// NextY the x value of this body a number of turns from now.
	// (0-how many you want).
	NextY(int64) int64

	// Spawn spawn a unit on some value of this celestial body.
	Spawn(float64, float64, string) bool
}

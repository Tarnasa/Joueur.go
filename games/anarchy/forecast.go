package anarchy

// Forecast is the weather effect that will be applied at the end of a
// turn, which causes fires to spread.
type Forecast interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// ControllingPlayer is the Player that can use WeatherStations to
	// control this Forecast when its the nextForecast.
	ControllingPlayer() Player

	// Direction is the direction the wind will blow fires in. Can be
	// 'north', 'east', 'south', or 'west'.
//
	// Literal Values: "North", "East", "South", or "West"
	Direction() string

	// Intensity is how much of a Building's fire that can be blown in
	// the direction of this Forecast. Fire is duplicated (copied),
	// not moved (transfered).
	Intensity() int64
}

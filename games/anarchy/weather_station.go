package anarchy

// WeatherStation is can be bribed to change the next Forecast in some way.
type WeatherStation interface {
	// Parent interfaces
	Building

	// -- Methods -- \\

	// Intensify bribe the weathermen to intensity the next Forecast
	// by 1 or -1
	Intensify(bool) bool

	// Rotate bribe the weathermen to change the direction of the next
	// Forecast by rotating it clockwise or counterclockwise.
	Rotate(bool) bool
}

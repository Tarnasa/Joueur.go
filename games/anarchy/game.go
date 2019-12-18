package anarchy

import "joueur/base"

// Game is two player grid based game where each player tries to burn down
// the other player's buildings. Let it burn.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// BaseBribesPerTurn is how many bribes players get at the
	// beginning of their turn, not counting their burned down
	// Buildings.
	BaseBribesPerTurn() int64

	// Buildings is all the buildings in the game.
	Buildings() []Building

	// CurrentForecast is the current Forecast, which will be applied
	// at the end of the turn.
	CurrentForecast() Forecast

	// CurrentPlayer is the player whose turn it is currently. That
	// player can send commands. Other players cannot.
	CurrentPlayer() Player

	// CurrentTurn is the current turn number, starting at 0 for the
	// first player's turn.
	CurrentTurn() int64

	// Forecasts is all the forecasts in the game, indexed by turn
	// number.
	Forecasts() []Forecast

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// MapHeight is the width of the entire map along the vertical (y)
	// axis.
	MapHeight() int64

	// MapWidth is the width of the entire map along the horizontal
	// (x) axis.
	MapWidth() int64

	// MaxFire is the maximum amount of fire value for any Building.
	MaxFire() int64

	// MaxForecastIntensity is the maximum amount of intensity value
	// for any Forecast.
	MaxForecastIntensity() int64

	// MaxTurns is the maximum number of turns before the game will
	// automatically end.
	MaxTurns() int64

	// NextForecast is the next Forecast, which will be applied at the
	// end of your opponent's turn. This is also the Forecast
	// WeatherStations can control this turn.
	NextForecast() Forecast

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string

	// TimeAddedPerTurn is the amount of time (in nano-seconds) added
	// after each player performs a turn.
	TimeAddedPerTurn() int64
}

package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// GameImpl is the struct that implements the container for Game instances
// in Anarchy.
type GameImpl struct {
	base.GameImpl

	baseBribesPerTurnImpl    int64
	buildingsImpl            []anarchy.Building
	currentForecastImpl      anarchy.Forecast
	currentPlayerImpl        anarchy.Player
	currentTurnImpl          int64
	forecastsImpl            []anarchy.Forecast
	gameObjectsImpl          map[string]anarchy.GameObject
	mapHeightImpl            int64
	mapWidthImpl             int64
	maxFireImpl              int64
	maxForecastIntensityImpl int64
	maxTurnsImpl             int64
	nextForecastImpl         anarchy.Forecast
	playersImpl              []anarchy.Player
	sessionImpl              string
	timeAddedPerTurnImpl     int64
}

// BaseBribesPerTurn returns how many bribes players get at the beginning
// of their turn, not counting their burned down Buildings.
func (gameImpl *GameImpl) BaseBribesPerTurn() int64 {
	return gameImpl.baseBribesPerTurnImpl
}

// Buildings returns all the buildings in the game.
func (gameImpl *GameImpl) Buildings() []anarchy.Building {
	return gameImpl.buildingsImpl
}

// CurrentForecast returns the current Forecast, which will be applied at
// the end of the turn.
func (gameImpl *GameImpl) CurrentForecast() anarchy.Forecast {
	return gameImpl.currentForecastImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() anarchy.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// Forecasts returns all the forecasts in the game, indexed by turn number.
func (gameImpl *GameImpl) Forecasts() []anarchy.Forecast {
	return gameImpl.forecastsImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]anarchy.GameObject {
	return gameImpl.gameObjectsImpl
}

// MapHeight returns the width of the entire map along the vertical (y)
// axis.
func (gameImpl *GameImpl) MapHeight() int64 {
	return gameImpl.mapHeightImpl
}

// MapWidth returns the width of the entire map along the horizontal (x)
// axis.
func (gameImpl *GameImpl) MapWidth() int64 {
	return gameImpl.mapWidthImpl
}

// MaxFire returns the maximum amount of fire value for any Building.
func (gameImpl *GameImpl) MaxFire() int64 {
	return gameImpl.maxFireImpl
}

// MaxForecastIntensity returns the maximum amount of intensity value for
// any Forecast.
func (gameImpl *GameImpl) MaxForecastIntensity() int64 {
	return gameImpl.maxForecastIntensityImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// NextForecast returns the next Forecast, which will be applied at the end
// of your opponent's turn. This is also the Forecast WeatherStations can
// control this turn.
//
// Value can be returned as a nil pointer.
func (gameImpl *GameImpl) NextForecast() anarchy.Forecast {
	return gameImpl.nextForecastImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []anarchy.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.baseBribesPerTurnImpl = 0
	gameImpl.buildingsImpl = []anarchy.Building{}
	gameImpl.currentForecastImpl = nil
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.forecastsImpl = []anarchy.Forecast{}
	gameImpl.gameObjectsImpl = make(map[string]anarchy.GameObject)
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxFireImpl = 0
	gameImpl.maxForecastIntensityImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.nextForecastImpl = nil
	gameImpl.playersImpl = []anarchy.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.timeAddedPerTurnImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Game.
func (gameImpl *GameImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := gameImpl.GameImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	anarchyDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'anarchy.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "baseBribesPerTurn":
		gameImpl.baseBribesPerTurnImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "buildings":
		gameImpl.buildingsImpl = anarchyDeltaMerge.ArrayOfBuilding(&gameImpl.buildingsImpl, delta)
		return true, nil
	case "currentForecast":
		gameImpl.currentForecastImpl = anarchyDeltaMerge.Forecast(delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = anarchyDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "forecasts":
		gameImpl.forecastsImpl = anarchyDeltaMerge.ArrayOfForecast(&gameImpl.forecastsImpl, delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = anarchyDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "maxFire":
		gameImpl.maxFireImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "maxForecastIntensity":
		gameImpl.maxForecastIntensityImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "nextForecast":
		gameImpl.nextForecastImpl = anarchyDeltaMerge.Forecast(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = anarchyDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

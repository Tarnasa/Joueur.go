package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// ForecastImpl is the struct that implements the container for Forecast
// instances in Anarchy.
type ForecastImpl struct {
	GameObjectImpl

	controllingPlayerImpl anarchy.Player
	directionImpl         string
	intensityImpl         int64
}

// ControllingPlayer returns the Player that can use WeatherStations to
// control this Forecast when its the nextForecast.
func (forecastImpl *ForecastImpl) ControllingPlayer() anarchy.Player {
	return forecastImpl.controllingPlayerImpl
}

// Direction returns the direction the wind will blow fires in. Can be
// 'north', 'east', 'south', or 'west'.
//
// Literal Values: "North", "East", "South", or "West"
func (forecastImpl *ForecastImpl) Direction() string {
	return forecastImpl.directionImpl
}

// Intensity returns how much of a Building's fire that can be blown in the
// direction of this Forecast. Fire is duplicated (copied), not moved
// (transfered).
func (forecastImpl *ForecastImpl) Intensity() int64 {
	return forecastImpl.intensityImpl
}

// InitImplDefaults initializes safe defaults for all fields in Forecast.
func (forecastImpl *ForecastImpl) InitImplDefaults() {
	forecastImpl.GameObjectImpl.InitImplDefaults()

	forecastImpl.controllingPlayerImpl = nil
	forecastImpl.directionImpl = ""
	forecastImpl.intensityImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Forecast.
func (forecastImpl *ForecastImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := forecastImpl.GameObjectImpl.DeltaMerge(
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
	case "controllingPlayer":
		forecastImpl.controllingPlayerImpl = anarchyDeltaMerge.Player(delta)
		return true, nil
	case "direction":
		forecastImpl.directionImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "intensity":
		forecastImpl.intensityImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl
import (
	"errors"
	"joueur/base"
)

// WeatherStationImpl is the struct that implements the container for
// WeatherStation instances in Anarchy.
type WeatherStationImpl struct {
	BuildingImpl

}

// Intensify runs logic that bribe the weathermen to intensity the next
// Forecast by 1 or -1
func (weatherStationImpl *WeatherStationImpl) Intensify(negative bool) bool {
	return weatherStationImpl.RunOnServer("intensify", map[string]interface{}{
		"negative": negative,
	}).(bool)
}

// Rotate runs logic that bribe the weathermen to change the direction of
// the next Forecast by rotating it clockwise or counterclockwise.
func (weatherStationImpl *WeatherStationImpl) Rotate(counterclockwise bool) bool {
	return weatherStationImpl.RunOnServer("rotate", map[string]interface{}{
		"counterclockwise": counterclockwise,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in WeatherStation.
func (weatherStationImpl *WeatherStationImpl) InitImplDefaults() {
	weatherStationImpl.BuildingImpl.InitImplDefaults()

}

// DeltaMerge merges the delta for a given attribute in WeatherStation.
func (weatherStationImpl *WeatherStationImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := weatherStationImpl.BuildingImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	_, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'anarchy.impl.DeltaMerge'",
		)
	}

	return false, nil // no errors in delta merging
}

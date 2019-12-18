package impl

import (
	"joueur/base"
	"joueur/games/anarchy"
)

// DeltaMerge is the set of functions that can delta merge a
// Anarchy game.
type DeltaMerge interface {
	base.DeltaMerge

	Building(interface{}) anarchy.Building
	FireDepartment(interface{}) anarchy.FireDepartment
	Forecast(interface{}) anarchy.Forecast
	GameObject(interface{}) anarchy.GameObject
	Player(interface{}) anarchy.Player
	PoliceDepartment(interface{}) anarchy.PoliceDepartment
	Warehouse(interface{}) anarchy.Warehouse
	WeatherStation(interface{}) anarchy.WeatherStation

	ArrayOfBuilding(*[]anarchy.Building, interface{}) []anarchy.Building
	ArrayOfFireDepartment(*[]anarchy.FireDepartment, interface{}) []anarchy.FireDepartment
	ArrayOfForecast(*[]anarchy.Forecast, interface{}) []anarchy.Forecast
	ArrayOfPlayer(*[]anarchy.Player, interface{}) []anarchy.Player
	ArrayOfPoliceDepartment(*[]anarchy.PoliceDepartment, interface{}) []anarchy.PoliceDepartment
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfWarehouse(*[]anarchy.Warehouse, interface{}) []anarchy.Warehouse
	ArrayOfWeatherStation(*[]anarchy.WeatherStation, interface{}) []anarchy.WeatherStation
	MapOfStringToGameObject(*map[string]anarchy.GameObject, interface{}) map[string]anarchy.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Anarchy Game Object Deltas -- \\

// Building attempts to return the instance of Building
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Building(delta interface{}) anarchy.Building {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBuilding, isBuilding := baseGameObject.(anarchy.Building)
	if !isBuilding {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.Building", delta)
	}

	return asBuilding
}

// FireDepartment attempts to return the instance of FireDepartment
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) FireDepartment(delta interface{}) anarchy.FireDepartment {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asFireDepartment, isFireDepartment := baseGameObject.(anarchy.FireDepartment)
	if !isFireDepartment {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.FireDepartment", delta)
	}

	return asFireDepartment
}

// Forecast attempts to return the instance of Forecast
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Forecast(delta interface{}) anarchy.Forecast {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asForecast, isForecast := baseGameObject.(anarchy.Forecast)
	if !isForecast {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.Forecast", delta)
	}

	return asForecast
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) anarchy.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(anarchy.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.GameObject", delta)
	}

	return asGameObject
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) anarchy.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(anarchy.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.Player", delta)
	}

	return asPlayer
}

// PoliceDepartment attempts to return the instance of PoliceDepartment
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) PoliceDepartment(delta interface{}) anarchy.PoliceDepartment {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPoliceDepartment, isPoliceDepartment := baseGameObject.(anarchy.PoliceDepartment)
	if !isPoliceDepartment {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.PoliceDepartment", delta)
	}

	return asPoliceDepartment
}

// Warehouse attempts to return the instance of Warehouse
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Warehouse(delta interface{}) anarchy.Warehouse {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asWarehouse, isWarehouse := baseGameObject.(anarchy.Warehouse)
	if !isWarehouse {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.Warehouse", delta)
	}

	return asWarehouse
}

// WeatherStation attempts to return the instance of WeatherStation
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) WeatherStation(delta interface{}) anarchy.WeatherStation {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asWeatherStation, isWeatherStation := baseGameObject.(anarchy.WeatherStation)
	if !isWeatherStation {
		(*deltaMergeImpl).CannotConvertDeltaTo("anarchy.WeatherStation", delta)
	}

	return asWeatherStation
}

// -- Deep Deltas -- \\

// ArrayOfBuilding delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfBuilding(state *[]anarchy.Building, delta interface{}) []anarchy.Building {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.Building, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Building(deltaValue)
	}
	return newArray
}

// ArrayOfFireDepartment delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfFireDepartment(state *[]anarchy.FireDepartment, delta interface{}) []anarchy.FireDepartment {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.FireDepartment, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.FireDepartment(deltaValue)
	}
	return newArray
}

// ArrayOfForecast delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfForecast(state *[]anarchy.Forecast, delta interface{}) []anarchy.Forecast {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.Forecast, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Forecast(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]anarchy.Player, delta interface{}) []anarchy.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfPoliceDepartment delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPoliceDepartment(state *[]anarchy.PoliceDepartment, delta interface{}) []anarchy.PoliceDepartment {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.PoliceDepartment, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.PoliceDepartment(deltaValue)
	}
	return newArray
}

// ArrayOfString delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfString(state *[]string, delta interface{}) []string {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]string, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.String(deltaValue)
	}
	return newArray
}

// ArrayOfWarehouse delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfWarehouse(state *[]anarchy.Warehouse, delta interface{}) []anarchy.Warehouse {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.Warehouse, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Warehouse(deltaValue)
	}
	return newArray
}

// ArrayOfWeatherStation delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfWeatherStation(state *[]anarchy.WeatherStation, delta interface{}) []anarchy.WeatherStation {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]anarchy.WeatherStation, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.WeatherStation(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]anarchy.GameObject, delta interface{}) map[string]anarchy.GameObject {
	deltaMap := (*deltaMergeImpl).ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if (*deltaMergeImpl).IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = deltaMergeImpl.GameObject(deltaValue)
		}
	}
	return *state
}

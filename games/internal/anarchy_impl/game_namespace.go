// Package impl contains all the Anarchy game implimentation logic.
package impl
// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// AnarchyNamespace is the collection of implimentation logic for the Anarchy game.
type AnarchyNamespace struct{}

// Name returns the name of the Anarchy game.
func (*AnarchyNamespace) Name() string {
	return "Anarchy"
}

// Version returns the current version hash as last generated for the Anarchy game.
func (*AnarchyNamespace) Version() string {
	return "2bc66f9a5d7babd553079e149c7466feb6f553935b608ff722872e195fbadab8"
}

// PlayerName returns the desired name of the AI in the Anarchy game.
func (*AnarchyNamespace) PlayerName() string {
	return anarchy.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Anarchy game.
func (*AnarchyNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Building":
		newBuilding := BuildingImpl{}
		newBuilding.InitImplDefaults()
		return &newBuilding, nil
	case "FireDepartment":
		newFireDepartment := FireDepartmentImpl{}
		newFireDepartment.InitImplDefaults()
		return &newFireDepartment, nil
	case "Forecast":
		newForecast := ForecastImpl{}
		newForecast.InitImplDefaults()
		return &newForecast, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "PoliceDepartment":
		newPoliceDepartment := PoliceDepartmentImpl{}
		newPoliceDepartment.InitImplDefaults()
		return &newPoliceDepartment, nil
	case "Warehouse":
		newWarehouse := WarehouseImpl{}
		newWarehouse.InitImplDefaults()
		return &newWarehouse, nil
	case "WeatherStation":
		newWeatherStation := WeatherStationImpl{}
		newWeatherStation.InitImplDefaults()
		return &newWeatherStation, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Anarchy' can be created")
}

// CreateGame is the factory method for Game the Anarchy game.
func (*AnarchyNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Anarchy game.
func (*AnarchyNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := anarchy.AI{}
	return &ai, &ai.AIImpl
}

func (*AnarchyNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Anarchy game.
func (*AnarchyNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*anarchy.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid anarchy.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

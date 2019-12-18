// Package impl contains all the Newtonian game implimentation logic.
package impl
// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// NewtonianNamespace is the collection of implimentation logic for the Newtonian game.
type NewtonianNamespace struct{}

// Name returns the name of the Newtonian game.
func (*NewtonianNamespace) Name() string {
	return "Newtonian"
}

// Version returns the current version hash as last generated for the Newtonian game.
func (*NewtonianNamespace) Version() string {
	return "7c19f909ee5faa0ac3faf4e989032b5a37ba94aeb5d6ae7654a15a2bb1401bbe"
}

// PlayerName returns the desired name of the AI in the Newtonian game.
func (*NewtonianNamespace) PlayerName() string {
	return newtonian.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Newtonian game.
func (*NewtonianNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Job":
		newJob := JobImpl{}
		newJob.InitImplDefaults()
		return &newJob, nil
	case "Machine":
		newMachine := MachineImpl{}
		newMachine.InitImplDefaults()
		return &newMachine, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	case "Unit":
		newUnit := UnitImpl{}
		newUnit.InitImplDefaults()
		return &newUnit, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Newtonian' can be created")
}

// CreateGame is the factory method for Game the Newtonian game.
func (*NewtonianNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Newtonian game.
func (*NewtonianNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := newtonian.AI{}
	return &ai, &ai.AIImpl
}

func (*NewtonianNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Newtonian game.
func (*NewtonianNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*newtonian.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid newtonian.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

// Package impl contains all the Stumped game implimentation logic.
package impl
// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// StumpedNamespace is the collection of implimentation logic for the Stumped game.
type StumpedNamespace struct{}

// Name returns the name of the Stumped game.
func (*StumpedNamespace) Name() string {
	return "Stumped"
}

// Version returns the current version hash as last generated for the Stumped game.
func (*StumpedNamespace) Version() string {
	return "7de307cae4a9a163a9b3600cb20c4b376b9f9cc42f1b990852878fea0127eed3"
}

// PlayerName returns the desired name of the AI in the Stumped game.
func (*StumpedNamespace) PlayerName() string {
	return stumped.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Stumped game.
func (*StumpedNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Beaver":
		newBeaver := BeaverImpl{}
		newBeaver.InitImplDefaults()
		return &newBeaver, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Job":
		newJob := JobImpl{}
		newJob.InitImplDefaults()
		return &newJob, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Spawner":
		newSpawner := SpawnerImpl{}
		newSpawner.InitImplDefaults()
		return &newSpawner, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Stumped' can be created")
}

// CreateGame is the factory method for Game the Stumped game.
func (*StumpedNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Stumped game.
func (*StumpedNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := stumped.AI{}
	return &ai, &ai.AIImpl
}

func (*StumpedNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Stumped game.
func (*StumpedNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*stumped.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid stumped.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

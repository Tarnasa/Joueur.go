// Package impl contains all the Stardash game implimentation logic.
package impl
// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// StardashNamespace is the collection of implimentation logic for the Stardash game.
type StardashNamespace struct{}

// Name returns the name of the Stardash game.
func (*StardashNamespace) Name() string {
	return "Stardash"
}

// Version returns the current version hash as last generated for the Stardash game.
func (*StardashNamespace) Version() string {
	return "0fa378e83ac567ebdf3e9805d3f130023f936e2740acda173d238b37f2b5d541"
}

// PlayerName returns the desired name of the AI in the Stardash game.
func (*StardashNamespace) PlayerName() string {
	return stardash.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Stardash game.
func (*StardashNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Body":
		newBody := BodyImpl{}
		newBody.InitImplDefaults()
		return &newBody, nil
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
	case "Projectile":
		newProjectile := ProjectileImpl{}
		newProjectile.InitImplDefaults()
		return &newProjectile, nil
	case "Unit":
		newUnit := UnitImpl{}
		newUnit.InitImplDefaults()
		return &newUnit, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Stardash' can be created")
}

// CreateGame is the factory method for Game the Stardash game.
func (*StardashNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Stardash game.
func (*StardashNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := stardash.AI{}
	return &ai, &ai.AIImpl
}

func (*StardashNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Stardash game.
func (*StardashNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*stardash.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid stardash.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

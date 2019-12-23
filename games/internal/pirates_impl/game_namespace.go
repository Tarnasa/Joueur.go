// Package impl contains all the Pirates game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// PiratesNamespace is the collection of implimentation logic for the Pirates game.
type PiratesNamespace struct{}

// Name returns the name of the Pirates game.
func (*PiratesNamespace) Name() string {
	return "Pirates"
}

// Version returns the current version hash as last generated for the Pirates game.
func (*PiratesNamespace) Version() string {
	return "d51fca49d06cb7164f9dbf9c3515ab0f9b5a17113a5946bddcc75aaba125967f"
}

// PlayerName returns the desired name of the AI in the Pirates game.
func (*PiratesNamespace) PlayerName() string {
	return pirates.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Pirates game.
func (*PiratesNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Port":
		newPort := PortImpl{}
		newPort.InitImplDefaults()
		return &newPort, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	case "Unit":
		newUnit := UnitImpl{}
		newUnit.InitImplDefaults()
		return &newUnit, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Pirates' can be created")
}

// CreateGame is the factory method for Game the Pirates game.
func (*PiratesNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Pirates game.
func (*PiratesNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := pirates.AI{}
	return &ai, &ai.AIImpl
}

func (*PiratesNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Pirates game.
func (*PiratesNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*pirates.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid pirates.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

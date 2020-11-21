// Package impl contains all the Galapagos game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// GalapagosNamespace is the collection of implimentation logic for the Galapagos game.
type GalapagosNamespace struct{}

// Name returns the name of the Galapagos game.
func (*GalapagosNamespace) Name() string {
	return "Galapagos"
}

// Version returns the current version hash as last generated for the Galapagos game.
func (*GalapagosNamespace) Version() string {
	return "9337bc3d0f54640d13df30becf2e1bbda2e85a7be44fea6b072e638afc7e8dd7"
}

// PlayerName returns the desired name of the AI in the Galapagos game.
func (*GalapagosNamespace) PlayerName() string {
	return galapagos.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Galapagos game.
func (*GalapagosNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Creature":
		newCreature := CreatureImpl{}
		newCreature.InitImplDefaults()
		return &newCreature, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Plant":
		newPlant := PlantImpl{}
		newPlant.InitImplDefaults()
		return &newPlant, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Galapagos' can be created")
}

// CreateGame is the factory method for Game the Galapagos game.
func (*GalapagosNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Galapagos game.
func (*GalapagosNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := galapagos.AI{}
	return &ai, &ai.AIImpl
}

func (*GalapagosNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Galapagos game.
func (*GalapagosNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*galapagos.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid galapagos.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

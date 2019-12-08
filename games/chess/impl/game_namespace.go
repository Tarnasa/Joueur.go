// Package impl contains all the Chess implimentation
// logic and structures required by aa client to play with a game server.
// To start coding your AI open ./ai.go
package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/chess"
)

// ChessNamespace is the collection of implimentation logic for the Chess game.
type ChessNamespace struct {}

// Name returns the name of the Chess game.
func (*ChessNamespace) Name() string {
	return "Chess"
}

// Version returns the current version hash as last generated for the Chess game.
func (*ChessNamespace) Version() string {
	return "cfa5f5c1685087ce2899229c04c26e39f231e897ecc8fe036b44bc22103ef801"
}

// PlayerName returns the desired name of the AI in the Chess game.
func (*ChessNamespace) PlayerName() string {
	return chess.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Chess game.
func (*ChessNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch (gameObjectName) {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Chess' can be created")
}

// CreateGame is the factory method for Game the Chess game.
func (*ChessNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Chess game.
func (*ChessNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := chess.AI{}
	return &ai, &ai.AIImpl
}

func (*ChessNamespace) CreateDeltaMerge(deltaMergeImpl base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Chess game.
func (*ChessNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*chess.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid chess.AI to order!")
	}
	switch functionName {
	case "makeMove":
		return (*ai).MakeMove(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

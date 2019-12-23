// Package impl contains all the Checkers game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/checkers"
)

// CheckersNamespace is the collection of implimentation logic for the Checkers game.
type CheckersNamespace struct{}

// Name returns the name of the Checkers game.
func (*CheckersNamespace) Name() string {
	return "Checkers"
}

// Version returns the current version hash as last generated for the Checkers game.
func (*CheckersNamespace) Version() string {
	return "49f1e5586cc4c62b6f74081e803d8edf9f54e8315f221c62c638f963cea8ab31"
}

// PlayerName returns the desired name of the AI in the Checkers game.
func (*CheckersNamespace) PlayerName() string {
	return checkers.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Checkers game.
func (*CheckersNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Checker":
		newChecker := CheckerImpl{}
		newChecker.InitImplDefaults()
		return &newChecker, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Checkers' can be created")
}

// CreateGame is the factory method for Game the Checkers game.
func (*CheckersNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Checkers game.
func (*CheckersNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := checkers.AI{}
	return &ai, &ai.AIImpl
}

func (*CheckersNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Checkers game.
func (*CheckersNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*checkers.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid checkers.AI to order!")
	}
	switch functionName {
	case "gotCaptured":
		arg0 := args[0].(checkers.Checker)
		(*ai).GotCaptured(arg0)
		return nil, nil
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

// Package impl contains all the Coreminer game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// CoreminerNamespace is the collection of implimentation logic for the Coreminer game.
type CoreminerNamespace struct{}

// Name returns the name of the Coreminer game.
func (*CoreminerNamespace) Name() string {
	return "Coreminer"
}

// Version returns the current version hash as last generated for the Coreminer game.
func (*CoreminerNamespace) Version() string {
	return "b559778acd8e4c689b8d028ca6cc154714ce79c39b09cd6d171b50faf88ef747"
}

// PlayerName returns the desired name of the AI in the Coreminer game.
func (*CoreminerNamespace) PlayerName() string {
	return coreminer.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Coreminer game.
func (*CoreminerNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Bomb":
		newBomb := BombImpl{}
		newBomb.InitImplDefaults()
		return &newBomb, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Miner":
		newMiner := MinerImpl{}
		newMiner.InitImplDefaults()
		return &newMiner, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	case "Upgrade":
		newUpgrade := UpgradeImpl{}
		newUpgrade.InitImplDefaults()
		return &newUpgrade, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Coreminer' can be created")
}

// CreateGame is the factory method for Game the Coreminer game.
func (*CoreminerNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Coreminer game.
func (*CoreminerNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := coreminer.AI{}
	return &ai, &ai.AIImpl
}

func (*CoreminerNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Coreminer game.
func (*CoreminerNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*coreminer.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid coreminer.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

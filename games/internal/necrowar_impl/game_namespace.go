// Package impl contains all the Necrowar game implimentation logic.
package impl
// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// NecrowarNamespace is the collection of implimentation logic for the Necrowar game.
type NecrowarNamespace struct{}

// Name returns the name of the Necrowar game.
func (*NecrowarNamespace) Name() string {
	return "Necrowar"
}

// Version returns the current version hash as last generated for the Necrowar game.
func (*NecrowarNamespace) Version() string {
	return "935f0e64ba290cdce31688a40bd90d1eb5375f36aeebd67482238fc0da25ef86"
}

// PlayerName returns the desired name of the AI in the Necrowar game.
func (*NecrowarNamespace) PlayerName() string {
	return necrowar.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Necrowar game.
func (*NecrowarNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Tile":
		newTile := TileImpl{}
		newTile.InitImplDefaults()
		return &newTile, nil
	case "Tower":
		newTower := TowerImpl{}
		newTower.InitImplDefaults()
		return &newTower, nil
	case "TowerJob":
		newTowerJob := TowerJobImpl{}
		newTowerJob.InitImplDefaults()
		return &newTowerJob, nil
	case "Unit":
		newUnit := UnitImpl{}
		newUnit.InitImplDefaults()
		return &newUnit, nil
	case "UnitJob":
		newUnitJob := UnitJobImpl{}
		newUnitJob.InitImplDefaults()
		return &newUnitJob, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Necrowar' can be created")
}

// CreateGame is the factory method for Game the Necrowar game.
func (*NecrowarNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Necrowar game.
func (*NecrowarNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := necrowar.AI{}
	return &ai, &ai.AIImpl
}

func (*NecrowarNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Necrowar game.
func (*NecrowarNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*necrowar.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid necrowar.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

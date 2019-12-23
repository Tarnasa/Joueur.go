// Package impl contains all the Saloon game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// SaloonNamespace is the collection of implimentation logic for the Saloon game.
type SaloonNamespace struct{}

// Name returns the name of the Saloon game.
func (*SaloonNamespace) Name() string {
	return "Saloon"
}

// Version returns the current version hash as last generated for the Saloon game.
func (*SaloonNamespace) Version() string {
	return "fbaeac2bae9020bdd5a8816cb9ae38215c277e4bf7874a2f70c3995cd8eee8d3"
}

// PlayerName returns the desired name of the AI in the Saloon game.
func (*SaloonNamespace) PlayerName() string {
	return saloon.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Saloon game.
func (*SaloonNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "Bottle":
		newBottle := BottleImpl{}
		newBottle.InitImplDefaults()
		return &newBottle, nil
	case "Cowboy":
		newCowboy := CowboyImpl{}
		newCowboy.InitImplDefaults()
		return &newCowboy, nil
	case "Furnishing":
		newFurnishing := FurnishingImpl{}
		newFurnishing.InitImplDefaults()
		return &newFurnishing, nil
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
	case "YoungGun":
		newYoungGun := YoungGunImpl{}
		newYoungGun.InitImplDefaults()
		return &newYoungGun, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Saloon' can be created")
}

// CreateGame is the factory method for Game the Saloon game.
func (*SaloonNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Saloon game.
func (*SaloonNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := saloon.AI{}
	return &ai, &ai.AIImpl
}

func (*SaloonNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Saloon game.
func (*SaloonNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*saloon.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid saloon.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

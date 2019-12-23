// Package impl contains all the Spiders game implimentation logic.
package impl

// To start coding your AI please open ./ai.go

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// SpidersNamespace is the collection of implimentation logic for the Spiders game.
type SpidersNamespace struct{}

// Name returns the name of the Spiders game.
func (*SpidersNamespace) Name() string {
	return "Spiders"
}

// Version returns the current version hash as last generated for the Spiders game.
func (*SpidersNamespace) Version() string {
	return "a8df6038306b6855bb35959d7698f8dcbf98f48e7e148de59fef940ccb241bdf"
}

// PlayerName returns the desired name of the AI in the Spiders game.
func (*SpidersNamespace) PlayerName() string {
	return spiders.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Spiders game.
func (*SpidersNamespace) CreateGameObject(gameObjectName string) (base.DeltaMergeableGameObject, error) {
	switch gameObjectName {
	case "BroodMother":
		newBroodMother := BroodMotherImpl{}
		newBroodMother.InitImplDefaults()
		return &newBroodMother, nil
	case "Cutter":
		newCutter := CutterImpl{}
		newCutter.InitImplDefaults()
		return &newCutter, nil
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, nil
	case "Nest":
		newNest := NestImpl{}
		newNest.InitImplDefaults()
		return &newNest, nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, nil
	case "Spider":
		newSpider := SpiderImpl{}
		newSpider.InitImplDefaults()
		return &newSpider, nil
	case "Spiderling":
		newSpiderling := SpiderlingImpl{}
		newSpiderling.InitImplDefaults()
		return &newSpiderling, nil
	case "Spitter":
		newSpitter := SpitterImpl{}
		newSpitter.InitImplDefaults()
		return &newSpitter, nil
	case "Weaver":
		newWeaver := WeaverImpl{}
		newWeaver.InitImplDefaults()
		return &newWeaver, nil
	case "Web":
		newWeb := WebImpl{}
		newWeb.InitImplDefaults()
		return &newWeb, nil
	}
	return nil, errors.New("no game object named '" + gameObjectName + "' for game 'Spiders' can be created")
}

// CreateGame is the factory method for Game the Spiders game.
func (*SpidersNamespace) CreateGame() base.DeltaMergeableGame {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game
}

// CreateAI is the factory method for the AI in the Spiders game.
func (*SpidersNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := spiders.AI{}
	return &ai, &ai.AIImpl
}

func (*SpidersNamespace) CreateDeltaMerge(deltaMergeImpl *base.DeltaMergeImpl) base.DeltaMerge {
	return &DeltaMergeImpl{
		DeltaMergeImpl: *deltaMergeImpl,
	}
}

// OrderAI handles an order for the AI in the Spiders game.
func (*SpidersNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*spiders.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid spiders.AI to order!")
	}
	switch functionName {
	case "runTurn":
		return (*ai).RunTurn(), nil
	}

	return nil, errors.New("Cannot find functionName " + functionName + " to run in S{game_name} AI")
}

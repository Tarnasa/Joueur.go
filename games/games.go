package games

import (
	"errors"
	"joueur/base"
)

type GameNamespace interface {
	Name() string
	Version() string
	PlayerName() string
	CreateAI() (base.BaseAI, *base.BaseAIImpl)
	CreateGame() (base.BaseGame, *base.BaseDeltaMergeableImpl)
	CreateGameObject(string) (base.BaseGameObject, *base.BaseDeltaMergeableImpl, error)
	OrderAI(base.BaseAI, string, []interface{}) (interface{}, error)
}

var gamesNamespaceTypes = make(map[string]GameNamespace)

func Register(gameName string, gameNamespace GameNamespace) {
	gamesNamespaceTypes[gameName] = gameNamespace
}

func Get(gameName string) (GameNamespace, error) {
	if gameNamespace, found := gamesNamespaceTypes[gameName]; found {
		return gameNamespace, nil
	}

	return nil, errors.New("Cannot get and create namespace for game " + gameName)
}

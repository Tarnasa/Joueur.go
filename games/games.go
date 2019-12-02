package games

import (
	"errors"
	"joueur/base"
	"reflect"
)

type GameNamespace interface {
	Name() string
	Version() string
	PlayerName() string
	CreateAI() *base.BaseAI
	CreateGame() *base.BaseGame
	CreateGameObject() (*base.BaseGameObject, error)
}

var gamesNamespaceTypes = make(map[string](reflect.Type))

func Register(gameName string, namespaceType reflect.type) {
	gamesNamespaceTypes[gameName] = namespace
}

func Get(gameName string) (*GameNamespace, error) {
	if namespaceType, ok := gamesNamespaceTypes[gameName]; ok {
		reflectedNamespace = reflect.New(namespaceType.Elem())
		if reflectedNamespace.IsValid() {
			namespace := reflectedNamespace.Interface(GameNamespace)

			return &namespace, nil
		}

	}

	return nil, errors.New("Cannot get and create namespace for game " + gameName)
}

package games

import (
	"errors"
	"reflect"
)

type GameNamespaceTypes struct {
	Game reflect.Type
	// AI reflect.Type
}

type GameNamespace struct {
	Version string
	Types   GameNamespaceTypes
}

var gamesRegistry = make(map[string]*GameNamespace)

func Register(gameName string, namespace *GameNamespace) {
	gamesRegistry[gameName] = namespace
}

func Get(gameName string) (*GameNamespace, error) {
	if namespace, ok := gamesRegistry[gameName]; ok {
		return namespace, nil
	} else {
		return nil, errors.New("Cannot get game " + gameName)
	}
}

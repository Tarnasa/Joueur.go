package games

import (
	"errors"
	"reflect"
)

type GameNamespace struct {
	Version         string
	GameType        reflect.Type
	AIType          reflect.Type
	GameObjectTypes map[string]reflect.Type
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

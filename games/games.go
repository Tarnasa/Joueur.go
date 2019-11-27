package games

import (
	"errors"
	"reflect"
)

type GameStructs = map[string]reflect.Type

var gamesRegistry = make(map[string]GameStructs)

func Register(gameName string, structs GameStructs) {
	gamesRegistry[gameName] = structs
}

func Get(gameName string) (GameStructs, error) {
	if structs, ok := gamesRegistry[gameName]; ok {
		return structs, nil
	} else {
		return nil, errors.New("Cannot get game " + gameName)
	}
}

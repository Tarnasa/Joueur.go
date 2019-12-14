// Package games collects and registers all the available games this
// Joueur client can play.
package games

import (
	"errors"
	"joueur/base"
)

var gamesNamespaceTypes = make(map[string]base.GameNamespace)

func Register(gameName string, gameNamespace base.GameNamespace) {
	gamesNamespaceTypes[gameName] = gameNamespace
}

func Get(gameName string) (base.GameNamespace, error) {
	if gameNamespace, found := gamesNamespaceTypes[gameName]; found {
		return gameNamespace, nil
	}

	return nil, errors.New(
		"Cannot get and create namespace for game " + gameName,
	)
}

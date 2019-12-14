// Package games collects and registers all the available games this
// Joueur client can play.
package games

import (
	"errors"
	"joueur/base"
	"joueur/games/internal"
)

// Get attempts to retrive the GameNamespace for a given game by its name.
// If none is found nil and an error are returned.
func Get(gameName string) (base.GameNamespace, error) {
	if gameNamespace, found := internal.GameNamespaceRegistry[gameName]; found {
		return gameNamespace, nil
	}

	return nil, errors.New(
		"Cannot get and create namespace for game " + gameName,
	)
}

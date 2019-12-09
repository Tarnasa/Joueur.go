package base

import (
	"errors"
)

// Game is the base interface all games should implement for their Game interfaces.
type Game interface {
	GetGameObject(string) (GameObject, bool)
}

// DeltaMergeableGame is a Game that is also DeltaMergeable
type DeltaMergeableGame interface {
	DeltaMergeable
	Game

	AddGameObject(string, GameObject) error
}

// GameImpl is the implimentation struct for the Game interface.
type GameImpl struct {
	DeltaMergeableImpl

	gameObjectsImpl map[string]GameObject
}

// GetGameObject simply attempts to get a game object from inside its gameObjects map.
func (gameImpl *GameImpl) GetGameObject(id string) (GameObject, bool) {
	gameObject, found := gameImpl.gameObjectsImpl[id]
	return gameObject, found
}

// AddGameObject adds a new gae object to the game.
// However if the GameObject is already present it returns an error.
func (gameImpl *GameImpl) AddGameObject(id string, gameObject GameObject) error {
	if _, alreadyHasGameObject := gameImpl.gameObjectsImpl[id]; alreadyHasGameObject {
		return errors.New("GameObject #" + id + " cannot be added. Already present in Game")
	}

	gameImpl.gameObjectsImpl[id] = gameObject
	return nil
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.gameObjectsImpl = make(map[string]GameObject)
}

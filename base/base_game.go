package base

import (
	"errors"
	"joueur/runtime/errorhandler"
)

// Game is the base interface all games should implement for their Game interfaces.
type Game interface {
	GetGameObject(string) (GameObject, bool)
}

// GameImpl is the implimentation struct for the Game interface.
type GameImpl struct {
	DeltaMergeableImpl
}

// GetGameObject simply attempts to get a game object from inside its gameObjects map.
func (gameImpl GameImpl) GetGameObject(id string) (GameObject, bool) {
	gameObjectsRaw, found := (gameImpl.InternalDataMap)["gameObjects"]
	gameObjectsMap, isMap := gameObjectsRaw.(map[string]GameObject)
	if !found || gameObjectsMap == nil || !isMap {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("gameObjects not found as a map of strings to GameObjects in Game"),
		)
	}

	gameObject, found := gameObjectsMap[id]
	return gameObject, found
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
}

package base

import (
	"errors"
	"joueur/runtime/errorhandler"
)

// RunOnServerCallback is the callback function for the game manager to hook
// into so RunOnServer works once the client is connected
var RunOnServerCallback func(GameObject, string, map[string]interface{}) interface{}

// GameObject is the base interface all GameObjects in all games should implement
type GameObject interface {
	GameObjectName() string
	ID() string
}

// DeltaMergeableGameObject is a GameObject that is also DeltaMergeable
type DeltaMergeableGameObject interface {
	DeltaMergeable
	GameObject
}

// GameObjectImpl is the implimentation struct for BaseGameObject
type GameObjectImpl struct {
	DeltaMergeableImpl

	game               Game
	gameObjectNameImpl string
	idImpl             string
}

// RunOnServer is a slim wrapper that attempts to run game logic on behalf
// of this gameObject on the server.
func (gameObjectImpl *GameObjectImpl) RunOnServer(functionName string, args map[string]interface{}) interface{} {
	if RunOnServerCallback == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Cannot invoke function "+functionName+"on server without callback set!"),
		)
	}

	return RunOnServerCallback(gameObjectImpl, functionName, args)
}

// GameObjectName returns string representing the top level Class that this
// game object is an instance of. Used for reflection to create new
// instances on clients, but exposed for convenience should AIs want this
// data.
func (gameObjectImpl *GameObjectImpl) GameObjectName() string {
	return gameObjectImpl.gameObjectNameImpl
}

// ID returns a unique id for each instance of a GameObject or a sub class.
// Used for client and server communication. Should never change value after being set.
func (gameObjectImpl *GameObjectImpl) ID() string {
	return gameObjectImpl.idImpl
}

// InitImplDefaults initializes safe defaults for all fields in GameObject.
func (gameObjectImpl *GameObjectImpl) InitImplDefaults() {
	gameObjectImpl.idImpl = ""
}

// DeltaMerge merges the delta for a given attribute in GameObject.
func (gameObjectImpl *GameObjectImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) (bool, error) {
	switch attribute {
	case "gameObjectName":
		gameObjectImpl.gameObjectNameImpl = deltaMerge.String(delta)
		return true, nil
	case "id":
		gameObjectImpl.idImpl = deltaMerge.String(delta)
		return true, nil
	}

	return false, nil
}

// String returns the string friendly representation of this game object
func (gameObjectImpl *GameObjectImpl) String() string {
	// format is 'GameObjectName #1234'
	return gameObjectImpl.GameObjectName() + " #" + gameObjectImpl.ID()
}

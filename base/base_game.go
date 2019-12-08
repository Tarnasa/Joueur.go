package base

// Game is the base interface all games should implement for their Game interfaces.
type Game interface {
	GetGameObject(string) (GameObject, bool)
}

// GameImpl is the implimentation struct for the Game interface.
type GameImpl struct {
	DeltaMergeableImpl

	GameObjectsImpl map[string]GameObject
}

// GetGameObject simply attempts to get a game object from inside its gameObjects map.
func (gameImpl *GameImpl) GetGameObject(id string) (GameObject, bool) {
	gameObject, found := gameImpl.GameObjectsImpl[id]
	return gameObject, found
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameObjectsImpl = make(map[string]GameObject)
}

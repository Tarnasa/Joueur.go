package base

type BaseGameImpl struct {
	// All the game objects in the game, indexed by their game object ID.
	GameObjects map[string]*BaseGameObject
}

type BaseGame interface{}

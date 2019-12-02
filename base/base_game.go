package base

type BaseGameImpl struct {
	// All the game objects in the game, indexed by their game object ID.
	GameObjects map[string]*BaseGameObject
}

func (_ BaseGameImpl) RunOnServer(functionName string, args map[string]interface{}) interface{} {
}

type BaseGame interface{}

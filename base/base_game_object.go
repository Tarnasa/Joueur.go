package base

type BaseGameObject interface {
	/*
		// A unique ID (unique to the game instance) of the game object.
		// Will never change, and IDs are never re-used.
		Id() string

		GameObjectName() string
	*/
}

type BaseGameObjectImpl struct {
	game *BaseGame
	data map[string]interface{}
}

/*
func (this BaseGameObjectImpl) Id() string {
	return this.data["id"].(string)
}

func (this BaseGameObjectImpl) GameObjectName() string {
	return this.data["gameObjectName"].(string)
}
*/

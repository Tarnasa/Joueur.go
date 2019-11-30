package base

type BaseGameObject struct {
	// A unique ID (unique to the game instance) of the game object.
	// Will never change, and IDs are never re-used.
	Id string

	GameObjectName string
}

func (gameObject BaseGameObject) runOnServer(name string) {
	// TODO: do
}

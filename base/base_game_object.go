package base

type BaseGameObject struct {
	// A unique ID (unique to the game instance) of the game object.
	// Will never change, and IDs are never re-used.
	Id string
}

func (gameObject BaseGameObject) RunOnServer(name string) {
	// TODO: do
}

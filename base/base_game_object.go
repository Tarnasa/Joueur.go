package base

type BaseGameObject interface {
	// A unique ID (unique to the game instance) of the game object.
	// Will never change, and IDs are never re-used.
	Id() string

	GameObjectName() string
}

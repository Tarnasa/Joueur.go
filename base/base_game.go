package base

type BaseGame struct {
	// All the game objects in the game, indexed by their game object ID.
	GameObjects map[string]*BaseGameObject

	// The players in this game.
	Players []*BasePlayer

	// The session this game is occuring in.
	Session string
}

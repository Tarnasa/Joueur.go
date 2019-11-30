package chess

import "joueur/base"

// An object in the game. The most basic class that all game classes should inherit from automatically.
type GameObject struct {
	base.BaseGameObject

	// The reference to the Game instance this GameObject is in.
	game *Game

	// String representing the top level Class that this game object is an instance of. Used for reflection to create new instances on clients, but exposed for convenience should AIs want this data.
	GameObjectName string

	// A unique id for each instance of a GameObject or a sub class. Used for client and server communication. Should never change value after being set.
	Id string

	// Any strings logged will be stored here. Intended for debugging.
	Logs []string
}

// Chess specific GameObject actions

// Adds a message to this GameObject's logs. Intended for your own debugging purposes, as strings stored here are saved in the gamelog.
func (gameObject GameObject) Log(
	// A string to add to this GameObject's log. Intended for debugging.
	message string,
) {
	// TODO: somehow we need to client.RunOnServer(args)
	return
}

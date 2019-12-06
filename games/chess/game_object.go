package chess

import "joueur/base"

// An object in the game. The most basic class that all game classes should inherit from automatically.
type GameObject interface {
	// Parent interfaces
	base.GameObject

	// -- Attributes -- \\

	// Any strings logged will be stored here. Intended for debugging.
	Logs() []string

	// -- Methods -- \\

	// Adds a message to this GameObject's logs. Intended for your own debugging purposes, as strings stored here are saved in the gamelog.
	Log(string)

}

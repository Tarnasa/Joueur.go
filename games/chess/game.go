package chess

import "joueur/base"

// The traditional 8x8 chess board with pieces.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Forsyth-Edwards Notation (fen), a notation that describes the game board state.
	Fen() string

	// A mapping of every game object's ID to the actual game object. Primarily used by the server and client to easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// The list of [known] moves that have occurred in the game, in Standard Algebraic Notation (SAN) format. The first element is the first move, with the last being the most recent.
	History() []string

	// List of all the players in the game.
	Players() []Player

	// A unique identifier for the game instance that is being played.
	Session() string

}

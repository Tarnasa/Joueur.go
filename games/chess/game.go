package chess

import "joueur/base"

// Game is the traditional 8x8 chess board with pieces.
type Game interface {
	// Parent interfaces
	base.Game

	// -- Attributes -- \\

	// Fen is forsyth-Edwards Notation (fen), a notation that
	// describes the game board state.
	Fen() string

	// GameObjects is a mapping of every game object's ID to the
	// actual game object. Primarily used by the server and client to
	// easily refer to the game objects via ID.
	GameObjects() map[string]GameObject

	// History is the array of [known] moves that have occurred in the
	// game, in Standard Algebraic Notation (SAN) format. The first
	// element is the first move, with the last being the most recent.
	History() []string

	// Players is array of all the players in the game.
	Players() []Player

	// Session is a unique identifier for the game instance that is
	// being played.
	Session() string
}

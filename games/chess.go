package games

// This file registers the game.
// Removing/modifying it means your AI may not work correctly as the game won't exist!

import "joueur/games/chess/impl"

func init() {
	Register("Chess", &impl.ChessNamespace{})
}

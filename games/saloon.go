package games

// This file registers the game.
// Removing/modifying it means your AI may not work correctly as the game won't exist!

import "joueur/games/saloon/impl"

func init() {
	Register("Saloon", &impl.SaloonNamespace{})
}

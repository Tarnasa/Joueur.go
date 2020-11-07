package internal

// This file registers the game.
// Removing/modifying it means your AI may not work correctly as the game won't exist!

import "joueur/games/internal/coreminer_impl"

func init() {
	register("Coreminer", &impl.CoreminerNamespace{})
}

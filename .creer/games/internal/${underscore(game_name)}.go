package internal

// This file registers the game.
// Removing/modifying it means your AI may not work correctly as the game won't exist!

import "joueur/games/internal/${underscore(game_name)}_impl"

func init() {
	register("${game_name}", &impl.${game_name}Namespace{})
}

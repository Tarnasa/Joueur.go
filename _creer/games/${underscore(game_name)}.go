package games

// This file registers the game.
// Removing/modifying it means your AI may not work correctly as the game won't exist!

import "joueur/games/${underscore(game_name)}/impl"

func init() {
	Register("${game_name}", &impl.${game_name}Namespace{})
}

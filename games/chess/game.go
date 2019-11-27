// DO NOT MODIFY THIS FILE
// Never try to directly create an instance of this class, or modify it.
// Instead, you should only be reading its variables and calling its functions.

package chess

import "joueur/base"

type Game struct {
	base.BaseGame

	Fen string
	// GameObjects map[string]GameObject
	History []string
}

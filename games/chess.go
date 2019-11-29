package games

import (
	"joueur/games/chess"
	"reflect"
)

func init() {
	Register("Chess", &GameNamespace{
		Version: "cfa5f5c1685087ce2899229c04c26e39f231e897ecc8fe036b44bc22103ef801",
		Types: GameNamesapceTypes{
			AI: reflect.TypeOf((*chess.AI)(nil)).Elem(),
			GameObject: reflect.TypeOf((*chess.GameObject)(nil)).Elem(),
			Player: reflect.TypeOf((*chess.Player)(nil)).Elem(),
		},
	})
}

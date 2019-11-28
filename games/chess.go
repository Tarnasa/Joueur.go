package games

import (
	"joueur/games/chess"
	"reflect"
)

func init() {
	Register("Chess", &GameNamespace{
		Version: "version hash will go here",
		Types: GameNamesapceTypes{
			Game: reflect.TypeOf((*chess.Game)(nil)).Elem(),
		},
	})
}

package games

import (
	"joueur/games/chess"
	"reflect"
)

func init() {
	Register("Chess", map[string]reflect.Type{
		"Game": reflect.TypeOf((*chess.Game)(nil)).Elem(),
	})
}

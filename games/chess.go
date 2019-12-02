package games

import (
	"joueur/games/chess/impl"
	"reflect"
)

func init() {
	Register("Chess", reflect.TypeOf((*(impl.ChessNamespace))(nil)).Elem(),)
}

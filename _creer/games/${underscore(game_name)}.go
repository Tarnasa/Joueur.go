package games

import (
	"joueur/games/${underscore(game_name)}/impl"
	"reflect"
)

func init() {
	Register("${game_name}", reflect.TypeOf((*(impl.${game_name}Namespace))(nil)).Elem(),)
}

package games

import (
	"joueur/games/${underscore(game_name)}"
	"reflect"
)

func init() {
	Register("${game_name}", &GameNamespace{
		Version: "${game_version}",
		Types: GameNamesapceTypes{
% for game_obj_name in (['AI'] + game_obj_names):
			${game_obj_name}: reflect.TypeOf((*chess.${game_obj_name})(nil)).Elem(),
% endfor
		},
	})
}

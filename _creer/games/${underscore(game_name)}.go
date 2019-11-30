package games

import (
	"joueur/games/${underscore(game_name)}"
	"reflect"
)

func init() {
	Register("${game_name}", &GameNamespace{
		Name: "${game_name}",
		Version: "${game_version}",
		GameType: reflect.TypeOf((*${underscore(game_name)}.Game)(nil)).Elem(),
		AIType: reflect.TypeOf((*${underscore(game_name)}.AI)(nil)).Elem(),
		GameObjectTypes: map[string]reflect.Type{
% for game_obj_name in (['AI'] + game_obj_names):
			"${game_obj_name}": reflect.TypeOf((*${underscore(game_name)}.${game_obj_name})(nil)).Elem(),
% endfor
		},
	})
}

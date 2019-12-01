// This package contains all the structs, methods, and the AI required as
// a client to play the ${game_name} with a game server.
// To start coding your AI open ./ai.go
package ${lowercase_first(game_name)}
<%include file='functions.noCreer' />
import (
	"errors"
	"joueur/base"
)
% for obj_name in game_obj_names:
<%
	obj = game_objs[obj_name]
	parents = list()
	if 'parentClasses' in obj:
		parents.extend(obj['parentClasses'])
	if obj_name in ['Game', 'GameObject']:
		parents.append('base.Base' + obj_name + "Impl")
%>
type ${obj_name}Impl struct {
%	for parent in parents:
	${parent}
%	endfor
%	if obj_name == 'GameObject':
	game *Game
%   endif
	data map[string]interface{}
}
%   if obj_name == 'GameObject':

func (this GameObjectImpl) Game() *Game {
	return this.game
}
%   endif
%   for attr_name in obj['attribute_names']:
<%
		attr = obj['attributes'][attr_name]
		ret_type = shared['go']['type'](attr['type'])
%>
func (this ${obj_name}Impl) ${upcase_first(attr_name)}() ${ret_type} {
	return this.data["${attr_name}"].(${ret_type})
}
%   endfor
%   for func_name in obj['function_names']:
<%
		func = obj['functions'][func_name]
		ret_type = shared['go']['type'](func['returns']['type']) if func['returns'] else ''
		argify = lambda a : '{} {}'.format(a['name'], shared['go']['type'](a['type']))
		args = ', '.join([argify(a) for a in func['arguments']])
%>
func (this ${obj_name}Impl) ${upcase_first(func_name)}(${args})${' ' if ret_type else ''} {
	${'return ' if ret_type else ''}this.RunOnServer(make(map[string]interface{
%		for arg in func['arguments']:
		"${arg['name']}": ${arg['name']},
%		endfor
	}))${('.('+ret_type+')') if ret_type else ''}
}
%   endfor
% endfor

// Factory functions

func CreateGameObject(gameObjectName string) (*GameObject, error) {
	switch (gameObjectName) {
% for game_obj_name in game_obj_names:
	case "${game_obj_name}":
		return &(${game_obj_name}Impl{}), nil
% endfor
	}
	return nil, errors.New("No game object named " + gameObjectName + " for game ${game_name}")
}

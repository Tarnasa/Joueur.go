// This package contains all the structs, methods, and the AI required as
// a client to play the ${game_name} with a game server.
// To start coding your AI open ./ai.go
package impl
<%include file='functions.noCreer' /><%
	package_name = underscore(game_name)
	ptype = lambda t : shared['go']['type'](t, package_name)
%>
import (
	"errors"
	"joueur/base"
	"joueur/games/${underscore(game_name)}"
)
% for obj_name in (['Game'] + game_obj_names):
<%
	obj = game if obj_name == 'Game' else game_objs[obj_name]
	parents = list()
	if 'parentClasses' in obj:
		parents.extend(obj['parentClasses'])
	if obj_name in ['Game', 'GameObject']:
		parents.append('base.Base' + obj_name)
%>
// -- ${obj_name} -- ${'\\\\'}

type ${obj_name}Impl struct {
%	for parent in parents:
	${parent}Impl
%	endfor
%	if obj_name == 'GameObject':
	game *${package_name}.Game
%   endif
}
%   if obj_name == 'GameObject':

func (this GameObjectImpl) Game() *${package_name}.Game {
	return this.game
}
%   endif
%   for attr_name in obj['attribute_names']:
<%
		attr = obj['attributes'][attr_name]
		ret_type = ptype(attr['type'])
%>
func (this ${obj_name}Impl) ${upcase_first(attr_name)}() ${ret_type} {
	return (*this.InternalDataMap)["${attr_name}"].(${ret_type})
}
%   endfor
%   for func_name in obj['function_names']:
<%
		func = obj['functions'][func_name]
		ret_type = ptype(func['returns']['type']) if func['returns'] else ''
		argify = lambda a : '{} {}'.format(a['name'], ptype(a['type']))
		args = ', '.join([argify(a) for a in func['arguments']])
%>
func (this ${obj_name}Impl) ${upcase_first(func_name)}(${args})${' ' if ret_type else ''} {
	args := make(map[string]interface{})
%		for arg in func['arguments']:
	args["${arg['name']}"] = ${arg['name']}
%		endfor
	${'return ' if ret_type else ''}this.RunOnServer("${func_name}", args)${('.('+ret_type+')') if ret_type else ''}
}
%   endfor

func defaultInernalDataMapFor${obj_name}() *map[string]interface{} {
	data := make(map[string]interface{})
%	if 'parentClasses' in obj:
%		for i, parent in enumerate(obj['parentClasses']):
	parentData${i} := defaultInernalDataMapFor${parent}()
	for key, value := range *parentData${i} {
		data[key] = value
	}
%		endfor
%	endif
%   for attr_name in obj['attribute_names']:
	data["${attr_name}"] = ${shared['go']['default_value'](obj['attributes'][attr_name]['type'], package_name)}
%   endfor

	return &data
}
% endfor

// -- Namespace -- \\
<% ns = game_name + 'Namespace' %>
type ${ns} struct {}

func (_ ${ns}) Name() string {
	return "${game_name}"
}

func (_ ${ns}) Version() string {
	return "${game_version}"
}

func (_ ${ns}) PlayerName() string {
	return ${package_name}.PlayerName()
}

func (_ ${ns}) CreateGameObject(gameObjectName string) (${package_name}.GameObject, error) {
	switch (gameObjectName) {
% for game_obj_name in game_obj_names:
	case "${game_obj_name}":
		new${game_obj_name} := ${game_obj_name}Impl{}
		new${game_obj_name}.InternalDataMap = defaultInernalDataMapFor${game_obj_name}()
		return &new${game_obj_name}, nil
% endfor
	}
	return nil, errors.New("No game object named " + gameObjectName + " for game ${game_name}")
}

func (_ ${ns}) CreateGame() ${package_name}.Game {
	return &GameImpl{}
}

func (_ ${ns}) CreateAI() ${package_name}.AI {
	return &${package_name}.AI{}
}

func (_ ${ns}) OrderAI(ai *${package_name}.AI, functionName string, args []interface{}) (interface{}, error) {
	switch (functionName) {
% for func_name in ai['function_names']:
<% func = ai['functions'][func_name]
%>	case "${func_name}":
%	for i, arg in enumerate(func['arguments']):
		arg${i} := args[${i}].(ptype(arg['type']))
%	endfor
		return (*ai).${upcase_first(func_name)}(${', '.join('arg{}'.format(i) for i in range(len(func['arguments'])))}), nil
% endfor
	}

	return nil, errors.New("Cannot find functionName "+functionName+" to run in S{game_name} AI")
}

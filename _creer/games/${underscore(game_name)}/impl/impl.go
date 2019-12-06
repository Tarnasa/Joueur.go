// Package impl contains all the ${game_name } implimentation
// logic and structures required by aa client to play with a game server.
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
		parents.append('base.' + obj_name)
	longest_attr_name_len = len(sorted(list(obj['attribute_names']), key=len)[-1])
%>
// -- ${obj_name} -- ${'\\\\'}

// ${obj_name}Impl is the struct that implements the container for ${obj_name} instances in ${game_name}.
type ${obj_name}Impl struct {
%	for parent in parents:
	${parent}Impl
%	endfor
%	if obj_name == 'GameObject':
	game${' ' * (5 + longest_attr_name_len - len('game'))}*GameImpl
%   endif
%	for attr_name in obj['attribute_names']:
%	if not (obj_name == 'GameObject' and attr_name == 'id'):
	${upcase_first(attr_name)}Impl${' ' * (1 + longest_attr_name_len - len(attr_name))}${ptype(obj['attributes'][attr_name]['type'])}
% 	endif
%	endfor
}
%   if obj_name == 'GameObject':

// Game returns a pointer to the ${game_name} Game instance
func (gameObjectImpl *GameObjectImpl) Game() ${package_name}.Game {
	return gameObjectImpl.game
}
%   endif
%   for attr_name in obj['attribute_names']:
<%
		if obj_name == 'GameObject' and attr_name == 'id':
			continue
		attr = obj['attributes'][attr_name]
		ret_type = ptype(attr['type'])
%>
// ${upcase_first(attr_name)} returns ${lowercase_first(shared["go"]["description"](attr))}
func (${lowercase_first(obj_name)}Impl *${obj_name}Impl) ${upcase_first(attr_name)}() ${ret_type} {
	return ${lowercase_first(obj_name)}Impl.${upcase_first(attr_name)}Impl
}
%   endfor
%   for func_name in obj['function_names']:
<%
		func = obj['functions'][func_name]
		ret_type = ptype(func['returns']['type']) if func['returns'] else ''
		argify = lambda a : '{} {}'.format(a['name'], ptype(a['type']))
		args = ', '.join([argify(a) for a in func['arguments']])
%>
// ${upcase_first(func_name)} runs logic that ${lowercase_first(shared["go"]["description"](func))}
func (${lowercase_first(obj_name)}Impl *${obj_name}Impl) ${upcase_first(func_name)}(${args})${' ' if ret_type else ''} {
	${'return ' if ret_type else ''}${lowercase_first(obj_name)}Impl.RunOnServer("${func_name}", map[string]interface{}{
%		for arg in func['arguments']:
		"${arg['name']}": ${arg['name']},
%		endfor
	})${('.('+ret_type+')') if ret_type else ''}
}
%   endfor

// InitImplDefaults initializes safe defaults for all fields in ${obj_name}.
func (${lowercase_first(obj_name)}Impl *${obj_name}Impl) InitImplDefaults() {
%		for i, parent in enumerate(obj['parentClasses'] + [] if not obj_name in ['Game', 'GameObject'] else [obj_name]):
	${lowercase_first(obj_name)}Impl.${parent}Impl.InitImplDefaults()
%		endfor

%   for attr_name in obj['attribute_names']:
	${lowercase_first(obj_name)}Impl.${upcase_first(attr_name)}Impl = ${shared['go']['default_value'](obj['attributes'][attr_name]['type'], package_name)}
%   endfor
}
% endfor

// -- Namespace -- ${'\\\\'}
<% ns = game_name + 'Namespace' %>
// ${ns} is the collection of implimentation logic for the ${game_name} game.
type ${ns} struct {}

// Name returns the name of the ${game_name} game.
func (*${ns}) Name() string {
	return "${game_name}"
}

// Version returns the current version hash as last generated for the ${game_name} game.
func (*${ns}) Version() string {
	return "${game_version}"
}

// PlayerName returns the desired name of the AI in the ${game_name} game.
func (*${ns}) PlayerName() string {
	return ${package_name}.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the ${game_name} game.
func (*${ns}) CreateGameObject(gameObjectName string) (base.GameObject, *base.DeltaMergeableImpl, error) {
	switch (gameObjectName) {
% for game_obj_name in game_obj_names:
	case "${game_obj_name}":
		new${game_obj_name} := ${game_obj_name}Impl{}
		new${game_obj_name}.InitImplDefaults()
		return &new${game_obj_name}, &(new${game_obj_name}.GameObjectImpl.DeltaMergeableImpl), nil
% endfor
	}
	return nil, nil, errors.New("No game object named " + gameObjectName + " for game ${game_name}")
}

// CreateGame is the factory method for Game the ${game_name} game.
func (*${ns}) CreateGame() (base.Game, *base.DeltaMergeableImpl) {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game, &(game.GameImpl.DeltaMergeableImpl)
}

// CreateAI is the factory method for the AI in the ${game_name} game.
func (*${ns}) CreateAI() (base.AI, *base.AIImpl) {
	ai := ${package_name}.AI{}
	return &ai, &ai.AIImpl
}

// OrderAI handles an order for the AI in the ${game_name} game.
func (*${ns}) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*chess.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid chess.AI to order!")
	}
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

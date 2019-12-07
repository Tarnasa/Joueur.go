// Package impl contains all the ${game_name } implimentation
// logic and structures required by aa client to play with a game server.
// To start coding your AI open ./ai.go
package impl
<%include file='functions.noCreer' />
import (
	"errors"
	"joueur/base"
	"joueur/games/${underscore(game_name)}"
)
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
	return ${shared['go']['package_name']}.PlayerName()
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
	ai := ${shared['go']['package_name']}.AI{}
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
		arg${i} := args[${i}].(${shared['go']['type'](arg['type'])})
%	endfor
		return (*ai).${upcase_first(func_name)}(${', '.join('arg{}'.format(i) for i in range(len(func['arguments'])))}), nil
% endfor
	}

	return nil, errors.New("Cannot find functionName "+functionName+" to run in S{game_name} AI")
}

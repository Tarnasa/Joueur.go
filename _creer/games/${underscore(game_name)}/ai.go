package ${lowercase_first(game_name)}
<%include file='impl/functions.noCreer' />
import "joueur/base"

// PlayerName should return the string name of your Player in games it plays.
func PlayerName() string {
${merge(
	'\t// ', 'getName',
	'\treturn "' + game_name + ' Go Player"',
	help=False
)}
}

// AI is your personal AI implimentation.
type AI struct {
	base.AIImpl
${merge(
	'\t// ', 'fields',
	'\t// You can add new fields here',
	help=False
)}
}

// Game returns the instance of the Game this AI is currently playing.
func (ai *AI) Game() Game {
	return ai.AIImpl.Game().(Game)
}

// Player returns the instance of the Player this AI is represented by in the
// game this AI is playing.
func (ai *AI) Player() Player {
	return ai.AIImpl.Player().(Player)
}

// Start is called once the game starts and your AI has a Player and Game.
// You can initialize your AI struct here.
func (ai *AI) Start() {
${merge(
	'\t// ', 'start',
	'\t// pass',
	help=False
)}
}

// GameUpdated is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai *AI) GameUpdated() {
${merge(
	'\t// ', 'game-updated',
	'\t// pass',
	help=False
)}
}

// Ended is called when the game ends, you can clean up your data and dump
// files here if need be.
func (ai *AI) Ended(won bool, reason string) {
${merge(
	'\t// ', 'ended',
	'\t// pass',
	help=False
)}
}

// -- ${game_name} specific AI actions -- ${'\\\\'}
% for function_name in ai['function_names']:
<% function_params = ai['functions'][function_name]%>
${shared['go']['function_top'](function_name, function_params, 'AI')}
${merge(
	'\t// ', function_name,
"""	// Put your game logic here for {}
	return{}
""".format(
		function_name,
		(' ' + shared['go']['default_value'](function_params['returns']['type'])) if function_params['returns'] else ''
	),
	help=False
)}
}
% endfor

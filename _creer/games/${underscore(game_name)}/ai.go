package ${lowercase_first(game_name)}
<%include file='impl/functions.noCreer' />
import "joueur/base"

func PlayerName() string {
${merge(
	'\t// ', 'getName',
	'\treturn "' + game_name + ' Go Player"'
)}
}

type AI struct {
	base.BaseAIImpl

	// The reference to the Game instance this AI is playing.
	Game Game

	// The reference to the Player this AI controls in the Game.
	Player Player
}

// This is called once the game starts and your AI knows its playerID and game.
// You can initialize your AI here.
func (ai AI) Start() {
${merge(
	'\t// ', 'start',
	'\t// pass'
)}
}

// This is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai AI) GameUpdated() {
${merge(
	'\t// ', 'game-updated',
	'\t// pass'
)}
}

// This is called when the game ends, you can clean up your data and dump files here if need be.
//
// @param won True means you won, false means you lost.
// @param reason The human readable string explaining why you won or lost.
func (ai AI) Ended(won bool, reason string) {
${merge(
	'\t// ', 'ended',
	'\t// pass'
)}
}

// Chess specific AI actions
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
	)
)}
}
% endfor

package ${lowercase_first(game_name)}<%include file='functions.noCreer' />
% if obj_key in ['GameObject', 'Game']:
import "joueur/base"

% endif
// ${obj['description']}
type ${obj_key} struct {
% if obj_key in ['GameObject', 'Game']:
	base.Base${obj_key}
% else:
	${obj_key}
% endif
% if obj_key != 'Game':

	// The reference to the Game instance this ${obj_key} is in.
	game Game
% endif
% for attr_name in obj['attribute_names']:
<% attr = obj['attributes'][attr_name] %>
	// ${attr['description']}
	${upcase_first(attr_name)} ${shared['go']['type'](attr['type'])}
% endfor
}
% if len(obj['function_names']) > 0:

// Chess specific ${obj_key} actions
% for function_name in obj['function_names']:
<% function_params = obj['functions'][function_name]%>
${shared['go']['function_top'](function_name, function_params, obj_key)}
	// TODO: somehow we need to client.RunOnServer(args)
	return
}
% endfor
% endif

<%include file='functions.noCreer' /><%
def find_deep_types(type_obj, types):
	if type_obj['name'] in ['list', 'dictionary']:
		types.append(type_obj)
		find_deep_types(type_obj['valueType'], types)
		if type_obj['name'] == 'dictionary':
			find_deep_types(type_obj['keyType'], types)

def find_deep_type_name(type_obj):
	type_name = type_obj['name']
	if type_name == 'list':
		return 'ArrayOf{}'.format(find_deep_type_name(type_obj['valueType']))
	elif type_name == 'dictionary':
		return 'MapOf{}To{}'.format(find_deep_type_name(type_obj['keyType']), find_deep_type_name(type_obj['valueType']))
	else:
		return upcase_first(type_name)

deep_types = []
for obj_name in game_obj_names + ['Game']:
	obj = game if obj_name == 'Game' else game_objs[obj_name]
	for attr_name in obj['attribute_names']:
		find_deep_types(obj['attributes'][attr_name]['type'], deep_types)

name_to_deep_type = dict()
for deep_type in deep_types:
	deep_type_name = find_deep_type_name(deep_type)
	name_to_deep_type[deep_type_name] = deep_type

%>import (
	"joueur/base"
)

type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

% for deep_type_name in sort_dict_keys(name_to_deep_type):
<%
deep_type = name_to_deep_type[deep_type_name]
%>func (deltaMergeImpl DeltaMergeImpl) ${deep_type_name}(state *${shared['go']['type'](deep_type)}, delta interface{}) {
	// TODO: do
}
% endfor

package impl
<%include file='functions.noCreer' /><%
package_name = lowercase_first(game_name)

def find_deep_types(type_obj, types):
	if shared['go']['is_type_deep'](type_obj):
		types.append(type_obj)
		find_deep_types(type_obj['valueType'], types)
		if type_obj['name'] == 'dictionary':
			find_deep_types(type_obj['keyType'], types)

deep_types = []
for obj_name in game_obj_names + ['Game']:
	obj = game if obj_name == 'Game' else game_objs[obj_name]
	for attr_name in obj['attribute_names']:
		find_deep_types(obj['attributes'][attr_name]['type'], deep_types)

name_to_deep_type = dict()
for deep_type in deep_types:
	deep_type_name = shared['go']['find_deep_type_name'](deep_type)
	name_to_deep_type[deep_type_name] = deep_type

%>
import (
	"joueur/base"
	"joueur/games/${package_name}"
)

type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- ${game_name} Game Object Deltas -- ${'\\\\'}

% for game_obj_name in game_obj_names:
// ${game_obj_name} attempts to return the instance of ${game_obj_name}
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) ${game_obj_name}(delta interface{}) ${package_name}.${game_obj_name} {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	as${game_obj_name}, is${game_obj_name} := baseGameObject.(${package_name}.${game_obj_name})
	if !is${game_obj_name} {
		(*deltaMergeImpl).CannotConvertDeltaTo("${package_name}.${game_obj_name}", delta)
	}

	return as${game_obj_name}
}

% endfor
// -- Deep Deltas -- ${'\\\\'}

% for deep_type_name in sort_dict_keys(name_to_deep_type):
<%
deep_type = name_to_deep_type[deep_type_name]
is_list = deep_type['name'] == 'list'

go_type = shared['go']['type'](deep_type, package_name)
go_value_type = shared['go']['type'](deep_type['valueType'], package_name)

value_type = deep_type['valueType']

valueCall = 'deltaMergeImpl.{}({}deltaValue)'.format(
	shared['go']['find_deep_type_name'](value_type),
	'&newArray[index], ' if shared['go']['is_type_deep'](value_type) else ''
)
%>// ${deep_type_name} delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ${deep_type_name}(state *${go_type}, delta interface{}) ${go_type} {
%	if is_list:
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make(${go_type}, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = ${valueCall}
	}
	return newArray
%	else: # it's a map, and for now assume key_type is string because that's all that is really supported
	deltaMap := (*deltaMergeImpl).ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if (*deltaMergeImpl).IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = ${valueCall}
		}
	}
	return *state
%	endif
}

% endfor
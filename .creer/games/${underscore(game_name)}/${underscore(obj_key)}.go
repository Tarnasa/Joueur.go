package ${lowercase_first(game_name)}
<%include file='functions.noCreer' />
${shared['go']['imports_for'](obj_key)}${shared['go']['interface_for'](obj_key, obj)}
% if 'tiled' in parent_data_names and obj_key == 'Tile':
// TileDirections are all the direction strings that Tile's neighbors can be
// in.
var TileDirections = [...]string{"North", "South", "East", "West"}
% endif

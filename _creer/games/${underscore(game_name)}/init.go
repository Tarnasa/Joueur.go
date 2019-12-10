// Package ${lowercase_first(game_name)} contains all the interfaces and AI
// that are required to represent and play ${game_name}.
package ${lowercase_first(game_name)}

// If you need to add code to initialize the ENTIRE ${game_name} package,
// do so here
func init() {
${merge(
	'\t// ', 'init',
	'\t// package initialization logic can go here',
	help=False
)}
}

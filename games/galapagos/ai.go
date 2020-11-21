package galapagos

import "joueur/base"

// PlayerName should return the string name of your Player in games it plays.
func PlayerName() string {
	return "Galapagos Go Player"
}

// AI is your personal AI implimentation.
type AI struct {
	base.AIImpl
	// You can add new fields here
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
	// pass
}

// GameUpdated is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai *AI) GameUpdated() {
	// pass
}

// Ended is called when the game ends, you can clean up your data and dump
// files here if need be.
func (ai *AI) Ended(won bool, reason string) {
	// pass
}

// -- Galapagos specific AI actions -- \\

// RunTurn this is called every time it is this AI.player's turn.
func (ai *AI) RunTurn() bool {
	// Put your game logic here for runTurn
	return true

}

// -- Tiled Game Utils -- \\

// findPath is a very basic path finding algorithm (Breadth First Search),
// that when given a starting Tile, will return a valid path to the goal Tile.
// If none is found an empty array will be returned.
func (ai *AI) findPath(start Tile, goal Tile) []Tile {
	if start == goal {
		// no need to make a path to here...
		return []Tile{}
	}

	// Queue of the tiles that will have their neighbors searched for 'goal',
	// with start as the first tile to have its neighbors searched.
	fringe := []Tile{start}

	// How we got to each tile that went into the fringe.
	cameFrom := map[Tile]Tile{
		start: start, // so we don't do back to the start when searching
	}

	// keep exploring neighbors of neighbors... until there are no more.
	for len(fringe) > 0 {
		// the tile we are currently exploring.
		inspect := fringe[len(fringe)-1] // pop off the last time
		fringe = fringe[:len(fringe)-1]

		// cycle through the tile's neighbors.
		for _, neighbor := range inspect.GetNeighbors() {
			// if we found the goal, we have the path!
			if neighbor == goal {
				// Follow the path backward to the start from the goal and
				// return it.
				path := []Tile{goal}

				// Starting at the tile we are currently at, insert them
				// retracing our steps till we get to the starting tile
				for inspect != start {
					// add inspect to the front of the path
					path = append([]Tile{inspect}, path...)
					inspect = cameFrom[inspect]
				}
				return path
			}
			// else we did not find the goal, so enqueue this tile's
			// neighbors to be inspected

			// if the tile exists, has not been explored or added to the
			// fringe yet, and it is pathable
			_, neighborInCameFrom := cameFrom[neighbor]
			if !neighborInCameFrom && neighbor.IsPathable() {
				// add it to the tiles to be explored and add where it came
				// from for path reconstruction.
				fringe = append(fringe, neighbor)
				cameFrom[neighbor] = inspect
			}
		}
	}

	// if you're here, that means that there was not a path to get to where
	// you want to go; in that case, we'll just return an empty path.
	return []Tile{}
}

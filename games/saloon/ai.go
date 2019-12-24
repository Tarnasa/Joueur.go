package saloon

import (
	"fmt"
	"joueur/base"
	"math/rand"
	"time"
)

// PlayerName should return the string name of your Player in games it plays.
func PlayerName() string {
	return "Saloon Go Player"
}

// AI is your personal AI implimentation.
type AI struct {
	base.AIImpl
	// You can add new fields here

	random *rand.Rand
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
	seed := rand.NewSource(time.Now().Unix())
	ai.random = rand.New(seed)
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

// -- Saloon specific AI actions -- \\

// RunTurn this is called every time it is this AI.Player()'s turn.
func (ai *AI) RunTurn() bool {
	// This is "ShellAI", some basic code we've provided that does
	// everything in the game for demo purposed, but poorly so you
	// can get to optimizing or overwriting it ASAP
	//
	// ShellAI does a few things:
	// 1. Tries to spawn a new Cowboy
	// 2. Tries to move to a Piano
	// 3. Tries to play a Piano
	// 4. Tries to act

	fmt.Println("Start of my turn:", ai.Game().CurrentTurn())

	//--- 1. Try to spawn a Cowboy --\\

	// Randomly select a job.
	callInJobIndex := ai.random.Intn(len(ai.Game().Jobs()))
	callInJob := ai.Game().Jobs()[callInJobIndex]
	var jobCount int64 = 0
	for _, myCowboy := range ai.Player().Cowboys() {
		if !myCowboy.IsDead() && myCowboy.Job() == callInJob {
			jobCount++
		}
	}

	// Call in the new cowboy with that job if there aren't too many
	//   cowboys with that job already.
	if ai.Player().YoungGun().CanCallIn() && jobCount < ai.Game().MaxCowboysPerJob() {
		fmt.Println("1. Calling in:", callInJob)
		ai.Player().YoungGun().CallIn(callInJob)
	}

	// for steps 2, 3, and 4 we will use this cowboy:
	var activeCowboy Cowboy = nil
	for _, myCowboy := range ai.Player().Cowboys() {
		if !myCowboy.IsDead() {
			activeCowboy = myCowboy
			break
		}
	}

	// Now let's use them
	if activeCowboy != nil {
		//--- 2. Try to move to a Piano ---\\

		// find a piano
		var piano Furnishing = nil
		for _, furnishing := range ai.Game().Furnishings() {
			if furnishing.IsPiano() && !furnishing.IsDestroyed() {
				piano = furnishing
				break
			}
		}

		// There will always be pianos or the game will end. No need to check for existence.
		// Attempt to move toward the piano by finding a path.
		if activeCowboy.CanMove() && !activeCowboy.IsDead() {
			fmt.Println("Trying to do stuff with", activeCowboy)

			// find a path from the Tile this cowboy is on to the tile the piano is on
			path := ai.findPath(activeCowboy.Tile(), piano.Tile())

			// if there is a path, move to it
			//	  length 0 means no path could be found to the tile
			//	  length 1 means the piano is adjacent, and we can't move onto the same tile as the piano
			if len(path) > 1 {
				fmt.Println("2. Moving to", path[0])
				activeCowboy.Move(path[0])
			}
		}

		//--- 3. Try to play a piano ---\\\

		// make sure the cowboy is alive and is not busy
		if !activeCowboy.IsDead() && activeCowboy.TurnsBusy() == 0 {
			// look at all the neighboring (adjacent) tiles, and if they have a piano, play it
			neighbors := activeCowboy.Tile().GetNeighbors()
			for _, neighbor := range neighbors {
				// if the neighboring tile has a piano
				if neighbor.Furnishing() != nil && neighbor.Furnishing().IsPiano() {
					// then play it
					fmt.Println("3. Playing Furnishing (piano) #", neighbor.Furnishing().ID())
					activeCowboy.Play(neighbor.Furnishing())
					break
				}
			}
		}

		//--- 4. Try to act ---\\

		// make sure the cowboy is alive and is not busy
		if !activeCowboy.IsDead() && activeCowboy.TurnsBusy() == 0 {
			// Get a random neighboring tile.
			neighbors := activeCowboy.Tile().GetNeighbors()
			randomNeighborIndex := ai.random.Intn(len(neighbors))
			randomNeighbor := neighbors[randomNeighborIndex]

			// Based on job, act accordingly.
			switch activeCowboy.Job() {
			case "Bartender":
				// Bartenders throw Bottles in a direction, and the Bottle makes cowboys drunk which causes them to walk in random directions
				// so throw the bottle on a random neighboring tile, and make drunks move in a random direction
				directionIndex := ai.random.Intn(len(TileDirections))
				direction := TileDirections[directionIndex]

				fmt.Println("4. Bartender acting on ", randomNeighbor, "with drunkDirection", direction)
				activeCowboy.Act(randomNeighbor, direction)
			case "Brawler":
				// Brawlers cannot act, they instead automatically attack all neighboring tiles on the end of their owner's turn.
				fmt.Println("4. Brawlers cannot act.")
			case "Sharpshooter":
				// Sharpshooters build focus by standing still, they can then act(tile) on a neighboring tile to fire in that direction
				if activeCowboy.Focus() > 0 {
					fmt.Println("4. Sharpshooter acting on", randomNeighbor)
					activeCowboy.Act(randomNeighbor, "") // fire in a random direction
				} else {
					fmt.Println("4. Sharpshooter doesn't have enough focus. (focus ==", activeCowboy.Focus(), ")")
				}
			}
		}
	}

	fmt.Println("Ending my turn.")

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

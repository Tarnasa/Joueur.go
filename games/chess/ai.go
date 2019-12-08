package chess

import "joueur/base"

import "fmt"

func PlayerName() string {
	// <<-- Creer-Merge: getName -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	return "Chess Go Player"
	// <<-- /Creer-Merge: getName -->>
}

type AI struct {
	base.AIImpl
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

// This is called once the game starts and your AI knows its playerID and game.
// You can initialize your AI here.
func (ai *AI) Start() {
	// <<-- Creer-Merge: start -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: start -->>
}

// This is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai *AI) GameUpdated() {
	// <<-- Creer-Merge: game-updated -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: game-updated -->>
}

// This is called when the game ends, you can clean up your data and dump files here if need be.
//
// @param won True means you won, false means you lost.
// @param reason The human readable string explaining why you won or lost.
func (ai *AI) Ended(won bool, reason string) {
	// <<-- Creer-Merge: ended -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: ended -->>
}

// Chess specific AI actions

// This is called every time it is this AI.player's turn to make a move.
func (ai *AI) MakeMove() string {
	// <<-- Creer-Merge: makeMove -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// Put your game logic here for makeMove
	fmt.Println("make move..., Game:", ai.Game().Fen(), "Player:", ai.Player())
	return "b5"
	// <<-- /Creer-Merge: makeMove -->>
}

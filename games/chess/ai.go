package chess

import "joueur/base"

type AI struct {
	base.BaseAI

	// The reference to the Game instance this AI is playing.
	Game *Game

	// The reference to the Player this AI controls in the Game.
	player *Player
}

func (ai AI) GetPlayerName() string {
	// <<-- Creer-Merge: getName -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	return "Chess Go Player"
	// <<-- /Creer-Merge: getName -->>
}

// This is called once the game starts and your AI knows its playerID and game.
// You can initialize your AI here.
func (ai AI) Start() {
	// <<-- Creer-Merge: start -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: start -->>
}

// This is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai AI) GameUpdated() {
	// <<-- Creer-Merge: game-updated -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: game-updated -->>
}

// This is called when the game ends, you can clean up your data and dump files here if need be.
//
// @param won True means you won, false means you lost.
// @param reason The human readable string explaining why you won or lost.
func (ai AI) Ended(won bool, reason string) {
	// <<-- Creer-Merge: ended -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// pass
	// <<-- /Creer-Merge: ended -->>
}

// Chess specific AI actions

// This is called every time it is this AI.player's turn to make a move.
func (ai AI) MakeMove() string {
	// <<-- Creer-Merge: makeMove -->> - Code you add between this comment and the end comment will be preserved between Creer re-runs.
	// Put your game logic here for makeMove
	return ""
	// <<-- /Creer-Merge: makeMove -->>
}

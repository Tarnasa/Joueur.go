package base

import (
	"github.com/fatih/color"
)

// AISettings is kind of hacky; basically exposing a global for all AI structs,
// however only 1 AI should ever be constructed per game so should be an ok assumption
var AISettings map[string]([]string)

// AI is the base interface all game AIs should implement to correctly
// interface and play their game.
type AI interface {
	Start()
	Ended(bool, string)
	GameUpdated()
	// Invalid()
}

// AIImpl is the implimentation struct for the AI interface
type AIImpl struct {
	game   Game
	player Player
}

// Game returns the game this [base] AI is in
func (ai *AIImpl) Game() Game {
	return ai.game
}

// Player returns the game this [base] AI is in
func (ai *AIImpl) Player() Player {
	return ai.player
}

// Start is called once the game starts and your AI knows its Player and
// Game. You can initialize your AI here.
func (ai AIImpl) Start() {
	// pass
}

// Ended is called when the game ends.
// You can clean up your data and dump files here if need be.
func (ai AIImpl) Ended(won bool, reason string) {
	// pass
}

// GameUpdated is called every time the game's state updates,
// so if you are tracking anything you can update it here.
func (ai AIImpl) GameUpdated() {
	// pass
}

// Invalid is automatically called after this AI sends some run that is
// has arguments that are invalidated for some reason.
func (ai AIImpl) Invalid(message string) {
	color.Yellow("Invalid: " + message)
}

// GetSetting gets an AI setting passed to the program via the --aiSettings
// flag. If the flag was set it will be returned as a string value, undefined otherwise.
func (ai AIImpl) GetSetting(key string) ([]string, bool) {
	setting, ok := AISettings[key]
	return setting, ok
}

// InjectIntoAI injects into an AI its "private" player and game
func InjectIntoAI(ai *AIImpl, game Game, player Player) {
	ai.game = game
	ai.player = player
}

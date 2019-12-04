package base

import (
	"github.com/fatih/color"
)

// AISettings - kind of hacky, basically exposing a global for all AI structs,
// however only 1 AI should ever be constructed per game so should be an ok assumption
var AISettings map[string]([]string)

type BaseAIImpl struct {
	Game   BaseGame
	Player BasePlayer
}

type BaseAI interface {
	Start()
	Ended(bool, string)
	GameUpdated()
	// Invalid()
}

func (ai BaseAIImpl) GetPlayerName() string {
	return "Go Player"
}

func (ai BaseAIImpl) Start() {
	// pass
}

func (ai BaseAIImpl) Ended(won bool, reason string) {
	// pass
}

func (ai BaseAIImpl) GameUpdated() {
	// pass
}

func (ai BaseAIImpl) Invalid(message string) {
	color.Yellow("Invalid: " + message)
}

func (ai BaseAIImpl) GetSetting(key string) ([]string, bool) {
	setting, ok := AISettings[key]
	return setting, ok
}

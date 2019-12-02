package base

import (
	"github.com/fatih/color"
)

// kind of hacky, basically exposing a global for all AI structs,
// however only 1 AI should ever be constructed per game so should be an ok assumption
var AISettings map[string]([]string)

type BaseAI struct {
	Game   *BaseGame
	Player *BasePlayer
}

type InterfaceAI interface {
	GetPlayerName() string
	Start()
	Ended(bool, string)
	GameUpdated()
	// Invalid()
}

func (ai BaseAI) GetPlayerName() string {
	return "Go Player"
}

func (ai BaseAI) Start() {
	// pass
}

func (ai BaseAI) Ended(won bool, reason string) {
	// pass
}

func (ai BaseAI) GameUpdated() {
	// pass
}

func (ai BaseAI) Invalid(message string) {
	color.Yellow("Invalid: " + message)
}

func (ai BaseAI) GetSetting(key string) ([]string, bool) {
	setting, ok := AISettings[key]
	return setting, ok
}

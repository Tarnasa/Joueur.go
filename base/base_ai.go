package base

import (
	"github.com/fatih/color"
)

type BaseAI struct {
	Game   BaseGame
	Player BasePlayer

	Settings map[string]([]string)
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
	setting, ok := ai.Settings[key]
	return setting, ok
}

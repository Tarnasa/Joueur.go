package gamemanager

import (
	"joueur/base"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"

	"errors"
	"net/url"
	"reflect"
)

type GameManager struct {
	ServerConstants client.ServerConstants
	GameNamespace   *games.GameNamespace
	InterfaceAI     *base.InterfaceAI
	ReflectAI       *reflect.Value

	reflectGame *reflect.Value
}

func New(gameManager *GameManager, aiSettings string) *GameManager {
	reflectGame := reflect.New((*gameManager.GameNamespace).GameType)
	gameManager.reflectGame = &reflectGame

	if !reflectGame.IsValid() {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not create Game instance for "+(*(gameManager.GameNamespace)).Name),
		)
	}
	settings := make(map[string]([]string))
	parsedSettings, parseErr := url.ParseQuery(aiSettings)
	if parseErr != nil {
		errorhandler.HandleError(
			errorhandler.InvalidArgs,
			parseErr,
			"Error occured while parsing AI Settings query string. Ensure the format is correct",
		)
	}

	for key, value := range parsedSettings {
		settings[key] = value
	}

	rai := (*gameManager.ReflectAI).Elem()
	rai.FieldByName("Settings").Set(reflect.ValueOf(settings))
	// rai.FieldByName("Game").Set(reflectGame.Elem())

	return gameManager
}

func (gameManager GameManager) Start(playerID string) {
	// TODO: set player in ai by this ID
}

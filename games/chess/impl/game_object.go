package impl

import (
	"joueur/base"
	"joueur/games/chess"
)

// GameObjectImpl is the struct that implements the container for GameObject instances in Chess.
type GameObjectImpl struct {
	base.GameObjectImpl
	game               *GameImpl
	implgameObjectName string
	impllogs           []string
}

// Game returns a pointer to the Chess Game instance
func (gameObjectImpl *GameObjectImpl) Game() chess.Game {
	return gameObjectImpl.game
}

// GameObjectName returns string representing the top level Class that this game object is an instance of. Used for reflection to create new instances on clients, but exposed for convenience should AIs want this data.
func (gameObjectImpl *GameObjectImpl) GameObjectName() string {
	return gameObjectImpl.GameObjectNameImpl
}

// Logs returns any strings logged will be stored here. Intended for debugging.
func (gameObjectImpl *GameObjectImpl) Logs() []string {
	return gameObjectImpl.LogsImpl
}

// Log runs logic that adds a message to this GameObject's logs. Intended for your own debugging purposes, as strings stored here are saved in the gamelog.
func (gameObjectImpl *GameObjectImpl) Log(message string) {
	gameObjectImpl.RunOnServer("log", map[string]interface{}{
		"message": message,
	})
}

// InitImplDefaults initializes safe defaults for all fields in GameObject.
func (gameObjectImpl *GameObjectImpl) InitImplDefaults() {
	gameObjectImpl.GameObjectImpl.InitImplDefaults()

	gameObjectImpl.implgameObjectName = ""
	gameObjectImpl.implid = ""
	gameObjectImpl.impllogs = make([]string, 0)
}

// DeltaMerge merged the delta for a given attribute in GameObject.
func (gameObjectImpl *GameObjectImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) {
	switch(attribute) {
	case "gameObjectName":
		(*gameObjectImpl).gameObjectNameImpl = deltaMerge.String(delta)
		break
	case "id":
		(*gameObjectImpl).idImpl = deltaMerge.String(delta)
		break
	case "logs":
		(*gameObjectImpl).logsImpl = deltaMerge.ArrayOfString((*gameObjectImpl).logsImpl, )
		break
	}
}

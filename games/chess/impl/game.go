package impl

import (
	"joueur/base"
	"joueur/games/chess"
)

// GameImpl is the struct that implements the container for Game instances in Chess.
type GameImpl struct {
	base.GameImpl
	implfen         string
	implgameObjects map[string]chess.GameObject
	implhistory     []string
	implplayers     []chess.Player
	implsession     string
}

// Fen returns forsyth-Edwards Notation (fen), a notation that describes the game board state.
func (gameImpl *GameImpl) Fen() string {
	return gameImpl.FenImpl
}

// GameObjects returns a mapping of every game object's ID to the actual game object. Primarily used by the server and client to easily refer to the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]chess.GameObject {
	return gameImpl.GameObjectsImpl
}

// History returns the array of [known] moves that have occurred in the game, in Standard Algebraic Notation (SAN) format. The first element is the first move, with the last being the most recent.
func (gameImpl *GameImpl) History() []string {
	return gameImpl.HistoryImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []chess.Player {
	return gameImpl.PlayersImpl
}

// Session returns a unique identifier for the game instance that is being played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.SessionImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.implfen = ""
	gameImpl.implgameObjects = make(map[string]chess.GameObject)
	gameImpl.implhistory = make([]string, 0)
	gameImpl.implplayers = make([]chess.Player, 0)
	gameImpl.implsession = ""
}

// DeltaMerge merged the delta for a given attribute in Game.
func (gameImpl *GameImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) {
	switch(attribute) {
	case "fen":
		(*gameImpl).fenImpl = deltaMerge.String(delta)
		break
	case "gameObjects":
		(*gameImpl).gameObjectsImpl = deltaMerge.MapOfStringToGameObject((*gameImpl).gameObjectsImpl, )
		break
	case "history":
		(*gameImpl).historyImpl = deltaMerge.ArrayOfString((*gameImpl).historyImpl, )
		break
	case "players":
		(*gameImpl).playersImpl = deltaMerge.ArrayOfPlayer((*gameImpl).playersImpl, )
		break
	case "session":
		(*gameImpl).sessionImpl = deltaMerge.String(delta)
		break
	}
}

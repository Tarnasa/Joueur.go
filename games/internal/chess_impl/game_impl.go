package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/chess"
)

// GameImpl is the struct that implements the container for Game instances
// in Chess.
type GameImpl struct {
	base.GameImpl

	fenImpl         string
	gameObjectsImpl map[string]chess.GameObject
	historyImpl     []string
	playersImpl     []chess.Player
	sessionImpl     string
}

// Fen returns forsyth-Edwards Notation (fen), a notation that describes
// the game board state.
func (gameImpl *GameImpl) Fen() string {
	return gameImpl.fenImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]chess.GameObject {
	return gameImpl.gameObjectsImpl
}

// History returns the array of [known] moves that have occurred in the
// game, in Universal Chess Inferface (UCI) format. The first element is
// the first move, with the last element being the most recent.
func (gameImpl *GameImpl) History() []string {
	return gameImpl.historyImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []chess.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.fenImpl = ""
	gameImpl.gameObjectsImpl = make(map[string]chess.GameObject)
	gameImpl.historyImpl = []string{}
	gameImpl.playersImpl = []chess.Player{}
	gameImpl.sessionImpl = ""
}

// DeltaMerge merges the delta for a given attribute in Game.
func (gameImpl *GameImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := gameImpl.GameImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	chessDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'chess.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "fen":
		gameImpl.fenImpl = chessDeltaMerge.String(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = chessDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "history":
		gameImpl.historyImpl = chessDeltaMerge.ArrayOfString(&gameImpl.historyImpl, delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = chessDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = chessDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

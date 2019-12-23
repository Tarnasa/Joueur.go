package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/checkers"
)

// GameImpl is the struct that implements the container for Game instances
// in Checkers.
type GameImpl struct {
	base.GameImpl

	boardHeightImpl        int64
	boardWidthImpl         int64
	checkerMovedImpl       checkers.Checker
	checkerMovedJumpedImpl bool
	checkersImpl           []checkers.Checker
	currentPlayerImpl      checkers.Player
	currentTurnImpl        int64
	gameObjectsImpl        map[string]checkers.GameObject
	maxTurnsImpl           int64
	playersImpl            []checkers.Player
	sessionImpl            string
	timeAddedPerTurnImpl   int64
}

// BoardHeight returns the height of the board for the Y component of a
// checker.
func (gameImpl *GameImpl) BoardHeight() int64 {
	return gameImpl.boardHeightImpl
}

// BoardWidth returns the width of the board for X component of a checker.
func (gameImpl *GameImpl) BoardWidth() int64 {
	return gameImpl.boardWidthImpl
}

// CheckerMoved returns the checker that last moved and must be moved
// because only one checker can move during each players turn.
//
// Value can be returned as a nil pointer.
func (gameImpl *GameImpl) CheckerMoved() checkers.Checker {
	return gameImpl.checkerMovedImpl
}

// CheckerMovedJumped returns if the last checker that moved jumped,
// meaning it can move again.
func (gameImpl *GameImpl) CheckerMovedJumped() bool {
	return gameImpl.checkerMovedJumpedImpl
}

// Checkers returns all the checkers currently in the game.
func (gameImpl *GameImpl) Checkers() []checkers.Checker {
	return gameImpl.checkersImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() checkers.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]checkers.GameObject {
	return gameImpl.gameObjectsImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []checkers.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.boardHeightImpl = 0
	gameImpl.boardWidthImpl = 0
	gameImpl.checkerMovedImpl = nil
	gameImpl.checkerMovedJumpedImpl = true
	gameImpl.checkersImpl = []checkers.Checker{}
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]checkers.GameObject)
	gameImpl.maxTurnsImpl = 0
	gameImpl.playersImpl = []checkers.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.timeAddedPerTurnImpl = 0
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

	checkersDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'checkers.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "boardHeight":
		gameImpl.boardHeightImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	case "boardWidth":
		gameImpl.boardWidthImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	case "checkerMoved":
		gameImpl.checkerMovedImpl = checkersDeltaMerge.Checker(delta)
		return true, nil
	case "checkerMovedJumped":
		gameImpl.checkerMovedJumpedImpl = checkersDeltaMerge.Boolean(delta)
		return true, nil
	case "checkers":
		gameImpl.checkersImpl = checkersDeltaMerge.ArrayOfChecker(&gameImpl.checkersImpl, delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = checkersDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = checkersDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = checkersDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = checkersDeltaMerge.String(delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

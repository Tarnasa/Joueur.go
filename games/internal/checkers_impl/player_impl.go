package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/checkers"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Checkers.
type PlayerImpl struct {
	GameObjectImpl

	checkersImpl      []checkers.Checker
	clientTypeImpl    string
	lostImpl          bool
	nameImpl          string
	opponentImpl      checkers.Player
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	wonImpl           bool
	yDirectionImpl    int64
}

// Checkers returns all the checkers currently in the game owned by this
// player.
func (playerImpl *PlayerImpl) Checkers() []checkers.Checker {
	return playerImpl.checkersImpl
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() checkers.Player {
	return playerImpl.opponentImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// YDirection returns the direction your checkers must go along the y-axis
// until kinged.
func (playerImpl *PlayerImpl) YDirection() int64 {
	return playerImpl.yDirectionImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.checkersImpl = []checkers.Checker{}
	playerImpl.clientTypeImpl = ""
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.wonImpl = true
	playerImpl.yDirectionImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := playerImpl.GameObjectImpl.DeltaMerge(
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
	case "checkers":
		playerImpl.checkersImpl = checkersDeltaMerge.ArrayOfChecker(&playerImpl.checkersImpl, delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = checkersDeltaMerge.String(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = checkersDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = checkersDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = checkersDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = checkersDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = checkersDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = checkersDeltaMerge.Float(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = checkersDeltaMerge.Boolean(delta)
		return true, nil
	case "yDirection":
		playerImpl.yDirectionImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

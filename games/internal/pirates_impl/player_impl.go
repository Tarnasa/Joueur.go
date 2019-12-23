package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Pirates.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	goldImpl          int64
	infamyImpl        int64
	lostImpl          bool
	nameImpl          string
	opponentImpl      pirates.Player
	portImpl          pirates.Port
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	unitsImpl         []pirates.Unit
	wonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Gold returns the amount of gold this Player has in reserve.
func (playerImpl *PlayerImpl) Gold() int64 {
	return playerImpl.goldImpl
}

// Infamy returns the amount of infamy this Player has.
func (playerImpl *PlayerImpl) Infamy() int64 {
	return playerImpl.infamyImpl
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
func (playerImpl *PlayerImpl) Opponent() pirates.Player {
	return playerImpl.opponentImpl
}

// Port returns the Port owned by this Player.
func (playerImpl *PlayerImpl) Port() pirates.Port {
	return playerImpl.portImpl
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

// Units returns every Unit owned by this Player.
func (playerImpl *PlayerImpl) Units() []pirates.Unit {
	return playerImpl.unitsImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.goldImpl = 0
	playerImpl.infamyImpl = 0
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.portImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.unitsImpl = []pirates.Unit{}
	playerImpl.wonImpl = true
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

	piratesDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'pirates.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "gold":
		playerImpl.goldImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "infamy":
		playerImpl.infamyImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = piratesDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = piratesDeltaMerge.Player(delta)
		return true, nil
	case "port":
		playerImpl.portImpl = piratesDeltaMerge.Port(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "units":
		playerImpl.unitsImpl = piratesDeltaMerge.ArrayOfUnit(&playerImpl.unitsImpl, delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = piratesDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

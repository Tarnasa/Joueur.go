package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Galapagos.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	creaturesImpl     []galapagos.Creature
	lostImpl          bool
	nameImpl          string
	opponentImpl      galapagos.Player
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	totalHealthImpl   int64
	wonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Creatures returns every creature owned by this Player.
func (playerImpl *PlayerImpl) Creatures() []galapagos.Creature {
	return playerImpl.creaturesImpl
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
func (playerImpl *PlayerImpl) Opponent() galapagos.Player {
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

// TotalHealth returns the combined health of all creatures this player had
// at the beginning of this turn.
func (playerImpl *PlayerImpl) TotalHealth() int64 {
	return playerImpl.totalHealthImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.creaturesImpl = []galapagos.Creature{}
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.totalHealthImpl = 0
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

	galapagosDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'galapagos.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = galapagosDeltaMerge.String(delta)
		return true, nil
	case "creatures":
		playerImpl.creaturesImpl = galapagosDeltaMerge.ArrayOfCreature(&playerImpl.creaturesImpl, delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = galapagosDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = galapagosDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = galapagosDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = galapagosDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = galapagosDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = galapagosDeltaMerge.Float(delta)
		return true, nil
	case "totalHealth":
		playerImpl.totalHealthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = galapagosDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/catastrophe"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Catastrophe.
type PlayerImpl struct {
	GameObjectImpl

	catImpl           catastrophe.Unit
	clientTypeImpl    string
	foodImpl          int64
	lostImpl          bool
	nameImpl          string
	opponentImpl      catastrophe.Player
	reasonLostImpl    string
	reasonWonImpl     string
	structuresImpl    []catastrophe.Structure
	timeRemainingImpl float64
	unitsImpl         []catastrophe.Unit
	upkeepImpl        int64
	wonImpl           bool
}

// Cat returns the overlord cat Unit owned by this Player.
func (playerImpl *PlayerImpl) Cat() catastrophe.Unit {
	return playerImpl.catImpl
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Food returns the amount of food owned by this player.
func (playerImpl *PlayerImpl) Food() int64 {
	return playerImpl.foodImpl
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
func (playerImpl *PlayerImpl) Opponent() catastrophe.Player {
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

// Structures returns every Structure owned by this Player.
func (playerImpl *PlayerImpl) Structures() []catastrophe.Structure {
	return playerImpl.structuresImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Units returns every Unit owned by this Player.
func (playerImpl *PlayerImpl) Units() []catastrophe.Unit {
	return playerImpl.unitsImpl
}

// Upkeep returns the total upkeep of every Unit owned by this Player. If
// there isn't enough food for every Unit, all Units become starved and do
// not consume food.
func (playerImpl *PlayerImpl) Upkeep() int64 {
	return playerImpl.upkeepImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.catImpl = nil
	playerImpl.clientTypeImpl = ""
	playerImpl.foodImpl = 0
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.structuresImpl = []catastrophe.Structure{}
	playerImpl.timeRemainingImpl = 0
	playerImpl.unitsImpl = []catastrophe.Unit{}
	playerImpl.upkeepImpl = 0
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

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "cat":
		playerImpl.catImpl = catastropheDeltaMerge.Unit(delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "food":
		playerImpl.foodImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = catastropheDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = catastropheDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "structures":
		playerImpl.structuresImpl = catastropheDeltaMerge.ArrayOfStructure(&playerImpl.structuresImpl, delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "units":
		playerImpl.unitsImpl = catastropheDeltaMerge.ArrayOfUnit(&playerImpl.unitsImpl, delta)
		return true, nil
	case "upkeep":
		playerImpl.upkeepImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = catastropheDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

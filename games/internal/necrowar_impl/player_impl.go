package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Necrowar.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	goldImpl          int64
	healthImpl        int64
	homeBaseImpl      []necrowar.Tile
	lostImpl          bool
	manaImpl          int64
	nameImpl          string
	opponentImpl      necrowar.Player
	reasonLostImpl    string
	reasonWonImpl     string
	sideImpl          []necrowar.Tile
	timeRemainingImpl float64
	towersImpl        []necrowar.Tower
	unitsImpl         []necrowar.Unit
	wonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Gold returns the amount of gold this Player has.
func (playerImpl *PlayerImpl) Gold() int64 {
	return playerImpl.goldImpl
}

// Health returns the amount of health remaining for this player's main
// unit.
func (playerImpl *PlayerImpl) Health() int64 {
	return playerImpl.healthImpl
}

// HomeBase returns the tile that the home base is located on.
func (playerImpl *PlayerImpl) HomeBase() []necrowar.Tile {
	return playerImpl.homeBaseImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Mana returns the amount of mana this player has.
func (playerImpl *PlayerImpl) Mana() int64 {
	return playerImpl.manaImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() necrowar.Player {
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

// Side returns all tiles that this player can build on and move workers
// on.
func (playerImpl *PlayerImpl) Side() []necrowar.Tile {
	return playerImpl.sideImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Towers returns every Tower owned by this player.
func (playerImpl *PlayerImpl) Towers() []necrowar.Tower {
	return playerImpl.towersImpl
}

// Units returns every Unit owned by this Player.
func (playerImpl *PlayerImpl) Units() []necrowar.Unit {
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
	playerImpl.healthImpl = 0
	playerImpl.homeBaseImpl = []necrowar.Tile{}
	playerImpl.lostImpl = true
	playerImpl.manaImpl = 0
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.sideImpl = []necrowar.Tile{}
	playerImpl.timeRemainingImpl = 0
	playerImpl.towersImpl = []necrowar.Tower{}
	playerImpl.unitsImpl = []necrowar.Unit{}
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

	necrowarDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'necrowar.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "gold":
		playerImpl.goldImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "health":
		playerImpl.healthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "homeBase":
		playerImpl.homeBaseImpl = necrowarDeltaMerge.ArrayOfTile(&playerImpl.homeBaseImpl, delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "mana":
		playerImpl.manaImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = necrowarDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "side":
		playerImpl.sideImpl = necrowarDeltaMerge.ArrayOfTile(&playerImpl.sideImpl, delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = necrowarDeltaMerge.Float(delta)
		return true, nil
	case "towers":
		playerImpl.towersImpl = necrowarDeltaMerge.ArrayOfTower(&playerImpl.towersImpl, delta)
		return true, nil
	case "units":
		playerImpl.unitsImpl = necrowarDeltaMerge.ArrayOfUnit(&playerImpl.unitsImpl, delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

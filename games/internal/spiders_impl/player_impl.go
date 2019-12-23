package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Spiders.
type PlayerImpl struct {
	GameObjectImpl

	broodMotherImpl             spiders.BroodMother
	clientTypeImpl              string
	lostImpl                    bool
	maxSpiderlingsImpl          int64
	nameImpl                    string
	numberOfNestsControlledImpl int64
	opponentImpl                spiders.Player
	reasonLostImpl              string
	reasonWonImpl               string
	spidersImpl                 []spiders.Spider
	timeRemainingImpl           float64
	wonImpl                     bool
}

// BroodMother returns this player's BroodMother. If it dies they lose the
// game.
func (playerImpl *PlayerImpl) BroodMother() spiders.BroodMother {
	return playerImpl.broodMotherImpl
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

// MaxSpiderlings returns the max number of Spiderlings players can spawn.
func (playerImpl *PlayerImpl) MaxSpiderlings() int64 {
	return playerImpl.maxSpiderlingsImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// NumberOfNestsControlled returns the number of nests this player
// controls.
func (playerImpl *PlayerImpl) NumberOfNestsControlled() int64 {
	return playerImpl.numberOfNestsControlledImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() spiders.Player {
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

// Spiders returns all the Spiders owned by this player.
func (playerImpl *PlayerImpl) Spiders() []spiders.Spider {
	return playerImpl.spidersImpl
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

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.broodMotherImpl = nil
	playerImpl.clientTypeImpl = ""
	playerImpl.lostImpl = true
	playerImpl.maxSpiderlingsImpl = 0
	playerImpl.nameImpl = ""
	playerImpl.numberOfNestsControlledImpl = 0
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.spidersImpl = []spiders.Spider{}
	playerImpl.timeRemainingImpl = 0
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

	spidersDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'spiders.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "broodMother":
		playerImpl.broodMotherImpl = spidersDeltaMerge.BroodMother(delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = spidersDeltaMerge.Boolean(delta)
		return true, nil
	case "maxSpiderlings":
		playerImpl.maxSpiderlingsImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "numberOfNestsControlled":
		playerImpl.numberOfNestsControlledImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = spidersDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "spiders":
		playerImpl.spidersImpl = spidersDeltaMerge.ArrayOfSpider(&playerImpl.spidersImpl, delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = spidersDeltaMerge.Float(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = spidersDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

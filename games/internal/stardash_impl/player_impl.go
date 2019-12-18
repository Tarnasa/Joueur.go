package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Stardash.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	homeBaseImpl      stardash.Body
	lostImpl          bool
	moneyImpl         int64
	nameImpl          string
	opponentImpl      stardash.Player
	projectilesImpl   []stardash.Projectile
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	unitsImpl         []stardash.Unit
	victoryPointsImpl int64
	wonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// HomeBase returns the home base of the player.
func (playerImpl *PlayerImpl) HomeBase() stardash.Body {
	return playerImpl.homeBaseImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Money returns the amount of money this Player has.
func (playerImpl *PlayerImpl) Money() int64 {
	return playerImpl.moneyImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() stardash.Player {
	return playerImpl.opponentImpl
}

// Projectiles returns every Projectile owned by this Player. The earlier
// in the array the older they are.
func (playerImpl *PlayerImpl) Projectiles() []stardash.Projectile {
	return playerImpl.projectilesImpl
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

// Units returns every Unit owned by this Player. The earlier in the array
// the older they are.
func (playerImpl *PlayerImpl) Units() []stardash.Unit {
	return playerImpl.unitsImpl
}

// VictoryPoints returns the number of victory points the player has.
func (playerImpl *PlayerImpl) VictoryPoints() int64 {
	return playerImpl.victoryPointsImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.homeBaseImpl = nil
	playerImpl.lostImpl = true
	playerImpl.moneyImpl = 0
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.projectilesImpl = []stardash.Projectile{}
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.unitsImpl = []stardash.Unit{}
	playerImpl.victoryPointsImpl = 0
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

	stardashDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'stardash.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "homeBase":
		playerImpl.homeBaseImpl = stardashDeltaMerge.Body(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = stardashDeltaMerge.Boolean(delta)
		return true, nil
	case "money":
		playerImpl.moneyImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = stardashDeltaMerge.Player(delta)
		return true, nil
	case "projectiles":
		playerImpl.projectilesImpl = stardashDeltaMerge.ArrayOfProjectile(&playerImpl.projectilesImpl, delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "units":
		playerImpl.unitsImpl = stardashDeltaMerge.ArrayOfUnit(&playerImpl.unitsImpl, delta)
		return true, nil
	case "victoryPoints":
		playerImpl.victoryPointsImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = stardashDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

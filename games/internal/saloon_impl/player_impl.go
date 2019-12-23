package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Saloon.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl    string
	cowboysImpl       []saloon.Cowboy
	killsImpl         int64
	lostImpl          bool
	nameImpl          string
	opponentImpl      saloon.Player
	reasonLostImpl    string
	reasonWonImpl     string
	rowdinessImpl     int64
	scoreImpl         int64
	siestaImpl        int64
	timeRemainingImpl float64
	wonImpl           bool
	youngGunImpl      saloon.YoungGun
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// Cowboys returns every Cowboy owned by this Player.
func (playerImpl *PlayerImpl) Cowboys() []saloon.Cowboy {
	return playerImpl.cowboysImpl
}

// Kills returns how many enemy Cowboys this player's team has killed.
func (playerImpl *PlayerImpl) Kills() int64 {
	return playerImpl.killsImpl
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
func (playerImpl *PlayerImpl) Opponent() saloon.Player {
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

// Rowdiness returns how rowdy their team is. When it gets too high their
// team takes a collective siesta.
func (playerImpl *PlayerImpl) Rowdiness() int64 {
	return playerImpl.rowdinessImpl
}

// Score returns how many times their team has played a piano.
func (playerImpl *PlayerImpl) Score() int64 {
	return playerImpl.scoreImpl
}

// Siesta returns 0 when not having a team siesta. When greater than 0
// represents how many turns left for the team siesta to complete.
func (playerImpl *PlayerImpl) Siesta() int64 {
	return playerImpl.siestaImpl
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

// YoungGun returns the YoungGun this Player uses to call in new Cowboys.
func (playerImpl *PlayerImpl) YoungGun() saloon.YoungGun {
	return playerImpl.youngGunImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.cowboysImpl = []saloon.Cowboy{}
	playerImpl.killsImpl = 0
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.rowdinessImpl = 0
	playerImpl.scoreImpl = 0
	playerImpl.siestaImpl = 0
	playerImpl.timeRemainingImpl = 0
	playerImpl.wonImpl = true
	playerImpl.youngGunImpl = nil
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

	saloonDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "cowboys":
		playerImpl.cowboysImpl = saloonDeltaMerge.ArrayOfCowboy(&playerImpl.cowboysImpl, delta)
		return true, nil
	case "kills":
		playerImpl.killsImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = saloonDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "rowdiness":
		playerImpl.rowdinessImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "score":
		playerImpl.scoreImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "siesta":
		playerImpl.siestaImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = saloonDeltaMerge.Float(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "youngGun":
		playerImpl.youngGunImpl = saloonDeltaMerge.YoungGun(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

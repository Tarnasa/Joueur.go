package impl

import (
	"joueur/base"
	"joueur/games/chess"
)

// PlayerImpl is the struct that implements the container for Player instances in Chess.
type PlayerImpl struct {
	GameObjectImpl
	implclientType    string
	implcolor         string
	impllost          bool
	implname          string
	implopponent      chess.Player
	implreasonLost    string
	implreasonWon     string
	impltimeRemaining float64
	implwon           bool
}

// ClientType returns what type of client this is, e.g. 'Python', 'JavaScript', or some other language. For potential data mining purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.ClientTypeImpl
}

// Color returns the color (side) of this player. Either 'white' or 'black', with the 'white' player having the first move.
func (playerImpl *PlayerImpl) Color() string {
	return playerImpl.ColorImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.LostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.NameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() chess.Player {
	return playerImpl.OpponentImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.ReasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.ReasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.TimeRemainingImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.WonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.implclientType = ""
	playerImpl.implcolor = ""
	playerImpl.impllost = false
	playerImpl.implname = ""
	playerImpl.implopponent = nil
	playerImpl.implreasonLost = ""
	playerImpl.implreasonWon = ""
	playerImpl.impltimeRemaining = 0
	playerImpl.implwon = false
}

// DeltaMerge merged the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) {
	switch(attribute) {
	case "clientType":
		(*playerImpl).clientTypeImpl = deltaMerge.String(delta)
		break
	case "color":
		(*playerImpl).colorImpl = deltaMerge.String(delta)
		break
	case "lost":
		(*playerImpl).lostImpl = deltaMerge.Boolean(delta)
		break
	case "name":
		(*playerImpl).nameImpl = deltaMerge.String(delta)
		break
	case "opponent":
		(*playerImpl).opponentImpl = deltaMerge.Player(delta)
		break
	case "reasonLost":
		(*playerImpl).reasonLostImpl = deltaMerge.String(delta)
		break
	case "reasonWon":
		(*playerImpl).reasonWonImpl = deltaMerge.String(delta)
		break
	case "timeRemaining":
		(*playerImpl).timeRemainingImpl = deltaMerge.Float(delta)
		break
	case "won":
		(*playerImpl).wonImpl = deltaMerge.Boolean(delta)
		break
	}
}

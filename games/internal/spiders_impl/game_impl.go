package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// GameImpl is the struct that implements the container for Game instances
// in Spiders.
type GameImpl struct {
	base.GameImpl

	currentPlayerImpl      spiders.Player
	currentTurnImpl        int64
	cutSpeedImpl           int64
	eggsScalarImpl         float64
	gameObjectsImpl        map[string]spiders.GameObject
	initialWebStrengthImpl int64
	maxTurnsImpl           int64
	maxWebStrengthImpl     int64
	movementSpeedImpl      int64
	nestsImpl              []spiders.Nest
	playersImpl            []spiders.Player
	sessionImpl            string
	spitSpeedImpl          int64
	timeAddedPerTurnImpl   int64
	weavePowerImpl         int64
	weaveSpeedImpl         int64
	websImpl               []spiders.Web
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() spiders.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// CutSpeed returns the speed at which Cutters work to do cut Webs.
func (gameImpl *GameImpl) CutSpeed() int64 {
	return gameImpl.cutSpeedImpl
}

// EggsScalar returns constant used to calculate how many eggs BroodMothers
// get on their owner's turns.
func (gameImpl *GameImpl) EggsScalar() float64 {
	return gameImpl.eggsScalarImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]spiders.GameObject {
	return gameImpl.gameObjectsImpl
}

// InitialWebStrength returns the starting strength for Webs.
func (gameImpl *GameImpl) InitialWebStrength() int64 {
	return gameImpl.initialWebStrengthImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// MaxWebStrength returns the maximum strength a web can be strengthened
// to.
func (gameImpl *GameImpl) MaxWebStrength() int64 {
	return gameImpl.maxWebStrengthImpl
}

// MovementSpeed returns the speed at which Spiderlings move on Webs.
func (gameImpl *GameImpl) MovementSpeed() int64 {
	return gameImpl.movementSpeedImpl
}

// Nests returns every Nest in the game.
func (gameImpl *GameImpl) Nests() []spiders.Nest {
	return gameImpl.nestsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []spiders.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// SpitSpeed returns the speed at which Spitters work to spit new Webs.
func (gameImpl *GameImpl) SpitSpeed() int64 {
	return gameImpl.spitSpeedImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// WeavePower returns how much web strength is added or removed from Webs
// when they are weaved.
func (gameImpl *GameImpl) WeavePower() int64 {
	return gameImpl.weavePowerImpl
}

// WeaveSpeed returns the speed at which Weavers work to do strengthens and
// weakens on Webs.
func (gameImpl *GameImpl) WeaveSpeed() int64 {
	return gameImpl.weaveSpeedImpl
}

// Webs returns every Web in the game.
func (gameImpl *GameImpl) Webs() []spiders.Web {
	return gameImpl.websImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.cutSpeedImpl = 0
	gameImpl.eggsScalarImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]spiders.GameObject)
	gameImpl.initialWebStrengthImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.maxWebStrengthImpl = 0
	gameImpl.movementSpeedImpl = 0
	gameImpl.nestsImpl = []spiders.Nest{}
	gameImpl.playersImpl = []spiders.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.spitSpeedImpl = 0
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.weavePowerImpl = 0
	gameImpl.weaveSpeedImpl = 0
	gameImpl.websImpl = []spiders.Web{}
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

	spidersDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'spiders.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "currentPlayer":
		gameImpl.currentPlayerImpl = spidersDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "cutSpeed":
		gameImpl.cutSpeedImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "eggsScalar":
		gameImpl.eggsScalarImpl = spidersDeltaMerge.Float(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = spidersDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "initialWebStrength":
		gameImpl.initialWebStrengthImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "maxWebStrength":
		gameImpl.maxWebStrengthImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "movementSpeed":
		gameImpl.movementSpeedImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "nests":
		gameImpl.nestsImpl = spidersDeltaMerge.ArrayOfNest(&gameImpl.nestsImpl, delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = spidersDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "spitSpeed":
		gameImpl.spitSpeedImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "weavePower":
		gameImpl.weavePowerImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "weaveSpeed":
		gameImpl.weaveSpeedImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "webs":
		gameImpl.websImpl = spidersDeltaMerge.ArrayOfWeb(&gameImpl.websImpl, delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

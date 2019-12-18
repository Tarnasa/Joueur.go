package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// GameImpl is the struct that implements the container for Game instances
// in Pirates.
type GameImpl struct {
	base.GameImpl

	buryInterestRateImpl     float64
	crewCostImpl             int64
	crewDamageImpl           int64
	crewHealthImpl           int64
	crewMovesImpl            int64
	crewRangeImpl            float64
	currentPlayerImpl        pirates.Player
	currentTurnImpl          int64
	gameObjectsImpl          map[string]pirates.GameObject
	healFactorImpl           float64
	mapHeightImpl            int64
	mapWidthImpl             int64
	maxTurnsImpl             int64
	merchantGoldRateImpl     float64
	merchantInterestRateImpl float64
	minInterestDistanceImpl  float64
	playersImpl              []pirates.Player
	portsImpl                []pirates.Port
	restRangeImpl            float64
	sessionImpl              string
	shipCostImpl             int64
	shipDamageImpl           int64
	shipHealthImpl           int64
	shipMovesImpl            int64
	shipRangeImpl            float64
	tilesImpl                []pirates.Tile
	timeAddedPerTurnImpl     int64
	unitsImpl                []pirates.Unit
}

// BuryInterestRate returns the rate buried gold increases each turn.
func (gameImpl *GameImpl) BuryInterestRate() float64 {
	return gameImpl.buryInterestRateImpl
}

// CrewCost returns how much gold it costs to construct a single crew.
func (gameImpl *GameImpl) CrewCost() int64 {
	return gameImpl.crewCostImpl
}

// CrewDamage returns how much damage crew deal to each other.
func (gameImpl *GameImpl) CrewDamage() int64 {
	return gameImpl.crewDamageImpl
}

// CrewHealth returns the maximum amount of health a crew member can have.
func (gameImpl *GameImpl) CrewHealth() int64 {
	return gameImpl.crewHealthImpl
}

// CrewMoves returns the number of moves Units with only crew are given
// each turn.
func (gameImpl *GameImpl) CrewMoves() int64 {
	return gameImpl.crewMovesImpl
}

// CrewRange returns a crew's attack range. Range is circular.
func (gameImpl *GameImpl) CrewRange() float64 {
	return gameImpl.crewRangeImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() pirates.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]pirates.GameObject {
	return gameImpl.gameObjectsImpl
}

// HealFactor returns how much health a Unit recovers when they rest.
func (gameImpl *GameImpl) HealFactor() float64 {
	return gameImpl.healFactorImpl
}

// MapHeight returns the number of Tiles in the map along the y (vertical)
// axis.
func (gameImpl *GameImpl) MapHeight() int64 {
	return gameImpl.mapHeightImpl
}

// MapWidth returns the number of Tiles in the map along the x (horizontal)
// axis.
func (gameImpl *GameImpl) MapWidth() int64 {
	return gameImpl.mapWidthImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// MerchantGoldRate returns how much gold merchant Ports get each turn.
func (gameImpl *GameImpl) MerchantGoldRate() float64 {
	return gameImpl.merchantGoldRateImpl
}

// MerchantInterestRate returns when a merchant ship spawns, the amount of
// additional gold it has relative to the Port's investment.
func (gameImpl *GameImpl) MerchantInterestRate() float64 {
	return gameImpl.merchantInterestRateImpl
}

// MinInterestDistance returns the Euclidean distance buried gold must be
// from the Player's Port to accumulate interest.
func (gameImpl *GameImpl) MinInterestDistance() float64 {
	return gameImpl.minInterestDistanceImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []pirates.Player {
	return gameImpl.playersImpl
}

// Ports returns every Port in the game. Merchant ports have owner set to
// nil.
func (gameImpl *GameImpl) Ports() []pirates.Port {
	return gameImpl.portsImpl
}

// RestRange returns how far a Unit can be from a Port to rest. Range is
// circular.
func (gameImpl *GameImpl) RestRange() float64 {
	return gameImpl.restRangeImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// ShipCost returns how much gold it costs to construct a ship.
func (gameImpl *GameImpl) ShipCost() int64 {
	return gameImpl.shipCostImpl
}

// ShipDamage returns how much damage ships deal to ships and ports.
func (gameImpl *GameImpl) ShipDamage() int64 {
	return gameImpl.shipDamageImpl
}

// ShipHealth returns the maximum amount of health a ship can have.
func (gameImpl *GameImpl) ShipHealth() int64 {
	return gameImpl.shipHealthImpl
}

// ShipMoves returns the number of moves Units with ships are given each
// turn.
func (gameImpl *GameImpl) ShipMoves() int64 {
	return gameImpl.shipMovesImpl
}

// ShipRange returns a ship's attack range. Range is circular.
func (gameImpl *GameImpl) ShipRange() float64 {
	return gameImpl.shipRangeImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []pirates.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// Units returns every Unit in the game. Merchant units have targetPort set
// to a port.
func (gameImpl *GameImpl) Units() []pirates.Unit {
	return gameImpl.unitsImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.buryInterestRateImpl = 0
	gameImpl.crewCostImpl = 0
	gameImpl.crewDamageImpl = 0
	gameImpl.crewHealthImpl = 0
	gameImpl.crewMovesImpl = 0
	gameImpl.crewRangeImpl = 0
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]pirates.GameObject)
	gameImpl.healFactorImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.merchantGoldRateImpl = 0
	gameImpl.merchantInterestRateImpl = 0
	gameImpl.minInterestDistanceImpl = 0
	gameImpl.playersImpl = []pirates.Player{}
	gameImpl.portsImpl = []pirates.Port{}
	gameImpl.restRangeImpl = 0
	gameImpl.sessionImpl = ""
	gameImpl.shipCostImpl = 0
	gameImpl.shipDamageImpl = 0
	gameImpl.shipHealthImpl = 0
	gameImpl.shipMovesImpl = 0
	gameImpl.shipRangeImpl = 0
	gameImpl.tilesImpl = []pirates.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.unitsImpl = []pirates.Unit{}
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

	piratesDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'pirates.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "buryInterestRate":
		gameImpl.buryInterestRateImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "crewCost":
		gameImpl.crewCostImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "crewDamage":
		gameImpl.crewDamageImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "crewHealth":
		gameImpl.crewHealthImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "crewMoves":
		gameImpl.crewMovesImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "crewRange":
		gameImpl.crewRangeImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = piratesDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = piratesDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "healFactor":
		gameImpl.healFactorImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "merchantGoldRate":
		gameImpl.merchantGoldRateImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "merchantInterestRate":
		gameImpl.merchantInterestRateImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "minInterestDistance":
		gameImpl.minInterestDistanceImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = piratesDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "ports":
		gameImpl.portsImpl = piratesDeltaMerge.ArrayOfPort(&gameImpl.portsImpl, delta)
		return true, nil
	case "restRange":
		gameImpl.restRangeImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = piratesDeltaMerge.String(delta)
		return true, nil
	case "shipCost":
		gameImpl.shipCostImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "shipDamage":
		gameImpl.shipDamageImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "shipHealth":
		gameImpl.shipHealthImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "shipMoves":
		gameImpl.shipMovesImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "shipRange":
		gameImpl.shipRangeImpl = piratesDeltaMerge.Float(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = piratesDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "units":
		gameImpl.unitsImpl = piratesDeltaMerge.ArrayOfUnit(&gameImpl.unitsImpl, delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) pirates.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x + y * gameImpl.mapWidthImpl]
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// GameImpl is the struct that implements the container for Game instances
// in Necrowar.
type GameImpl struct {
	base.GameImpl

	currentPlayerImpl       necrowar.Player
	currentTurnImpl         int64
	gameObjectsImpl         map[string]necrowar.GameObject
	goldIncomePerUnitImpl   int64
	islandIncomePerUnitImpl int64
	manaIncomePerUnitImpl   int64
	mapHeightImpl           int64
	mapWidthImpl            int64
	maxTurnsImpl            int64
	playersImpl             []necrowar.Player
	riverPhaseImpl          int64
	sessionImpl             string
	tilesImpl               []necrowar.Tile
	timeAddedPerTurnImpl    int64
	towerJobsImpl           []necrowar.TowerJob
	towersImpl              []necrowar.Tower
	unitJobsImpl            []necrowar.UnitJob
	unitsImpl               []necrowar.Unit
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() necrowar.Player {
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
func (gameImpl *GameImpl) GameObjects() map[string]necrowar.GameObject {
	return gameImpl.gameObjectsImpl
}

// GoldIncomePerUnit returns the amount of gold income per turn per unit in
// a mine.
func (gameImpl *GameImpl) GoldIncomePerUnit() int64 {
	return gameImpl.goldIncomePerUnitImpl
}

// IslandIncomePerUnit returns the amount of gold income per turn per unit
// in the island mine.
func (gameImpl *GameImpl) IslandIncomePerUnit() int64 {
	return gameImpl.islandIncomePerUnitImpl
}

// ManaIncomePerUnit returns the Amount of gold income per turn per unit
// fishing on the river side.
func (gameImpl *GameImpl) ManaIncomePerUnit() int64 {
	return gameImpl.manaIncomePerUnitImpl
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

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []necrowar.Player {
	return gameImpl.playersImpl
}

// RiverPhase returns the amount of turns it takes between the river
// changing phases.
func (gameImpl *GameImpl) RiverPhase() int64 {
	return gameImpl.riverPhaseImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []necrowar.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// TowerJobs returns an array of every tower type / job.
func (gameImpl *GameImpl) TowerJobs() []necrowar.TowerJob {
	return gameImpl.towerJobsImpl
}

// Towers returns every Tower in the game.
func (gameImpl *GameImpl) Towers() []necrowar.Tower {
	return gameImpl.towersImpl
}

// UnitJobs returns an array of every unit type / job.
func (gameImpl *GameImpl) UnitJobs() []necrowar.UnitJob {
	return gameImpl.unitJobsImpl
}

// Units returns every Unit in the game.
func (gameImpl *GameImpl) Units() []necrowar.Unit {
	return gameImpl.unitsImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]necrowar.GameObject)
	gameImpl.goldIncomePerUnitImpl = 0
	gameImpl.islandIncomePerUnitImpl = 0
	gameImpl.manaIncomePerUnitImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.playersImpl = []necrowar.Player{}
	gameImpl.riverPhaseImpl = 0
	gameImpl.sessionImpl = ""
	gameImpl.tilesImpl = []necrowar.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.towerJobsImpl = []necrowar.TowerJob{}
	gameImpl.towersImpl = []necrowar.Tower{}
	gameImpl.unitJobsImpl = []necrowar.UnitJob{}
	gameImpl.unitsImpl = []necrowar.Unit{}
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

	necrowarDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'necrowar.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "currentPlayer":
		gameImpl.currentPlayerImpl = necrowarDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = necrowarDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "goldIncomePerUnit":
		gameImpl.goldIncomePerUnitImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "islandIncomePerUnit":
		gameImpl.islandIncomePerUnitImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "manaIncomePerUnit":
		gameImpl.manaIncomePerUnitImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = necrowarDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "riverPhase":
		gameImpl.riverPhaseImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = necrowarDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "towerJobs":
		gameImpl.towerJobsImpl = necrowarDeltaMerge.ArrayOfTowerJob(&gameImpl.towerJobsImpl, delta)
		return true, nil
	case "towers":
		gameImpl.towersImpl = necrowarDeltaMerge.ArrayOfTower(&gameImpl.towersImpl, delta)
		return true, nil
	case "unitJobs":
		gameImpl.unitJobsImpl = necrowarDeltaMerge.ArrayOfUnitJob(&gameImpl.unitJobsImpl, delta)
		return true, nil
	case "units":
		gameImpl.unitsImpl = necrowarDeltaMerge.ArrayOfUnit(&gameImpl.unitsImpl, delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) necrowar.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x + y * gameImpl.mapWidthImpl]
}

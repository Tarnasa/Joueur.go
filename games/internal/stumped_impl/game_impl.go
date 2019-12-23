package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// GameImpl is the struct that implements the container for Game instances
// in Stumped.
type GameImpl struct {
	base.GameImpl

	beaversImpl                []stumped.Beaver
	currentPlayerImpl          stumped.Player
	currentTurnImpl            int64
	freeBeaversCountImpl       int64
	gameObjectsImpl            map[string]stumped.GameObject
	jobsImpl                   []stumped.Job
	lodgeCostConstantImpl      float64
	lodgesToWinImpl            int64
	mapHeightImpl              int64
	mapWidthImpl               int64
	maxTurnsImpl               int64
	playersImpl                []stumped.Player
	sessionImpl                string
	spawnerImpl                []stumped.Spawner
	spawnerHarvestConstantImpl float64
	spawnerTypesImpl           []string
	tilesImpl                  []stumped.Tile
	timeAddedPerTurnImpl       int64
}

// Beavers returns every Beaver in the game.
func (gameImpl *GameImpl) Beavers() []stumped.Beaver {
	return gameImpl.beaversImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() stumped.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// FreeBeaversCount returns when a Player has less Beavers than this
// number, then recruiting other Beavers is free.
func (gameImpl *GameImpl) FreeBeaversCount() int64 {
	return gameImpl.freeBeaversCountImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]stumped.GameObject {
	return gameImpl.gameObjectsImpl
}

// Jobs returns all the Jobs that Beavers can have in the game.
func (gameImpl *GameImpl) Jobs() []stumped.Job {
	return gameImpl.jobsImpl
}

// LodgeCostConstant returns constant number used to calculate what it
// costs to spawn a new lodge.
func (gameImpl *GameImpl) LodgeCostConstant() float64 {
	return gameImpl.lodgeCostConstantImpl
}

// LodgesToWin returns how many lodges must be owned by a Player at once to
// win the game.
func (gameImpl *GameImpl) LodgesToWin() int64 {
	return gameImpl.lodgesToWinImpl
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
func (gameImpl *GameImpl) Players() []stumped.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// Spawner returns every Spawner in the game.
func (gameImpl *GameImpl) Spawner() []stumped.Spawner {
	return gameImpl.spawnerImpl
}

// SpawnerHarvestConstant returns constant number used to calculate how
// many branches/food Beavers harvest from Spawners.
func (gameImpl *GameImpl) SpawnerHarvestConstant() float64 {
	return gameImpl.spawnerHarvestConstantImpl
}

// SpawnerTypes returns all the types of Spawners in the game.
func (gameImpl *GameImpl) SpawnerTypes() []string {
	return gameImpl.spawnerTypesImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []stumped.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.beaversImpl = []stumped.Beaver{}
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.freeBeaversCountImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]stumped.GameObject)
	gameImpl.jobsImpl = []stumped.Job{}
	gameImpl.lodgeCostConstantImpl = 0
	gameImpl.lodgesToWinImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.playersImpl = []stumped.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.spawnerImpl = []stumped.Spawner{}
	gameImpl.spawnerHarvestConstantImpl = 0
	gameImpl.spawnerTypesImpl = []string{}
	gameImpl.tilesImpl = []stumped.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
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

	stumpedDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stumped.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "beavers":
		gameImpl.beaversImpl = stumpedDeltaMerge.ArrayOfBeaver(&gameImpl.beaversImpl, delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = stumpedDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "freeBeaversCount":
		gameImpl.freeBeaversCountImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = stumpedDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "jobs":
		gameImpl.jobsImpl = stumpedDeltaMerge.ArrayOfJob(&gameImpl.jobsImpl, delta)
		return true, nil
	case "lodgeCostConstant":
		gameImpl.lodgeCostConstantImpl = stumpedDeltaMerge.Float(delta)
		return true, nil
	case "lodgesToWin":
		gameImpl.lodgesToWinImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = stumpedDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	case "spawner":
		gameImpl.spawnerImpl = stumpedDeltaMerge.ArrayOfSpawner(&gameImpl.spawnerImpl, delta)
		return true, nil
	case "spawnerHarvestConstant":
		gameImpl.spawnerHarvestConstantImpl = stumpedDeltaMerge.Float(delta)
		return true, nil
	case "spawnerTypes":
		gameImpl.spawnerTypesImpl = stumpedDeltaMerge.ArrayOfString(&gameImpl.spawnerTypesImpl, delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = stumpedDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) stumped.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x + y * gameImpl.mapWidthImpl]
}

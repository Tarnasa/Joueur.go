package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// GameImpl is the struct that implements the container for Game instances
// in Newtonian.
type GameImpl struct {
	base.GameImpl

	currentPlayerImpl    newtonian.Player
	currentTurnImpl      int64
	gameObjectsImpl      map[string]newtonian.GameObject
	internCapImpl        int64
	jobsImpl             []newtonian.Job
	machinesImpl         []newtonian.Machine
	managerCapImpl       int64
	mapHeightImpl        int64
	mapWidthImpl         int64
	materialSpawnImpl    int64
	maxTurnsImpl         int64
	physicistCapImpl     int64
	playersImpl          []newtonian.Player
	refinedValueImpl     int64
	regenerateRateImpl   float64
	sessionImpl          string
	spawnTimeImpl        int64
	stunTimeImpl         int64
	tilesImpl            []newtonian.Tile
	timeAddedPerTurnImpl int64
	timeImmuneImpl       int64
	unitsImpl            []newtonian.Unit
	victoryAmountImpl    int64
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() newtonian.Player {
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
func (gameImpl *GameImpl) GameObjects() map[string]newtonian.GameObject {
	return gameImpl.gameObjectsImpl
}

// InternCap returns the maximum number of interns a player can have.
func (gameImpl *GameImpl) InternCap() int64 {
	return gameImpl.internCapImpl
}

// Jobs returns an array of all jobs. The first element is intern, second
// is physicists, and third is manager.
func (gameImpl *GameImpl) Jobs() []newtonian.Job {
	return gameImpl.jobsImpl
}

// Machines returns every Machine in the game.
func (gameImpl *GameImpl) Machines() []newtonian.Machine {
	return gameImpl.machinesImpl
}

// ManagerCap returns the maximum number of managers a player can have.
func (gameImpl *GameImpl) ManagerCap() int64 {
	return gameImpl.managerCapImpl
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

// MaterialSpawn returns the number of materials that spawn per spawn
// cycle.
func (gameImpl *GameImpl) MaterialSpawn() int64 {
	return gameImpl.materialSpawnImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// PhysicistCap returns the maximum number of physicists a player can have.
func (gameImpl *GameImpl) PhysicistCap() int64 {
	return gameImpl.physicistCapImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []newtonian.Player {
	return gameImpl.playersImpl
}

// RefinedValue returns the amount of victory points added when a refined
// ore is consumed by the generator.
func (gameImpl *GameImpl) RefinedValue() int64 {
	return gameImpl.refinedValueImpl
}

// RegenerateRate returns the percent of max HP regained when a unit end
// their turn on a tile owned by their player.
func (gameImpl *GameImpl) RegenerateRate() float64 {
	return gameImpl.regenerateRateImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// SpawnTime returns the amount of turns it takes a unit to spawn.
func (gameImpl *GameImpl) SpawnTime() int64 {
	return gameImpl.spawnTimeImpl
}

// StunTime returns the amount of turns a unit cannot do anything when
// stunned.
func (gameImpl *GameImpl) StunTime() int64 {
	return gameImpl.stunTimeImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []newtonian.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// TimeImmune returns the number turns a unit is immune to being stunned.
func (gameImpl *GameImpl) TimeImmune() int64 {
	return gameImpl.timeImmuneImpl
}

// Units returns every Unit in the game.
func (gameImpl *GameImpl) Units() []newtonian.Unit {
	return gameImpl.unitsImpl
}

// VictoryAmount returns the amount of combined heat and pressure that you
// need to win.
func (gameImpl *GameImpl) VictoryAmount() int64 {
	return gameImpl.victoryAmountImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]newtonian.GameObject)
	gameImpl.internCapImpl = 0
	gameImpl.jobsImpl = []newtonian.Job{}
	gameImpl.machinesImpl = []newtonian.Machine{}
	gameImpl.managerCapImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.materialSpawnImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.physicistCapImpl = 0
	gameImpl.playersImpl = []newtonian.Player{}
	gameImpl.refinedValueImpl = 0
	gameImpl.regenerateRateImpl = 0
	gameImpl.sessionImpl = ""
	gameImpl.spawnTimeImpl = 0
	gameImpl.stunTimeImpl = 0
	gameImpl.tilesImpl = []newtonian.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.timeImmuneImpl = 0
	gameImpl.unitsImpl = []newtonian.Unit{}
	gameImpl.victoryAmountImpl = 0
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

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "currentPlayer":
		gameImpl.currentPlayerImpl = newtonianDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = newtonianDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "internCap":
		gameImpl.internCapImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "jobs":
		gameImpl.jobsImpl = newtonianDeltaMerge.ArrayOfJob(&gameImpl.jobsImpl, delta)
		return true, nil
	case "machines":
		gameImpl.machinesImpl = newtonianDeltaMerge.ArrayOfMachine(&gameImpl.machinesImpl, delta)
		return true, nil
	case "managerCap":
		gameImpl.managerCapImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "materialSpawn":
		gameImpl.materialSpawnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "physicistCap":
		gameImpl.physicistCapImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = newtonianDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "refinedValue":
		gameImpl.refinedValueImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "regenerateRate":
		gameImpl.regenerateRateImpl = newtonianDeltaMerge.Float(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "spawnTime":
		gameImpl.spawnTimeImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "stunTime":
		gameImpl.stunTimeImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = newtonianDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "timeImmune":
		gameImpl.timeImmuneImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "units":
		gameImpl.unitsImpl = newtonianDeltaMerge.ArrayOfUnit(&gameImpl.unitsImpl, delta)
		return true, nil
	case "victoryAmount":
		gameImpl.victoryAmountImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) newtonian.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x+y*gameImpl.mapWidthImpl]
}

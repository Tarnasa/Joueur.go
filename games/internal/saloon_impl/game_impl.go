package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// GameImpl is the struct that implements the container for Game instances
// in Saloon.
type GameImpl struct {
	base.GameImpl

	bartenderCooldownImpl  int64
	bottlesImpl            []saloon.Bottle
	brawlerDamageImpl      int64
	cowboysImpl            []saloon.Cowboy
	currentPlayerImpl      saloon.Player
	currentTurnImpl        int64
	furnishingsImpl        []saloon.Furnishing
	gameObjectsImpl        map[string]saloon.GameObject
	jobsImpl               []string
	mapHeightImpl          int64
	mapWidthImpl           int64
	maxCowboysPerJobImpl   int64
	maxTurnsImpl           int64
	playersImpl            []saloon.Player
	rowdinessToSiestaImpl  int64
	sessionImpl            string
	sharpshooterDamageImpl int64
	siestaLengthImpl       int64
	tilesImpl              []saloon.Tile
	timeAddedPerTurnImpl   int64
	turnsDrunkImpl         int64
}

// BartenderCooldown returns how many turns a Bartender will be busy for
// after throwing a Bottle.
func (gameImpl *GameImpl) BartenderCooldown() int64 {
	return gameImpl.bartenderCooldownImpl
}

// Bottles returns all the beer Bottles currently flying across the saloon
// in the game.
func (gameImpl *GameImpl) Bottles() []saloon.Bottle {
	return gameImpl.bottlesImpl
}

// BrawlerDamage returns how much damage is applied to neighboring things
// bit by the Sharpshooter between turns.
func (gameImpl *GameImpl) BrawlerDamage() int64 {
	return gameImpl.brawlerDamageImpl
}

// Cowboys returns every Cowboy in the game.
func (gameImpl *GameImpl) Cowboys() []saloon.Cowboy {
	return gameImpl.cowboysImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() saloon.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// Furnishings returns every furnishing in the game.
func (gameImpl *GameImpl) Furnishings() []saloon.Furnishing {
	return gameImpl.furnishingsImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]saloon.GameObject {
	return gameImpl.gameObjectsImpl
}

// Jobs returns all the jobs that Cowboys can be called in with.
func (gameImpl *GameImpl) Jobs() []string {
	return gameImpl.jobsImpl
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

// MaxCowboysPerJob returns the maximum number of Cowboys a Player can
// bring into the saloon of each specific job.
func (gameImpl *GameImpl) MaxCowboysPerJob() int64 {
	return gameImpl.maxCowboysPerJobImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []saloon.Player {
	return gameImpl.playersImpl
}

// RowdinessToSiesta returns when a player's rowdiness reaches or exceeds
// this number their Cowboys take a collective siesta.
func (gameImpl *GameImpl) RowdinessToSiesta() int64 {
	return gameImpl.rowdinessToSiestaImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// SharpshooterDamage returns how much damage is applied to things hit by
// Sharpshooters when they act.
func (gameImpl *GameImpl) SharpshooterDamage() int64 {
	return gameImpl.sharpshooterDamageImpl
}

// SiestaLength returns how long siestas are for a player's team.
func (gameImpl *GameImpl) SiestaLength() int64 {
	return gameImpl.siestaLengthImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []saloon.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// TurnsDrunk returns how many turns a Cowboy will be drunk for if a bottle
// breaks on it.
func (gameImpl *GameImpl) TurnsDrunk() int64 {
	return gameImpl.turnsDrunkImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.bartenderCooldownImpl = 0
	gameImpl.bottlesImpl = []saloon.Bottle{}
	gameImpl.brawlerDamageImpl = 0
	gameImpl.cowboysImpl = []saloon.Cowboy{}
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.furnishingsImpl = []saloon.Furnishing{}
	gameImpl.gameObjectsImpl = make(map[string]saloon.GameObject)
	gameImpl.jobsImpl = []string{}
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxCowboysPerJobImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.playersImpl = []saloon.Player{}
	gameImpl.rowdinessToSiestaImpl = 0
	gameImpl.sessionImpl = ""
	gameImpl.sharpshooterDamageImpl = 0
	gameImpl.siestaLengthImpl = 0
	gameImpl.tilesImpl = []saloon.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.turnsDrunkImpl = 0
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

	saloonDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bartenderCooldown":
		gameImpl.bartenderCooldownImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "bottles":
		gameImpl.bottlesImpl = saloonDeltaMerge.ArrayOfBottle(&gameImpl.bottlesImpl, delta)
		return true, nil
	case "brawlerDamage":
		gameImpl.brawlerDamageImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "cowboys":
		gameImpl.cowboysImpl = saloonDeltaMerge.ArrayOfCowboy(&gameImpl.cowboysImpl, delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = saloonDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "furnishings":
		gameImpl.furnishingsImpl = saloonDeltaMerge.ArrayOfFurnishing(&gameImpl.furnishingsImpl, delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = saloonDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "jobs":
		gameImpl.jobsImpl = saloonDeltaMerge.ArrayOfString(&gameImpl.jobsImpl, delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "maxCowboysPerJob":
		gameImpl.maxCowboysPerJobImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = saloonDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "rowdinessToSiesta":
		gameImpl.rowdinessToSiestaImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "sharpshooterDamage":
		gameImpl.sharpshooterDamageImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "siestaLength":
		gameImpl.siestaLengthImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = saloonDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "turnsDrunk":
		gameImpl.turnsDrunkImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) saloon.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x + y * gameImpl.mapWidthImpl]
}

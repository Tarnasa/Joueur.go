package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// GameImpl is the struct that implements the container for Game instances
// in Galapagos.
type GameImpl struct {
	base.GameImpl

	baseHealthImpl           int64
	creaturesImpl            []galapagos.Creature
	currentPlayerImpl        galapagos.Player
	currentTurnImpl          int64
	damageMultiplierImpl     int64
	gameObjectsImpl          map[string]galapagos.GameObject
	healthPerBreedImpl       int64
	healthPerCarnivorismImpl int64
	healthPerEnduranceImpl   int64
	healthPerHerbivorismImpl int64
	healthPerMoveImpl        int64
	healthPerTurnImpl        int64
	mapHeightImpl            int64
	mapWidthImpl             int64
	maxPlantSizeImpl         int64
	maxStartingCreaturesImpl int64
	maxStartingPlantsImpl    int64
	maxStatValueImpl         int64
	maxTurnsImpl             int64
	minStartingCreaturesImpl int64
	minStartingPlantsImpl    int64
	plantsImpl               []galapagos.Plant
	playersImpl              []galapagos.Player
	sessionImpl              string
	tilesImpl                []galapagos.Tile
	timeAddedPerTurnImpl     int64
}

// BaseHealth returns the amount of health that a creature with a 0
// endurance stat starts with.
func (gameImpl *GameImpl) BaseHealth() int64 {
	return gameImpl.baseHealthImpl
}

// Creatures returns every Creature in the game.
func (gameImpl *GameImpl) Creatures() []galapagos.Creature {
	return gameImpl.creaturesImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() galapagos.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// DamageMultiplier returns how much to damage an opponent per difference
// of carnivorism and defense.
func (gameImpl *GameImpl) DamageMultiplier() int64 {
	return gameImpl.damageMultiplierImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]galapagos.GameObject {
	return gameImpl.gameObjectsImpl
}

// HealthPerBreed returns the amount of extra health from both breeding
// creatures required if you have more total health than your opponent.
func (gameImpl *GameImpl) HealthPerBreed() int64 {
	return gameImpl.healthPerBreedImpl
}

// HealthPerCarnivorism returns multiplied by carnivorism to determine
// health gained from eating creatures.
func (gameImpl *GameImpl) HealthPerCarnivorism() int64 {
	return gameImpl.healthPerCarnivorismImpl
}

// HealthPerEndurance returns the amount of extra health for each point of
// endurance.
func (gameImpl *GameImpl) HealthPerEndurance() int64 {
	return gameImpl.healthPerEnduranceImpl
}

// HealthPerHerbivorism returns multiplied by herbivorism to determine
// health gained from biting plants.
func (gameImpl *GameImpl) HealthPerHerbivorism() int64 {
	return gameImpl.healthPerHerbivorismImpl
}

// HealthPerMove returns the amount of health required to move.
func (gameImpl *GameImpl) HealthPerMove() int64 {
	return gameImpl.healthPerMoveImpl
}

// HealthPerTurn returns the amount of health lost after each of your
// turns.
func (gameImpl *GameImpl) HealthPerTurn() int64 {
	return gameImpl.healthPerTurnImpl
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

// MaxPlantSize returns the maximum size a plant to grow to.
func (gameImpl *GameImpl) MaxPlantSize() int64 {
	return gameImpl.maxPlantSizeImpl
}

// MaxStartingCreatures returns the maximum number of creatures that each
// player will start with.
func (gameImpl *GameImpl) MaxStartingCreatures() int64 {
	return gameImpl.maxStartingCreaturesImpl
}

// MaxStartingPlants returns the maximum number of plants that the map will
// start with.
func (gameImpl *GameImpl) MaxStartingPlants() int64 {
	return gameImpl.maxStartingPlantsImpl
}

// MaxStatValue returns the maxmimum value that a stat (carnivorism,
// herbivorism, defense, endurance, speed) can have.
func (gameImpl *GameImpl) MaxStatValue() int64 {
	return gameImpl.maxStatValueImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// MinStartingCreatures returns the minimum number of creatures that each
// player will start with.
func (gameImpl *GameImpl) MinStartingCreatures() int64 {
	return gameImpl.minStartingCreaturesImpl
}

// MinStartingPlants returns the minimum number of plants that the map will
// start with.
func (gameImpl *GameImpl) MinStartingPlants() int64 {
	return gameImpl.minStartingPlantsImpl
}

// Plants returns every Plant in the game.
func (gameImpl *GameImpl) Plants() []galapagos.Plant {
	return gameImpl.plantsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []galapagos.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []galapagos.Tile {
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

	gameImpl.baseHealthImpl = 0
	gameImpl.creaturesImpl = []galapagos.Creature{}
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.damageMultiplierImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]galapagos.GameObject)
	gameImpl.healthPerBreedImpl = 0
	gameImpl.healthPerCarnivorismImpl = 0
	gameImpl.healthPerEnduranceImpl = 0
	gameImpl.healthPerHerbivorismImpl = 0
	gameImpl.healthPerMoveImpl = 0
	gameImpl.healthPerTurnImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxPlantSizeImpl = 0
	gameImpl.maxStartingCreaturesImpl = 0
	gameImpl.maxStartingPlantsImpl = 0
	gameImpl.maxStatValueImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.minStartingCreaturesImpl = 0
	gameImpl.minStartingPlantsImpl = 0
	gameImpl.plantsImpl = []galapagos.Plant{}
	gameImpl.playersImpl = []galapagos.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.tilesImpl = []galapagos.Tile{}
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

	galapagosDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'galapagos.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "baseHealth":
		gameImpl.baseHealthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "creatures":
		gameImpl.creaturesImpl = galapagosDeltaMerge.ArrayOfCreature(&gameImpl.creaturesImpl, delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = galapagosDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "damageMultiplier":
		gameImpl.damageMultiplierImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = galapagosDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "healthPerBreed":
		gameImpl.healthPerBreedImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "healthPerCarnivorism":
		gameImpl.healthPerCarnivorismImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "healthPerEndurance":
		gameImpl.healthPerEnduranceImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "healthPerHerbivorism":
		gameImpl.healthPerHerbivorismImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "healthPerMove":
		gameImpl.healthPerMoveImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "healthPerTurn":
		gameImpl.healthPerTurnImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "maxPlantSize":
		gameImpl.maxPlantSizeImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "maxStartingCreatures":
		gameImpl.maxStartingCreaturesImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "maxStartingPlants":
		gameImpl.maxStartingPlantsImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "maxStatValue":
		gameImpl.maxStatValueImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "minStartingCreatures":
		gameImpl.minStartingCreaturesImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "minStartingPlants":
		gameImpl.minStartingPlantsImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "plants":
		gameImpl.plantsImpl = galapagosDeltaMerge.ArrayOfPlant(&gameImpl.plantsImpl, delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = galapagosDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = galapagosDeltaMerge.String(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = galapagosDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) galapagos.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x+y*gameImpl.mapWidthImpl]
}

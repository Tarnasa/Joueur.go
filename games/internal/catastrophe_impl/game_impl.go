package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/catastrophe"
)

// GameImpl is the struct that implements the container for Game instances
// in Catastrophe.
type GameImpl struct {
	base.GameImpl

	catEnergyMultImpl        float64
	currentPlayerImpl        catastrophe.Player
	currentTurnImpl          int64
	gameObjectsImpl          map[string]catastrophe.GameObject
	harvestCooldownImpl      int64
	jobsImpl                 []catastrophe.Job
	lowerHarvestAmountImpl   int64
	mapHeightImpl            int64
	mapWidthImpl             int64
	maxTurnsImpl             int64
	monumentCostMultImpl     float64
	monumentMaterialsImpl    int64
	neutralMaterialsImpl     int64
	playersImpl              []catastrophe.Player
	sessionImpl              string
	shelterMaterialsImpl     int64
	startingFoodImpl         int64
	starvingEnergyMultImpl   float64
	structuresImpl           []catastrophe.Structure
	tilesImpl                []catastrophe.Tile
	timeAddedPerTurnImpl     int64
	turnsBetweenHarvestsImpl int64
	turnsToCreateHumanImpl   int64
	turnsToLowerHarvestImpl  int64
	unitsImpl                []catastrophe.Unit
	wallMaterialsImpl        int64
}

// CatEnergyMult returns the multiplier for the amount of energy
// regenerated when resting in a shelter with the cat overlord.
func (gameImpl *GameImpl) CatEnergyMult() float64 {
	return gameImpl.catEnergyMultImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() catastrophe.Player {
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
func (gameImpl *GameImpl) GameObjects() map[string]catastrophe.GameObject {
	return gameImpl.gameObjectsImpl
}

// HarvestCooldown returns the amount of turns it takes for a Tile that was
// just harvested to grow food again.
func (gameImpl *GameImpl) HarvestCooldown() int64 {
	return gameImpl.harvestCooldownImpl
}

// Jobs returns all the Jobs that Units can have in the game.
func (gameImpl *GameImpl) Jobs() []catastrophe.Job {
	return gameImpl.jobsImpl
}

// LowerHarvestAmount returns the amount that the harvest rate is lowered
// each season.
func (gameImpl *GameImpl) LowerHarvestAmount() int64 {
	return gameImpl.lowerHarvestAmountImpl
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

// MonumentCostMult returns the multiplier for the cost of actions when
// performing them in range of a monument. Does not effect pickup cost.
func (gameImpl *GameImpl) MonumentCostMult() float64 {
	return gameImpl.monumentCostMultImpl
}

// MonumentMaterials returns the number of materials in a monument.
func (gameImpl *GameImpl) MonumentMaterials() int64 {
	return gameImpl.monumentMaterialsImpl
}

// NeutralMaterials returns the number of materials in a neutral Structure.
func (gameImpl *GameImpl) NeutralMaterials() int64 {
	return gameImpl.neutralMaterialsImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []catastrophe.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// ShelterMaterials returns the number of materials in a shelter.
func (gameImpl *GameImpl) ShelterMaterials() int64 {
	return gameImpl.shelterMaterialsImpl
}

// StartingFood returns the amount of food Players start with.
func (gameImpl *GameImpl) StartingFood() int64 {
	return gameImpl.startingFoodImpl
}

// StarvingEnergyMult returns the multiplier for the amount of energy
// regenerated when resting while starving.
func (gameImpl *GameImpl) StarvingEnergyMult() float64 {
	return gameImpl.starvingEnergyMultImpl
}

// Structures returns every Structure in the game.
func (gameImpl *GameImpl) Structures() []catastrophe.Structure {
	return gameImpl.structuresImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []catastrophe.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() int64 {
	return gameImpl.timeAddedPerTurnImpl
}

// TurnsBetweenHarvests returns after a food tile is harvested, the number
// of turns before it can be harvested again.
func (gameImpl *GameImpl) TurnsBetweenHarvests() int64 {
	return gameImpl.turnsBetweenHarvestsImpl
}

// TurnsToCreateHuman returns the number of turns between fresh humans
// being spawned on the road.
func (gameImpl *GameImpl) TurnsToCreateHuman() int64 {
	return gameImpl.turnsToCreateHumanImpl
}

// TurnsToLowerHarvest returns the number of turns before the harvest rate
// is lowered (length of each season basically).
func (gameImpl *GameImpl) TurnsToLowerHarvest() int64 {
	return gameImpl.turnsToLowerHarvestImpl
}

// Units returns every Unit in the game.
func (gameImpl *GameImpl) Units() []catastrophe.Unit {
	return gameImpl.unitsImpl
}

// WallMaterials returns the number of materials in a wall.
func (gameImpl *GameImpl) WallMaterials() int64 {
	return gameImpl.wallMaterialsImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.catEnergyMultImpl = 0
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]catastrophe.GameObject)
	gameImpl.harvestCooldownImpl = 0
	gameImpl.jobsImpl = []catastrophe.Job{}
	gameImpl.lowerHarvestAmountImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.monumentCostMultImpl = 0
	gameImpl.monumentMaterialsImpl = 0
	gameImpl.neutralMaterialsImpl = 0
	gameImpl.playersImpl = []catastrophe.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.shelterMaterialsImpl = 0
	gameImpl.startingFoodImpl = 0
	gameImpl.starvingEnergyMultImpl = 0
	gameImpl.structuresImpl = []catastrophe.Structure{}
	gameImpl.tilesImpl = []catastrophe.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.turnsBetweenHarvestsImpl = 0
	gameImpl.turnsToCreateHumanImpl = 0
	gameImpl.turnsToLowerHarvestImpl = 0
	gameImpl.unitsImpl = []catastrophe.Unit{}
	gameImpl.wallMaterialsImpl = 0
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

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "catEnergyMult":
		gameImpl.catEnergyMultImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = catastropheDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = catastropheDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "harvestCooldown":
		gameImpl.harvestCooldownImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "jobs":
		gameImpl.jobsImpl = catastropheDeltaMerge.ArrayOfJob(&gameImpl.jobsImpl, delta)
		return true, nil
	case "lowerHarvestAmount":
		gameImpl.lowerHarvestAmountImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "monumentCostMult":
		gameImpl.monumentCostMultImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "monumentMaterials":
		gameImpl.monumentMaterialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "neutralMaterials":
		gameImpl.neutralMaterialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = catastropheDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "shelterMaterials":
		gameImpl.shelterMaterialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "startingFood":
		gameImpl.startingFoodImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "starvingEnergyMult":
		gameImpl.starvingEnergyMultImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "structures":
		gameImpl.structuresImpl = catastropheDeltaMerge.ArrayOfStructure(&gameImpl.structuresImpl, delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = catastropheDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "turnsBetweenHarvests":
		gameImpl.turnsBetweenHarvestsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "turnsToCreateHuman":
		gameImpl.turnsToCreateHumanImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "turnsToLowerHarvest":
		gameImpl.turnsToLowerHarvestImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "units":
		gameImpl.unitsImpl = catastropheDeltaMerge.ArrayOfUnit(&gameImpl.unitsImpl, delta)
		return true, nil
	case "wallMaterials":
		gameImpl.wallMaterialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) catastrophe.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x + y * gameImpl.mapWidthImpl]
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// GameImpl is the struct that implements the container for Game instances
// in Coreminer.
type GameImpl struct {
	base.GameImpl

	bombPriceImpl               int64
	bombSizeImpl                int64
	bombsImpl                   []coreminer.Bomb
	buildingMaterialPriceImpl   int64
	currentPlayerImpl           coreminer.Player
	currentTurnImpl             int64
	dirtPriceImpl               int64
	fallDamageImpl              int64
	fallWeightDamageImpl        int64
	gameObjectsImpl             map[string]coreminer.GameObject
	ladderCostImpl              int64
	ladderHealthImpl            int64
	largeCargoSizeImpl          int64
	largeMaterialSizeImpl       int64
	mapHeightImpl               int64
	mapWidthImpl                int64
	maxShieldingImpl            int64
	maxTurnsImpl                int64
	maxUpgradeLevelImpl         int64
	minersImpl                  []coreminer.Miner
	orePriceImpl                int64
	oreValueImpl                int64
	playersImpl                 []coreminer.Player
	sessionImpl                 string
	shieldCostImpl              int64
	shieldHealthImpl            int64
	spawnPriceImpl              int64
	suffocationDamageImpl       int64
	suffocationWeightDamageImpl int64
	supportCostImpl             int64
	supportHealthImpl           int64
	tilesImpl                   []coreminer.Tile
	timeAddedPerTurnImpl        float64
	upgradePriceImpl            int64
	upgradesImpl                []coreminer.Upgrade
	victoryAmountImpl           int64
}

// BombPrice returns the monetary price of a bomb when bought or sold.
func (gameImpl *GameImpl) BombPrice() int64 {
	return gameImpl.bombPriceImpl
}

// BombSize returns the amount of cargo space taken up by a Bomb.
func (gameImpl *GameImpl) BombSize() int64 {
	return gameImpl.bombSizeImpl
}

// Bombs returns every Bomb in the game.
func (gameImpl *GameImpl) Bombs() []coreminer.Bomb {
	return gameImpl.bombsImpl
}

// BuildingMaterialPrice returns the monetary price of building materials
// when bought.
func (gameImpl *GameImpl) BuildingMaterialPrice() int64 {
	return gameImpl.buildingMaterialPriceImpl
}

// CurrentPlayer returns the player whose turn it is currently. That player
// can send commands. Other players cannot.
func (gameImpl *GameImpl) CurrentPlayer() coreminer.Player {
	return gameImpl.currentPlayerImpl
}

// CurrentTurn returns the current turn number, starting at 0 for the first
// player's turn.
func (gameImpl *GameImpl) CurrentTurn() int64 {
	return gameImpl.currentTurnImpl
}

// DirtPrice returns the monetary price of dirt when bought or sold.
func (gameImpl *GameImpl) DirtPrice() int64 {
	return gameImpl.dirtPriceImpl
}

// FallDamage returns the amount of damage taken per Tile fallen.
func (gameImpl *GameImpl) FallDamage() int64 {
	return gameImpl.fallDamageImpl
}

// FallWeightDamage returns the amount of extra damage taken for falling
// while carrying a large amount of cargo.
func (gameImpl *GameImpl) FallWeightDamage() int64 {
	return gameImpl.fallWeightDamageImpl
}

// GameObjects returns a mapping of every game object's ID to the actual
// game object. Primarily used by the server and client to easily refer to
// the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]coreminer.GameObject {
	return gameImpl.gameObjectsImpl
}

// LadderCost returns the amount of building material required to build a
// ladder.
func (gameImpl *GameImpl) LadderCost() int64 {
	return gameImpl.ladderCostImpl
}

// LadderHealth returns the amount of mining power needed to remove a
// ladder from a Tile.
func (gameImpl *GameImpl) LadderHealth() int64 {
	return gameImpl.ladderHealthImpl
}

// LargeCargoSize returns the amount deemed as a large amount of cargo.
func (gameImpl *GameImpl) LargeCargoSize() int64 {
	return gameImpl.largeCargoSizeImpl
}

// LargeMaterialSize returns the amount deemed as a large amount of
// material.
func (gameImpl *GameImpl) LargeMaterialSize() int64 {
	return gameImpl.largeMaterialSizeImpl
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

// MaxShielding returns the maximum amount of shielding possible on a Tile.
func (gameImpl *GameImpl) MaxShielding() int64 {
	return gameImpl.maxShieldingImpl
}

// MaxTurns returns the maximum number of turns before the game will
// automatically end.
func (gameImpl *GameImpl) MaxTurns() int64 {
	return gameImpl.maxTurnsImpl
}

// MaxUpgradeLevel returns the highest upgrade level allowed on a Miner.
func (gameImpl *GameImpl) MaxUpgradeLevel() int64 {
	return gameImpl.maxUpgradeLevelImpl
}

// Miners returns every Miner in the game.
func (gameImpl *GameImpl) Miners() []coreminer.Miner {
	return gameImpl.minersImpl
}

// OrePrice returns the amount of money awarded when ore is dumped in the
// base and sold.
func (gameImpl *GameImpl) OrePrice() int64 {
	return gameImpl.orePriceImpl
}

// OreValue returns the amount of value awarded when ore is dumped in the
// base and sold.
func (gameImpl *GameImpl) OreValue() int64 {
	return gameImpl.oreValueImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []coreminer.Player {
	return gameImpl.playersImpl
}

// Session returns a unique identifier for the game instance that is being
// played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.sessionImpl
}

// ShieldCost returns the amount of building material required to shield a
// Tile.
func (gameImpl *GameImpl) ShieldCost() int64 {
	return gameImpl.shieldCostImpl
}

// ShieldHealth returns the amount of mining power needed to remove one
// unit of shielding off a Tile.
func (gameImpl *GameImpl) ShieldHealth() int64 {
	return gameImpl.shieldHealthImpl
}

// SpawnPrice returns the monetary price of spawning a Miner.
func (gameImpl *GameImpl) SpawnPrice() int64 {
	return gameImpl.spawnPriceImpl
}

// SuffocationDamage returns the amount of damage taken when suffocating
// inside a filled Tile.
func (gameImpl *GameImpl) SuffocationDamage() int64 {
	return gameImpl.suffocationDamageImpl
}

// SuffocationWeightDamage returns the amount of extra damage taken for
// suffocating under a large amount of material.
func (gameImpl *GameImpl) SuffocationWeightDamage() int64 {
	return gameImpl.suffocationWeightDamageImpl
}

// SupportCost returns the amount of building material required to build a
// support.
func (gameImpl *GameImpl) SupportCost() int64 {
	return gameImpl.supportCostImpl
}

// SupportHealth returns the amount of mining power needed to remove a
// support from a Tile.
func (gameImpl *GameImpl) SupportHealth() int64 {
	return gameImpl.supportHealthImpl
}

// Tiles returns all the tiles in the map, stored in Row-major order. Use
// `x + y * mapWidth` to access the correct index.
func (gameImpl *GameImpl) Tiles() []coreminer.Tile {
	return gameImpl.tilesImpl
}

// TimeAddedPerTurn returns the amount of time (in nano-seconds) added
// after each player performs a turn.
func (gameImpl *GameImpl) TimeAddedPerTurn() float64 {
	return gameImpl.timeAddedPerTurnImpl
}

// UpgradePrice returns the cost to upgrade a Miner.
func (gameImpl *GameImpl) UpgradePrice() int64 {
	return gameImpl.upgradePriceImpl
}

// Upgrades returns every Upgrade for a Miner in the game.
func (gameImpl *GameImpl) Upgrades() []coreminer.Upgrade {
	return gameImpl.upgradesImpl
}

// VictoryAmount returns the amount of victory points (value) required to
// win.
func (gameImpl *GameImpl) VictoryAmount() int64 {
	return gameImpl.victoryAmountImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.bombPriceImpl = 0
	gameImpl.bombSizeImpl = 0
	gameImpl.bombsImpl = []coreminer.Bomb{}
	gameImpl.buildingMaterialPriceImpl = 0
	gameImpl.currentPlayerImpl = nil
	gameImpl.currentTurnImpl = 0
	gameImpl.dirtPriceImpl = 0
	gameImpl.fallDamageImpl = 0
	gameImpl.fallWeightDamageImpl = 0
	gameImpl.gameObjectsImpl = make(map[string]coreminer.GameObject)
	gameImpl.ladderCostImpl = 0
	gameImpl.ladderHealthImpl = 0
	gameImpl.largeCargoSizeImpl = 0
	gameImpl.largeMaterialSizeImpl = 0
	gameImpl.mapHeightImpl = 0
	gameImpl.mapWidthImpl = 0
	gameImpl.maxShieldingImpl = 0
	gameImpl.maxTurnsImpl = 0
	gameImpl.maxUpgradeLevelImpl = 0
	gameImpl.minersImpl = []coreminer.Miner{}
	gameImpl.orePriceImpl = 0
	gameImpl.oreValueImpl = 0
	gameImpl.playersImpl = []coreminer.Player{}
	gameImpl.sessionImpl = ""
	gameImpl.shieldCostImpl = 0
	gameImpl.shieldHealthImpl = 0
	gameImpl.spawnPriceImpl = 0
	gameImpl.suffocationDamageImpl = 0
	gameImpl.suffocationWeightDamageImpl = 0
	gameImpl.supportCostImpl = 0
	gameImpl.supportHealthImpl = 0
	gameImpl.tilesImpl = []coreminer.Tile{}
	gameImpl.timeAddedPerTurnImpl = 0
	gameImpl.upgradePriceImpl = 0
	gameImpl.upgradesImpl = []coreminer.Upgrade{}
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

	coreminerDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'coreminer.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bombPrice":
		gameImpl.bombPriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "bombSize":
		gameImpl.bombSizeImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "bombs":
		gameImpl.bombsImpl = coreminerDeltaMerge.ArrayOfBomb(&gameImpl.bombsImpl, delta)
		return true, nil
	case "buildingMaterialPrice":
		gameImpl.buildingMaterialPriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "currentPlayer":
		gameImpl.currentPlayerImpl = coreminerDeltaMerge.Player(delta)
		return true, nil
	case "currentTurn":
		gameImpl.currentTurnImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "dirtPrice":
		gameImpl.dirtPriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "fallDamage":
		gameImpl.fallDamageImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "fallWeightDamage":
		gameImpl.fallWeightDamageImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "gameObjects":
		gameImpl.gameObjectsImpl = coreminerDeltaMerge.MapOfStringToGameObject(&gameImpl.gameObjectsImpl, delta)
		return true, nil
	case "ladderCost":
		gameImpl.ladderCostImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "ladderHealth":
		gameImpl.ladderHealthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "largeCargoSize":
		gameImpl.largeCargoSizeImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "largeMaterialSize":
		gameImpl.largeMaterialSizeImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "mapHeight":
		gameImpl.mapHeightImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "mapWidth":
		gameImpl.mapWidthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "maxShielding":
		gameImpl.maxShieldingImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "maxTurns":
		gameImpl.maxTurnsImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "maxUpgradeLevel":
		gameImpl.maxUpgradeLevelImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "miners":
		gameImpl.minersImpl = coreminerDeltaMerge.ArrayOfMiner(&gameImpl.minersImpl, delta)
		return true, nil
	case "orePrice":
		gameImpl.orePriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "oreValue":
		gameImpl.oreValueImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "players":
		gameImpl.playersImpl = coreminerDeltaMerge.ArrayOfPlayer(&gameImpl.playersImpl, delta)
		return true, nil
	case "session":
		gameImpl.sessionImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	case "shieldCost":
		gameImpl.shieldCostImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "shieldHealth":
		gameImpl.shieldHealthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "spawnPrice":
		gameImpl.spawnPriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "suffocationDamage":
		gameImpl.suffocationDamageImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "suffocationWeightDamage":
		gameImpl.suffocationWeightDamageImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "supportCost":
		gameImpl.supportCostImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "supportHealth":
		gameImpl.supportHealthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "tiles":
		gameImpl.tilesImpl = coreminerDeltaMerge.ArrayOfTile(&gameImpl.tilesImpl, delta)
		return true, nil
	case "timeAddedPerTurn":
		gameImpl.timeAddedPerTurnImpl = coreminerDeltaMerge.Float(delta)
		return true, nil
	case "upgradePrice":
		gameImpl.upgradePriceImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "upgrades":
		gameImpl.upgradesImpl = coreminerDeltaMerge.ArrayOfUpgrade(&gameImpl.upgradesImpl, delta)
		return true, nil
	case "victoryAmount":
		gameImpl.victoryAmountImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

// -- Tiled Game Utils -- \\

// GetTileAt returns the Tile at a give position (x, y).
func (gameImpl *GameImpl) GetTileAt(x int64, y int64) coreminer.Tile {
	if x < 0 || y < 0 || x >= gameImpl.mapWidthImpl || y >= gameImpl.mapHeightImpl {
		// out of bounds
		return nil
	}

	return gameImpl.tilesImpl[x+y*gameImpl.mapWidthImpl]
}

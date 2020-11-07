package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Coreminer.
type PlayerImpl struct {
	GameObjectImpl

	baseTileImpl      coreminer.Tile
	bombsImpl         []coreminer.Bomb
	clientTypeImpl    string
	hopperTilesImpl   []coreminer.Tile
	lostImpl          bool
	minersImpl        []coreminer.Miner
	moneyImpl         int64
	nameImpl          string
	opponentImpl      coreminer.Player
	reasonLostImpl    string
	reasonWonImpl     string
	timeRemainingImpl float64
	valueImpl         int64
	wonImpl           bool
}

// BaseTile returns the Tile this Player's base is on.
func (playerImpl *PlayerImpl) BaseTile() coreminer.Tile {
	return playerImpl.baseTileImpl
}

// Bombs returns every Bomb owned by this Player.
func (playerImpl *PlayerImpl) Bombs() []coreminer.Bomb {
	return playerImpl.bombsImpl
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// HopperTiles returns the Tiles this Player's hoppers are on.
func (playerImpl *PlayerImpl) HopperTiles() []coreminer.Tile {
	return playerImpl.hopperTilesImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Miners returns every Miner owned by this Player.
func (playerImpl *PlayerImpl) Miners() []coreminer.Miner {
	return playerImpl.minersImpl
}

// Money returns the amount of money this Player currently has.
func (playerImpl *PlayerImpl) Money() int64 {
	return playerImpl.moneyImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() coreminer.Player {
	return playerImpl.opponentImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Value returns the amount of value (victory points) this Player has
// gained.
func (playerImpl *PlayerImpl) Value() int64 {
	return playerImpl.valueImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// SpawnMiner runs logic that spawns a Miner on this Player's base Tile.
func (playerImpl *PlayerImpl) SpawnMiner() bool {
	return playerImpl.RunOnServer("spawnMiner", map[string]interface{}{}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.baseTileImpl = nil
	playerImpl.bombsImpl = []coreminer.Bomb{}
	playerImpl.clientTypeImpl = ""
	playerImpl.hopperTilesImpl = []coreminer.Tile{}
	playerImpl.lostImpl = true
	playerImpl.minersImpl = []coreminer.Miner{}
	playerImpl.moneyImpl = 0
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.valueImpl = 0
	playerImpl.wonImpl = true
}

// DeltaMerge merges the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := playerImpl.GameObjectImpl.DeltaMerge(
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
	case "baseTile":
		playerImpl.baseTileImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "bombs":
		playerImpl.bombsImpl = coreminerDeltaMerge.ArrayOfBomb(&playerImpl.bombsImpl, delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	case "hopperTiles":
		playerImpl.hopperTilesImpl = coreminerDeltaMerge.ArrayOfTile(&playerImpl.hopperTilesImpl, delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	case "miners":
		playerImpl.minersImpl = coreminerDeltaMerge.ArrayOfMiner(&playerImpl.minersImpl, delta)
		return true, nil
	case "money":
		playerImpl.moneyImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = coreminerDeltaMerge.Player(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = coreminerDeltaMerge.Float(delta)
		return true, nil
	case "value":
		playerImpl.valueImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = coreminerDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/coreminer"
)

// MinerImpl is the struct that implements the container for Miner
// instances in Coreminer.
type MinerImpl struct {
	GameObjectImpl

	bombsImpl             int64
	buildingMaterialsImpl int64
	currentUpgradeImpl    coreminer.Upgrade
	dirtImpl              int64
	healthImpl            int64
	miningPowerImpl       int64
	movesImpl             int64
	oreImpl               int64
	ownerImpl             coreminer.Player
	tileImpl              coreminer.Tile
	upgradeLevelImpl      int64
}

// Bombs returns the number of bombs being carried by this Miner.
func (minerImpl *MinerImpl) Bombs() int64 {
	return minerImpl.bombsImpl
}

// BuildingMaterials returns the number of building materials carried by
// this Miner.
func (minerImpl *MinerImpl) BuildingMaterials() int64 {
	return minerImpl.buildingMaterialsImpl
}

// CurrentUpgrade returns the Upgrade this Miner is on.
func (minerImpl *MinerImpl) CurrentUpgrade() coreminer.Upgrade {
	return minerImpl.currentUpgradeImpl
}

// Dirt returns the amount of dirt carried by this Miner.
func (minerImpl *MinerImpl) Dirt() int64 {
	return minerImpl.dirtImpl
}

// Health returns the remaining health of this Miner.
func (minerImpl *MinerImpl) Health() int64 {
	return minerImpl.healthImpl
}

// MiningPower returns the remaining mining power this Miner has this turn.
func (minerImpl *MinerImpl) MiningPower() int64 {
	return minerImpl.miningPowerImpl
}

// Moves returns the number of moves this Miner has left this turn.
func (minerImpl *MinerImpl) Moves() int64 {
	return minerImpl.movesImpl
}

// Ore returns the amount of ore carried by this Miner.
func (minerImpl *MinerImpl) Ore() int64 {
	return minerImpl.oreImpl
}

// Owner returns the Player that owns and can control this Miner.
func (minerImpl *MinerImpl) Owner() coreminer.Player {
	return minerImpl.ownerImpl
}

// Tile returns the Tile this Miner is on.
//
// Value can be returned as a nil pointer.
func (minerImpl *MinerImpl) Tile() coreminer.Tile {
	return minerImpl.tileImpl
}

// UpgradeLevel returns the upgrade level of this Miner. Starts at 0.
func (minerImpl *MinerImpl) UpgradeLevel() int64 {
	return minerImpl.upgradeLevelImpl
}

// Build runs logic that builds a support, shield, or ladder on Miner's
// Tile, or an adjacent Tile.
func (minerImpl *MinerImpl) Build(tile coreminer.Tile, typeVar string) bool {
	return minerImpl.RunOnServer("build", map[string]interface{}{
		"tile": tile,
		"type": typeVar,
	}).(bool)
}

// Buy runs logic that purchase a resource from the Player's base or
// hopper.
func (minerImpl *MinerImpl) Buy(resource string, amount int64) bool {
	return minerImpl.RunOnServer("buy", map[string]interface{}{
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// Dump runs logic that dumps materials from cargo to an adjacent Tile. If
// the Tile is a base or a hopper Tile, materials are sold instead of
// placed.
func (minerImpl *MinerImpl) Dump(tile coreminer.Tile, material string, amount int64) bool {
	return minerImpl.RunOnServer("dump", map[string]interface{}{
		"tile":     tile,
		"material": material,
		"amount":   amount,
	}).(bool)
}

// Mine runs logic that mines the Tile the Miner is on or an adjacent Tile.
func (minerImpl *MinerImpl) Mine(tile coreminer.Tile, amount int64) bool {
	return minerImpl.RunOnServer("mine", map[string]interface{}{
		"tile":   tile,
		"amount": amount,
	}).(bool)
}

// Move runs logic that moves this Miner from its current Tile to an
// adjacent Tile.
func (minerImpl *MinerImpl) Move(tile coreminer.Tile) bool {
	return minerImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Transfer runs logic that transfers a resource from the one Miner to
// another.
func (minerImpl *MinerImpl) Transfer(miner coreminer.Miner, resource string, amount int64) bool {
	return minerImpl.RunOnServer("transfer", map[string]interface{}{
		"miner":    miner,
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// Upgrade runs logic that upgrade this Miner by installing an upgrade
// module.
func (minerImpl *MinerImpl) Upgrade() bool {
	return minerImpl.RunOnServer("upgrade", map[string]interface{}{}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Miner.
func (minerImpl *MinerImpl) InitImplDefaults() {
	minerImpl.GameObjectImpl.InitImplDefaults()

	minerImpl.bombsImpl = 0
	minerImpl.buildingMaterialsImpl = 0
	minerImpl.currentUpgradeImpl = nil
	minerImpl.dirtImpl = 0
	minerImpl.healthImpl = 0
	minerImpl.miningPowerImpl = 0
	minerImpl.movesImpl = 0
	minerImpl.oreImpl = 0
	minerImpl.ownerImpl = nil
	minerImpl.tileImpl = nil
	minerImpl.upgradeLevelImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Miner.
func (minerImpl *MinerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := minerImpl.GameObjectImpl.DeltaMerge(
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
	case "bombs":
		minerImpl.bombsImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "buildingMaterials":
		minerImpl.buildingMaterialsImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "currentUpgrade":
		minerImpl.currentUpgradeImpl = coreminerDeltaMerge.Upgrade(delta)
		return true, nil
	case "dirt":
		minerImpl.dirtImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "health":
		minerImpl.healthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "miningPower":
		minerImpl.miningPowerImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		minerImpl.movesImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "ore":
		minerImpl.oreImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		minerImpl.ownerImpl = coreminerDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		minerImpl.tileImpl = coreminerDeltaMerge.Tile(delta)
		return true, nil
	case "upgradeLevel":
		minerImpl.upgradeLevelImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

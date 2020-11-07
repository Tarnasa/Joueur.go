package impl

import (
	"errors"
	"joueur/base"
)

// UpgradeImpl is the struct that implements the container for Upgrade
// instances in Coreminer.
type UpgradeImpl struct {
	GameObjectImpl

	cargoCapacityImpl int64
	healthImpl        int64
	miningPowerImpl   int64
	movesImpl         int64
	titleImpl         string
}

// CargoCapacity returns the amount of cargo capacity this Upgrade has.
func (upgradeImpl *UpgradeImpl) CargoCapacity() int64 {
	return upgradeImpl.cargoCapacityImpl
}

// Health returns the maximum amount of health this Upgrade has.
func (upgradeImpl *UpgradeImpl) Health() int64 {
	return upgradeImpl.healthImpl
}

// MiningPower returns the amount of mining power this Upgrade has per
// turn.
func (upgradeImpl *UpgradeImpl) MiningPower() int64 {
	return upgradeImpl.miningPowerImpl
}

// Moves returns the number of moves this Upgrade can make per turn.
func (upgradeImpl *UpgradeImpl) Moves() int64 {
	return upgradeImpl.movesImpl
}

// Title returns the Upgrade title.
func (upgradeImpl *UpgradeImpl) Title() string {
	return upgradeImpl.titleImpl
}

// InitImplDefaults initializes safe defaults for all fields in Upgrade.
func (upgradeImpl *UpgradeImpl) InitImplDefaults() {
	upgradeImpl.GameObjectImpl.InitImplDefaults()

	upgradeImpl.cargoCapacityImpl = 0
	upgradeImpl.healthImpl = 0
	upgradeImpl.miningPowerImpl = 0
	upgradeImpl.movesImpl = 0
	upgradeImpl.titleImpl = ""
}

// DeltaMerge merges the delta for a given attribute in Upgrade.
func (upgradeImpl *UpgradeImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := upgradeImpl.GameObjectImpl.DeltaMerge(
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
	case "cargoCapacity":
		upgradeImpl.cargoCapacityImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "health":
		upgradeImpl.healthImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "miningPower":
		upgradeImpl.miningPowerImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		upgradeImpl.movesImpl = coreminerDeltaMerge.Int(delta)
		return true, nil
	case "title":
		upgradeImpl.titleImpl = coreminerDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

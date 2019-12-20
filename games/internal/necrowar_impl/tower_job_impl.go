package impl
import (
	"errors"
	"joueur/base"
)

// TowerJobImpl is the struct that implements the container for TowerJob
// instances in Necrowar.
type TowerJobImpl struct {
	GameObjectImpl

	allUnitsImpl            bool
	damageImpl              int64
	goldCostImpl            int64
	healthImpl              int64
	manaCostImpl            int64
	rangeImpl               int64
	titleImpl               string
	turnsBetweenAttacksImpl int64
}

// AllUnits returns whether this tower type hits all of the units on a tile
// (true) or one at a time (false).
func (towerJobImpl *TowerJobImpl) AllUnits() bool {
	return towerJobImpl.allUnitsImpl
}

// Damage returns the amount of damage this type does per attack.
func (towerJobImpl *TowerJobImpl) Damage() int64 {
	return towerJobImpl.damageImpl
}

// GoldCost returns how much does this type cost in gold.
func (towerJobImpl *TowerJobImpl) GoldCost() int64 {
	return towerJobImpl.goldCostImpl
}

// Health returns the amount of starting health this type has.
func (towerJobImpl *TowerJobImpl) Health() int64 {
	return towerJobImpl.healthImpl
}

// ManaCost returns how much does this type cost in mana.
func (towerJobImpl *TowerJobImpl) ManaCost() int64 {
	return towerJobImpl.manaCostImpl
}

// Range returns the number of tiles this type can attack from.
func (towerJobImpl *TowerJobImpl) Range() int64 {
	return towerJobImpl.rangeImpl
}

// Title returns the type title. 'arrow', 'aoe', 'balarraya', 'cleansing',
// or 'castle'.
//
// Literal Values: "arrow", "aoe", "ballista", "cleansing", or "castle"
func (towerJobImpl *TowerJobImpl) Title() string {
	return towerJobImpl.titleImpl
}

// TurnsBetweenAttacks returns how many turns have to take place between
// this type's attacks.
func (towerJobImpl *TowerJobImpl) TurnsBetweenAttacks() int64 {
	return towerJobImpl.turnsBetweenAttacksImpl
}

// InitImplDefaults initializes safe defaults for all fields in TowerJob.
func (towerJobImpl *TowerJobImpl) InitImplDefaults() {
	towerJobImpl.GameObjectImpl.InitImplDefaults()

	towerJobImpl.allUnitsImpl = true
	towerJobImpl.damageImpl = 0
	towerJobImpl.goldCostImpl = 0
	towerJobImpl.healthImpl = 0
	towerJobImpl.manaCostImpl = 0
	towerJobImpl.rangeImpl = 0
	towerJobImpl.titleImpl = ""
	towerJobImpl.turnsBetweenAttacksImpl = 0
}

// DeltaMerge merges the delta for a given attribute in TowerJob.
func (towerJobImpl *TowerJobImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := towerJobImpl.GameObjectImpl.DeltaMerge(
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
	case "allUnits":
		towerJobImpl.allUnitsImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "damage":
		towerJobImpl.damageImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "goldCost":
		towerJobImpl.goldCostImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "health":
		towerJobImpl.healthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "manaCost":
		towerJobImpl.manaCostImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "range":
		towerJobImpl.rangeImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "title":
		towerJobImpl.titleImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	case "turnsBetweenAttacks":
		towerJobImpl.turnsBetweenAttacksImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

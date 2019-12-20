package impl
import (
	"errors"
	"joueur/base"
)

// UnitJobImpl is the struct that implements the container for UnitJob
// instances in Necrowar.
type UnitJobImpl struct {
	GameObjectImpl

	damageImpl   int64
	goldCostImpl int64
	healthImpl   int64
	manaCostImpl int64
	movesImpl    int64
	perTileImpl  int64
	rangeImpl    int64
	titleImpl    string
}

// Damage returns the amount of damage this type does per attack.
func (unitJobImpl *UnitJobImpl) Damage() int64 {
	return unitJobImpl.damageImpl
}

// GoldCost returns how much does this type cost in gold.
func (unitJobImpl *UnitJobImpl) GoldCost() int64 {
	return unitJobImpl.goldCostImpl
}

// Health returns the amount of starting health this type has.
func (unitJobImpl *UnitJobImpl) Health() int64 {
	return unitJobImpl.healthImpl
}

// ManaCost returns how much does this type cost in mana.
func (unitJobImpl *UnitJobImpl) ManaCost() int64 {
	return unitJobImpl.manaCostImpl
}

// Moves returns the number of moves this type can make per turn.
func (unitJobImpl *UnitJobImpl) Moves() int64 {
	return unitJobImpl.movesImpl
}

// PerTile returns how many of this type of unit can take up one tile.
func (unitJobImpl *UnitJobImpl) PerTile() int64 {
	return unitJobImpl.perTileImpl
}

// Range returns amount of tiles away this type has to be in order to be
// effective.
func (unitJobImpl *UnitJobImpl) Range() int64 {
	return unitJobImpl.rangeImpl
}

// Title returns the type title. 'worker', 'zombie', 'ghoul', 'hound',
// 'abomination', 'wraith' or 'horseman'.
//
// Literal Values: "worker", "zombie", "ghoul", "hound", "abomination",
// "wraith", or "horseman"
func (unitJobImpl *UnitJobImpl) Title() string {
	return unitJobImpl.titleImpl
}

// InitImplDefaults initializes safe defaults for all fields in UnitJob.
func (unitJobImpl *UnitJobImpl) InitImplDefaults() {
	unitJobImpl.GameObjectImpl.InitImplDefaults()

	unitJobImpl.damageImpl = 0
	unitJobImpl.goldCostImpl = 0
	unitJobImpl.healthImpl = 0
	unitJobImpl.manaCostImpl = 0
	unitJobImpl.movesImpl = 0
	unitJobImpl.perTileImpl = 0
	unitJobImpl.rangeImpl = 0
	unitJobImpl.titleImpl = ""
}

// DeltaMerge merges the delta for a given attribute in UnitJob.
func (unitJobImpl *UnitJobImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := unitJobImpl.GameObjectImpl.DeltaMerge(
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
	case "damage":
		unitJobImpl.damageImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "goldCost":
		unitJobImpl.goldCostImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "health":
		unitJobImpl.healthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "manaCost":
		unitJobImpl.manaCostImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		unitJobImpl.movesImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "perTile":
		unitJobImpl.perTileImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "range":
		unitJobImpl.rangeImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "title":
		unitJobImpl.titleImpl = necrowarDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

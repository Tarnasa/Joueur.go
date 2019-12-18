package impl
import (
	"errors"
	"joueur/base"
)

// JobImpl is the struct that implements the container for Job instances in
// Stardash.
type JobImpl struct {
	GameObjectImpl

	carryLimitImpl int64
	damageImpl     int64
	energyImpl     int64
	movesImpl      int64
	rangeImpl      int64
	shieldImpl     int64
	titleImpl      string
	unitCostImpl   int64
}

// CarryLimit returns how many combined resources a unit with this Job can
// hold at once.
func (jobImpl *JobImpl) CarryLimit() int64 {
	return jobImpl.carryLimitImpl
}

// Damage returns the amount of damage this Job does per attack.
func (jobImpl *JobImpl) Damage() int64 {
	return jobImpl.damageImpl
}

// Energy returns the amount of starting health this Job has.
func (jobImpl *JobImpl) Energy() int64 {
	return jobImpl.energyImpl
}

// Moves returns the distance this job can move per turn.
func (jobImpl *JobImpl) Moves() int64 {
	return jobImpl.movesImpl
}

// Range returns the distance at which this job can effect things.
func (jobImpl *JobImpl) Range() int64 {
	return jobImpl.rangeImpl
}

// Shield returns the reserve the martyr use to protect allies.
func (jobImpl *JobImpl) Shield() int64 {
	return jobImpl.shieldImpl
}

// Title returns the Job title. 'corvette', 'missileboat', 'martyr',
// 'transport', or 'miner'. (in this order from 0-4).
// Literal Values: "corvette", "missileboat", "martyr", "transport", or
// "miner"
func (jobImpl *JobImpl) Title() string {
	return jobImpl.titleImpl
}

// UnitCost returns how much money it costs to spawn a unit.
func (jobImpl *JobImpl) UnitCost() int64 {
	return jobImpl.unitCostImpl
}

// InitImplDefaults initializes safe defaults for all fields in Job.
func (jobImpl *JobImpl) InitImplDefaults() {
	jobImpl.GameObjectImpl.InitImplDefaults()

	jobImpl.carryLimitImpl = 0
	jobImpl.damageImpl = 0
	jobImpl.energyImpl = 0
	jobImpl.movesImpl = 0
	jobImpl.rangeImpl = 0
	jobImpl.shieldImpl = 0
	jobImpl.titleImpl = ""
	jobImpl.unitCostImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Job.
func (jobImpl *JobImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := jobImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	stardashDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'stardash.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "carryLimit":
		jobImpl.carryLimitImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "damage":
		jobImpl.damageImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "energy":
		jobImpl.energyImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		jobImpl.movesImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "range":
		jobImpl.rangeImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "shield":
		jobImpl.shieldImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "title":
		jobImpl.titleImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "unitCost":
		jobImpl.unitCostImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

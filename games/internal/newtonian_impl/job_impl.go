package impl

import (
	"errors"
	"joueur/base"
)

// JobImpl is the struct that implements the container for Job instances in
// Newtonian.
type JobImpl struct {
	GameObjectImpl

	carryLimitImpl int64
	damageImpl     int64
	healthImpl     int64
	movesImpl      int64
	titleImpl      string
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

// Health returns the amount of starting health this Job has.
func (jobImpl *JobImpl) Health() int64 {
	return jobImpl.healthImpl
}

// Moves returns the number of moves this Job can make per turn.
func (jobImpl *JobImpl) Moves() int64 {
	return jobImpl.movesImpl
}

// Title returns the Job title. 'intern', 'manager', or 'physicist'.
//
// Literal Values: "intern", "manager", or "physicist"
func (jobImpl *JobImpl) Title() string {
	return jobImpl.titleImpl
}

// InitImplDefaults initializes safe defaults for all fields in Job.
func (jobImpl *JobImpl) InitImplDefaults() {
	jobImpl.GameObjectImpl.InitImplDefaults()

	jobImpl.carryLimitImpl = 0
	jobImpl.damageImpl = 0
	jobImpl.healthImpl = 0
	jobImpl.movesImpl = 0
	jobImpl.titleImpl = ""
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

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "carryLimit":
		jobImpl.carryLimitImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "damage":
		jobImpl.damageImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "health":
		jobImpl.healthImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		jobImpl.movesImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "title":
		jobImpl.titleImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

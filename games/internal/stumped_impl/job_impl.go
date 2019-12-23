package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// JobImpl is the struct that implements the container for Job instances in
// Stumped.
type JobImpl struct {
	GameObjectImpl

	actionsImpl          int64
	carryLimitImpl       int64
	choppingImpl         int64
	costImpl             int64
	damageImpl           int64
	distractionPowerImpl int64
	healthImpl           int64
	movesImpl            int64
	munchingImpl         int64
	titleImpl            string
}

// Actions returns the number of actions this Job can make per turn.
func (jobImpl *JobImpl) Actions() int64 {
	return jobImpl.actionsImpl
}

// CarryLimit returns how many combined resources a beaver with this Job
// can hold at once.
func (jobImpl *JobImpl) CarryLimit() int64 {
	return jobImpl.carryLimitImpl
}

// Chopping returns scalar for how many branches this Job harvests at once.
func (jobImpl *JobImpl) Chopping() int64 {
	return jobImpl.choppingImpl
}

// Cost returns how much food this Job costs to recruit.
func (jobImpl *JobImpl) Cost() int64 {
	return jobImpl.costImpl
}

// Damage returns the amount of damage this Job does per attack.
func (jobImpl *JobImpl) Damage() int64 {
	return jobImpl.damageImpl
}

// DistractionPower returns how many turns a beaver attacked by this Job is
// distracted by.
func (jobImpl *JobImpl) DistractionPower() int64 {
	return jobImpl.distractionPowerImpl
}

// Health returns the amount of starting health this Job has.
func (jobImpl *JobImpl) Health() int64 {
	return jobImpl.healthImpl
}

// Moves returns the number of moves this Job can make per turn.
func (jobImpl *JobImpl) Moves() int64 {
	return jobImpl.movesImpl
}

// Munching returns scalar for how much food this Job harvests at once.
func (jobImpl *JobImpl) Munching() int64 {
	return jobImpl.munchingImpl
}

// Title returns the Job title.
func (jobImpl *JobImpl) Title() string {
	return jobImpl.titleImpl
}

// Recruit runs logic that recruits a Beaver of this Job to a lodge
func (jobImpl *JobImpl) Recruit(tile stumped.Tile) stumped.Beaver {
	return jobImpl.RunOnServer("recruit", map[string]interface{}{
		"tile": tile,
	}).(stumped.Beaver)
}

// InitImplDefaults initializes safe defaults for all fields in Job.
func (jobImpl *JobImpl) InitImplDefaults() {
	jobImpl.GameObjectImpl.InitImplDefaults()

	jobImpl.actionsImpl = 0
	jobImpl.carryLimitImpl = 0
	jobImpl.choppingImpl = 0
	jobImpl.costImpl = 0
	jobImpl.damageImpl = 0
	jobImpl.distractionPowerImpl = 0
	jobImpl.healthImpl = 0
	jobImpl.movesImpl = 0
	jobImpl.munchingImpl = 0
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

	stumpedDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stumped.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "actions":
		jobImpl.actionsImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "carryLimit":
		jobImpl.carryLimitImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "chopping":
		jobImpl.choppingImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "cost":
		jobImpl.costImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "damage":
		jobImpl.damageImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "distractionPower":
		jobImpl.distractionPowerImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "health":
		jobImpl.healthImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		jobImpl.movesImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "munching":
		jobImpl.munchingImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "title":
		jobImpl.titleImpl = stumpedDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

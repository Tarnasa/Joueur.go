package impl
import (
	"errors"
	"joueur/base"
)

// JobImpl is the struct that implements the container for Job instances in
// Catastrophe.
type JobImpl struct {
	GameObjectImpl

	actionCostImpl float64
	carryLimitImpl int64
	movesImpl      int64
	regenRateImpl  float64
	titleImpl      string
	upkeepImpl     int64
}

// ActionCost returns the amount of energy this Job normally uses to
// perform its actions.
func (jobImpl *JobImpl) ActionCost() float64 {
	return jobImpl.actionCostImpl
}

// CarryLimit returns how many combined resources a Unit with this Job can
// hold at once.
func (jobImpl *JobImpl) CarryLimit() int64 {
	return jobImpl.carryLimitImpl
}

// Moves returns the number of moves this Job can make per turn.
func (jobImpl *JobImpl) Moves() int64 {
	return jobImpl.movesImpl
}

// RegenRate returns the amount of energy normally regenerated when resting
// at a shelter.
func (jobImpl *JobImpl) RegenRate() float64 {
	return jobImpl.regenRateImpl
}

// Title returns the Job title.
// Literal Values: "fresh human", "cat overlord", "soldier", "gatherer",
// "builder", or "missionary"
func (jobImpl *JobImpl) Title() string {
	return jobImpl.titleImpl
}

// Upkeep returns the amount of food per turn this Unit consumes. If there
// isn't enough food for every Unit, all Units become starved and do not
// consume food.
func (jobImpl *JobImpl) Upkeep() int64 {
	return jobImpl.upkeepImpl
}

// InitImplDefaults initializes safe defaults for all fields in Job.
func (jobImpl *JobImpl) InitImplDefaults() {
	jobImpl.GameObjectImpl.InitImplDefaults()

	jobImpl.actionCostImpl = 0
	jobImpl.carryLimitImpl = 0
	jobImpl.movesImpl = 0
	jobImpl.regenRateImpl = 0
	jobImpl.titleImpl = ""
	jobImpl.upkeepImpl = 0
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

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "actionCost":
		jobImpl.actionCostImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "carryLimit":
		jobImpl.carryLimitImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		jobImpl.movesImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "regenRate":
		jobImpl.regenRateImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "title":
		jobImpl.titleImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	case "upkeep":
		jobImpl.upkeepImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

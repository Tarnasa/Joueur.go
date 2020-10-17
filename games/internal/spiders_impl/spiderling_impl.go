package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// SpiderlingImpl is the struct that implements the container for
// Spiderling instances in Spiders.
type SpiderlingImpl struct {
	SpiderImpl

	busyImpl              string
	movingOnWebImpl       spiders.Web
	movingToNestImpl      spiders.Nest
	numberOfCoworkersImpl int64
	workRemainingImpl     float64
}

// Busy returns when empty string this Spiderling is not busy, and can act.
// Otherwise a string representing what it is busy with, e.g. 'Moving',
// 'Attacking'.
//
// Literal Values: "", "Moving", "Attacking", "Strengthening", "Weakening",
// "Cutting", or "Spitting"
func (spiderlingImpl *SpiderlingImpl) Busy() string {
	return spiderlingImpl.busyImpl
}

// MovingOnWeb returns the Web this Spiderling is using to move. Nil if it
// is not moving.
//
// Value can be returned as a nil pointer.
func (spiderlingImpl *SpiderlingImpl) MovingOnWeb() spiders.Web {
	return spiderlingImpl.movingOnWebImpl
}

// MovingToNest returns the Nest this Spiderling is moving to. Nil if it is
// not moving.
//
// Value can be returned as a nil pointer.
func (spiderlingImpl *SpiderlingImpl) MovingToNest() spiders.Nest {
	return spiderlingImpl.movingToNestImpl
}

// NumberOfCoworkers returns the number of Spiderlings busy with the same
// work this Spiderling is doing, speeding up the task.
func (spiderlingImpl *SpiderlingImpl) NumberOfCoworkers() int64 {
	return spiderlingImpl.numberOfCoworkersImpl
}

// WorkRemaining returns how much work needs to be done for this Spiderling
// to finish being busy. See docs for the Work formula.
func (spiderlingImpl *SpiderlingImpl) WorkRemaining() float64 {
	return spiderlingImpl.workRemainingImpl
}

// Attack runs logic that attacks another Spiderling.
func (spiderlingImpl *SpiderlingImpl) Attack(spiderling spiders.Spiderling) bool {
	return spiderlingImpl.RunOnServer("attack", map[string]interface{}{
		"spiderling": spiderling,
	}).(bool)
}

// Move runs logic that starts moving the Spiderling across a Web to
// another Nest.
func (spiderlingImpl *SpiderlingImpl) Move(web spiders.Web) bool {
	return spiderlingImpl.RunOnServer("move", map[string]interface{}{
		"web": web,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Spiderling.
func (spiderlingImpl *SpiderlingImpl) InitImplDefaults() {
	spiderlingImpl.SpiderImpl.InitImplDefaults()

	spiderlingImpl.busyImpl = ""
	spiderlingImpl.movingOnWebImpl = nil
	spiderlingImpl.movingToNestImpl = nil
	spiderlingImpl.numberOfCoworkersImpl = 0
	spiderlingImpl.workRemainingImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Spiderling.
func (spiderlingImpl *SpiderlingImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := spiderlingImpl.SpiderImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	spidersDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'spiders.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "busy":
		spiderlingImpl.busyImpl = spidersDeltaMerge.String(delta)
		return true, nil
	case "movingOnWeb":
		spiderlingImpl.movingOnWebImpl = spidersDeltaMerge.Web(delta)
		return true, nil
	case "movingToNest":
		spiderlingImpl.movingToNestImpl = spidersDeltaMerge.Nest(delta)
		return true, nil
	case "numberOfCoworkers":
		spiderlingImpl.numberOfCoworkersImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "workRemaining":
		spiderlingImpl.workRemainingImpl = spidersDeltaMerge.Float(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

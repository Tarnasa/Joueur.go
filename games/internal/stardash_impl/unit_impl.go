package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// UnitImpl is the struct that implements the container for Unit instances
// in Stardash.
type UnitImpl struct {
	GameObjectImpl

	actedImpl       bool
	dashXImpl       float64
	dashYImpl       float64
	energyImpl      int64
	genariumImpl    int64
	isBusyImpl      bool
	jobImpl         stardash.Job
	legendariumImpl int64
	movesImpl       float64
	mythiciteImpl   int64
	ownerImpl       stardash.Player
	protectorImpl   stardash.Unit
	rariumImpl      int64
	shieldImpl      int64
	xImpl           float64
	yImpl           float64
}

// Acted returns whether or not this Unit has performed its action this
// turn.
func (unitImpl *UnitImpl) Acted() bool {
	return unitImpl.actedImpl
}

// DashX returns the x value this unit is dashing to.
func (unitImpl *UnitImpl) DashX() float64 {
	return unitImpl.dashXImpl
}

// DashY returns the y value this unit is dashing to.
func (unitImpl *UnitImpl) DashY() float64 {
	return unitImpl.dashYImpl
}

// Energy returns the remaining health of the unit.
func (unitImpl *UnitImpl) Energy() int64 {
	return unitImpl.energyImpl
}

// Genarium returns the amount of Genarium ore carried by this unit. (0 to
// job carry capacity - other carried items).
func (unitImpl *UnitImpl) Genarium() int64 {
	return unitImpl.genariumImpl
}

// IsBusy returns tracks wheither or not the ship is dashing or Mining. If
// true, it cannot do anything else.
func (unitImpl *UnitImpl) IsBusy() bool {
	return unitImpl.isBusyImpl
}

// Job returns the Job this Unit has.
func (unitImpl *UnitImpl) Job() stardash.Job {
	return unitImpl.jobImpl
}

// Legendarium returns the amount of Legendarium ore carried by this unit.
// (0 to job carry capacity - other carried items).
func (unitImpl *UnitImpl) Legendarium() int64 {
	return unitImpl.legendariumImpl
}

// Moves returns the distance this unit can still move.
func (unitImpl *UnitImpl) Moves() float64 {
	return unitImpl.movesImpl
}

// Mythicite returns the amount of Mythicite carried by this unit. (0 to
// job carry capacity - other carried items).
func (unitImpl *UnitImpl) Mythicite() int64 {
	return unitImpl.mythiciteImpl
}

// Owner returns the Player that owns and can control this Unit.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Owner() stardash.Player {
	return unitImpl.ownerImpl
}

// Protector returns the martyr ship that is currently shielding this ship
// if any.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Protector() stardash.Unit {
	return unitImpl.protectorImpl
}

// Rarium returns the amount of Rarium carried by this unit. (0 to job
// carry capacity - other carried items).
func (unitImpl *UnitImpl) Rarium() int64 {
	return unitImpl.rariumImpl
}

// Shield returns the sheild that a martyr ship has.
func (unitImpl *UnitImpl) Shield() int64 {
	return unitImpl.shieldImpl
}

// X returns the x value this unit is on.
func (unitImpl *UnitImpl) X() float64 {
	return unitImpl.xImpl
}

// Y returns the y value this unit is on.
func (unitImpl *UnitImpl) Y() float64 {
	return unitImpl.yImpl
}

// Attack runs logic that attacks the specified unit.
func (unitImpl *UnitImpl) Attack(enemy stardash.Unit) bool {
	return unitImpl.RunOnServer("attack", map[string]interface{}{
		"enemy": enemy,
	}).(bool)
}

// Dash runs logic that causes the unit to dash towards the designated
// destination.
func (unitImpl *UnitImpl) Dash(x float64, y float64) bool {
	return unitImpl.RunOnServer("dash", map[string]interface{}{
		"x": x,
		"y": y,
	}).(bool)
}

// Mine runs logic that allows a miner to mine a asteroid
func (unitImpl *UnitImpl) Mine(body stardash.Body) bool {
	return unitImpl.RunOnServer("mine", map[string]interface{}{
		"body": body,
	}).(bool)
}

// Move runs logic that moves this Unit from its current location to the
// new location specified.
func (unitImpl *UnitImpl) Move(x float64, y float64) bool {
	return unitImpl.RunOnServer("move", map[string]interface{}{
		"x": x,
		"y": y,
	}).(bool)
}

// Safe runs logic that tells you if your ship can move to that location
// from were it is without clipping the sun.
func (unitImpl *UnitImpl) Safe(x float64, y float64) bool {
	return unitImpl.RunOnServer("safe", map[string]interface{}{
		"x": x,
		"y": y,
	}).(bool)
}

// Shootdown runs logic that attacks the specified projectile.
func (unitImpl *UnitImpl) Shootdown(missile stardash.Projectile) bool {
	return unitImpl.RunOnServer("shootdown", map[string]interface{}{
		"missile": missile,
	}).(bool)
}

// Transfer runs logic that grab materials from a friendly unit. Doesn't
// use a action.
func (unitImpl *UnitImpl) Transfer(unit stardash.Unit, amount int64, material string) bool {
	return unitImpl.RunOnServer("transfer", map[string]interface{}{
		"unit": unit,
		"amount": amount,
		"material": material,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Unit.
func (unitImpl *UnitImpl) InitImplDefaults() {
	unitImpl.GameObjectImpl.InitImplDefaults()

	unitImpl.actedImpl = true
	unitImpl.dashXImpl = 0
	unitImpl.dashYImpl = 0
	unitImpl.energyImpl = 0
	unitImpl.genariumImpl = 0
	unitImpl.isBusyImpl = true
	unitImpl.jobImpl = nil
	unitImpl.legendariumImpl = 0
	unitImpl.movesImpl = 0
	unitImpl.mythiciteImpl = 0
	unitImpl.ownerImpl = nil
	unitImpl.protectorImpl = nil
	unitImpl.rariumImpl = 0
	unitImpl.shieldImpl = 0
	unitImpl.xImpl = 0
	unitImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Unit.
func (unitImpl *UnitImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := unitImpl.GameObjectImpl.DeltaMerge(
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
	case "acted":
		unitImpl.actedImpl = stardashDeltaMerge.Boolean(delta)
		return true, nil
	case "dashX":
		unitImpl.dashXImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "dashY":
		unitImpl.dashYImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "energy":
		unitImpl.energyImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "genarium":
		unitImpl.genariumImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "isBusy":
		unitImpl.isBusyImpl = stardashDeltaMerge.Boolean(delta)
		return true, nil
	case "job":
		unitImpl.jobImpl = stardashDeltaMerge.Job(delta)
		return true, nil
	case "legendarium":
		unitImpl.legendariumImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		unitImpl.movesImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "mythicite":
		unitImpl.mythiciteImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		unitImpl.ownerImpl = stardashDeltaMerge.Player(delta)
		return true, nil
	case "protector":
		unitImpl.protectorImpl = stardashDeltaMerge.Unit(delta)
		return true, nil
	case "rarium":
		unitImpl.rariumImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "shield":
		unitImpl.shieldImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "x":
		unitImpl.xImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "y":
		unitImpl.yImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

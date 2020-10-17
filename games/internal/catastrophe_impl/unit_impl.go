package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/catastrophe"
)

// UnitImpl is the struct that implements the container for Unit instances
// in Catastrophe.
type UnitImpl struct {
	GameObjectImpl

	actedImpl          bool
	energyImpl         float64
	foodImpl           int64
	jobImpl            catastrophe.Job
	materialsImpl      int64
	movementTargetImpl catastrophe.Tile
	movesImpl          int64
	ownerImpl          catastrophe.Player
	squadImpl          []catastrophe.Unit
	starvingImpl       bool
	tileImpl           catastrophe.Tile
	turnsToDieImpl     int64
}

// Acted returns whether this Unit has performed its action this turn.
func (unitImpl *UnitImpl) Acted() bool {
	return unitImpl.actedImpl
}

// Energy returns the amount of energy this Unit has (from 0.0 to 100.0).
func (unitImpl *UnitImpl) Energy() float64 {
	return unitImpl.energyImpl
}

// Food returns the amount of food this Unit is holding.
func (unitImpl *UnitImpl) Food() int64 {
	return unitImpl.foodImpl
}

// Job returns the Job this Unit was recruited to do.
func (unitImpl *UnitImpl) Job() catastrophe.Job {
	return unitImpl.jobImpl
}

// Materials returns the amount of materials this Unit is holding.
func (unitImpl *UnitImpl) Materials() int64 {
	return unitImpl.materialsImpl
}

// MovementTarget returns the tile this Unit is moving to. This only
// applies to neutral fresh humans spawned on the road. Otherwise, the tile
// this Unit is on.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) MovementTarget() catastrophe.Tile {
	return unitImpl.movementTargetImpl
}

// Moves returns how many moves this Unit has left this turn.
func (unitImpl *UnitImpl) Moves() int64 {
	return unitImpl.movesImpl
}

// Owner returns the Player that owns and can control this Unit, or nil if
// the Unit is neutral.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Owner() catastrophe.Player {
	return unitImpl.ownerImpl
}

// Squad returns the Units in the same squad as this Unit. Units in the
// same squad attack and defend together.
func (unitImpl *UnitImpl) Squad() []catastrophe.Unit {
	return unitImpl.squadImpl
}

// Starving returns whether this Unit is starving. Starving Units
// regenerate energy at half the rate they normally would while resting.
func (unitImpl *UnitImpl) Starving() bool {
	return unitImpl.starvingImpl
}

// Tile returns the Tile this Unit is on.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Tile() catastrophe.Tile {
	return unitImpl.tileImpl
}

// TurnsToDie returns the number of turns before this Unit dies. This only
// applies to neutral fresh humans created from combat. Otherwise, 0.
func (unitImpl *UnitImpl) TurnsToDie() int64 {
	return unitImpl.turnsToDieImpl
}

// Attack runs logic that attacks an adjacent Tile. Costs an action for
// each Unit in this Unit's squad. Units in the squad without an action
// don't participate in combat. Units in combat cannot move afterwards.
// Attacking structures will not give materials.
func (unitImpl *UnitImpl) Attack(tile catastrophe.Tile) bool {
	return unitImpl.RunOnServer("attack", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// ChangeJob runs logic that changes this Unit's Job. Must be at max energy
// (100) to change Jobs.
func (unitImpl *UnitImpl) ChangeJob(job string) bool {
	return unitImpl.RunOnServer("changeJob", map[string]interface{}{
		"job": job,
	}).(bool)
}

// Construct runs logic that constructs a Structure on an adjacent Tile.
func (unitImpl *UnitImpl) Construct(tile catastrophe.Tile, typeVar string) bool {
	return unitImpl.RunOnServer("construct", map[string]interface{}{
		"tile": tile,
		"type": typeVar,
	}).(bool)
}

// Convert runs logic that converts an adjacent Unit to your side.
func (unitImpl *UnitImpl) Convert(tile catastrophe.Tile) bool {
	return unitImpl.RunOnServer("convert", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Deconstruct runs logic that removes materials from an adjacent Tile's
// Structure. You cannot deconstruct friendly structures (see
// `Unit.attack`).
func (unitImpl *UnitImpl) Deconstruct(tile catastrophe.Tile) bool {
	return unitImpl.RunOnServer("deconstruct", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Drop runs logic that drops some of the given resource on or adjacent to
// the Unit's Tile. Does not count as an action.
func (unitImpl *UnitImpl) Drop(tile catastrophe.Tile, resource string, amount int64) bool {
	return unitImpl.RunOnServer("drop", map[string]interface{}{
		"tile":     tile,
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// Harvest runs logic that harvests the food on an adjacent Tile.
func (unitImpl *UnitImpl) Harvest(tile catastrophe.Tile) bool {
	return unitImpl.RunOnServer("harvest", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Move runs logic that moves this Unit from its current Tile to an
// adjacent Tile.
func (unitImpl *UnitImpl) Move(tile catastrophe.Tile) bool {
	return unitImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Pickup runs logic that picks up some materials or food on or adjacent to
// the Unit's Tile. Does not count as an action.
func (unitImpl *UnitImpl) Pickup(tile catastrophe.Tile, resource string, amount int64) bool {
	return unitImpl.RunOnServer("pickup", map[string]interface{}{
		"tile":     tile,
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// Rest runs logic that regenerates energy. Must be in range of a friendly
// shelter to rest. Costs an action. Units cannot move after resting.
func (unitImpl *UnitImpl) Rest() bool {
	return unitImpl.RunOnServer("rest", map[string]interface{}{}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Unit.
func (unitImpl *UnitImpl) InitImplDefaults() {
	unitImpl.GameObjectImpl.InitImplDefaults()

	unitImpl.actedImpl = true
	unitImpl.energyImpl = 0
	unitImpl.foodImpl = 0
	unitImpl.jobImpl = nil
	unitImpl.materialsImpl = 0
	unitImpl.movementTargetImpl = nil
	unitImpl.movesImpl = 0
	unitImpl.ownerImpl = nil
	unitImpl.squadImpl = []catastrophe.Unit{}
	unitImpl.starvingImpl = true
	unitImpl.tileImpl = nil
	unitImpl.turnsToDieImpl = 0
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

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "acted":
		unitImpl.actedImpl = catastropheDeltaMerge.Boolean(delta)
		return true, nil
	case "energy":
		unitImpl.energyImpl = catastropheDeltaMerge.Float(delta)
		return true, nil
	case "food":
		unitImpl.foodImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "job":
		unitImpl.jobImpl = catastropheDeltaMerge.Job(delta)
		return true, nil
	case "materials":
		unitImpl.materialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "movementTarget":
		unitImpl.movementTargetImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "moves":
		unitImpl.movesImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		unitImpl.ownerImpl = catastropheDeltaMerge.Player(delta)
		return true, nil
	case "squad":
		unitImpl.squadImpl = catastropheDeltaMerge.ArrayOfUnit(&unitImpl.squadImpl, delta)
		return true, nil
	case "starving":
		unitImpl.starvingImpl = catastropheDeltaMerge.Boolean(delta)
		return true, nil
	case "tile":
		unitImpl.tileImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "turnsToDie":
		unitImpl.turnsToDieImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

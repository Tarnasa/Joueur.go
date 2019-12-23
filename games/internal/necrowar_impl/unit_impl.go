package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// UnitImpl is the struct that implements the container for Unit instances
// in Necrowar.
type UnitImpl struct {
	GameObjectImpl

	actedImpl  bool
	healthImpl int64
	jobImpl    necrowar.UnitJob
	movesImpl  int64
	ownerImpl  necrowar.Player
	tileImpl   necrowar.Tile
}

// Acted returns whether or not this Unit has performed its action this
// turn (attack or build).
func (unitImpl *UnitImpl) Acted() bool {
	return unitImpl.actedImpl
}

// Health returns the remaining health of a unit.
func (unitImpl *UnitImpl) Health() int64 {
	return unitImpl.healthImpl
}

// Job returns the type of unit this is.
func (unitImpl *UnitImpl) Job() necrowar.UnitJob {
	return unitImpl.jobImpl
}

// Moves returns the number of moves this unit has left this turn.
func (unitImpl *UnitImpl) Moves() int64 {
	return unitImpl.movesImpl
}

// Owner returns the Player that owns and can control this Unit.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Owner() necrowar.Player {
	return unitImpl.ownerImpl
}

// Tile returns the Tile this Unit is on.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Tile() necrowar.Tile {
	return unitImpl.tileImpl
}

// Attack runs logic that attacks an enemy tower on an adjacent tile.
func (unitImpl *UnitImpl) Attack(tile necrowar.Tile) bool {
	return unitImpl.RunOnServer("attack", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Build runs logic that unit, if it is a worker, builds a tower on the
// tile it is on, only workers can do this.
func (unitImpl *UnitImpl) Build(title string) bool {
	return unitImpl.RunOnServer("build", map[string]interface{}{
		"title": title,
	}).(bool)
}

// Fish runs logic that stops adjacent to a river tile and begins fishing
// for mana.
func (unitImpl *UnitImpl) Fish(tile necrowar.Tile) bool {
	return unitImpl.RunOnServer("fish", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Mine runs logic that enters a mine and is put to work gathering
// resources.
func (unitImpl *UnitImpl) Mine(tile necrowar.Tile) bool {
	return unitImpl.RunOnServer("mine", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Move runs logic that moves this Unit from its current Tile to an
// adjacent Tile.
func (unitImpl *UnitImpl) Move(tile necrowar.Tile) bool {
	return unitImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Unit.
func (unitImpl *UnitImpl) InitImplDefaults() {
	unitImpl.GameObjectImpl.InitImplDefaults()

	unitImpl.actedImpl = true
	unitImpl.healthImpl = 0
	unitImpl.jobImpl = nil
	unitImpl.movesImpl = 0
	unitImpl.ownerImpl = nil
	unitImpl.tileImpl = nil
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

	necrowarDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'necrowar.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "acted":
		unitImpl.actedImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "health":
		unitImpl.healthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "job":
		unitImpl.jobImpl = necrowarDeltaMerge.UnitJob(delta)
		return true, nil
	case "moves":
		unitImpl.movesImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		unitImpl.ownerImpl = necrowarDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		unitImpl.tileImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

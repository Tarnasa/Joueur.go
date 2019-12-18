package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// UnitImpl is the struct that implements the container for Unit instances
// in Newtonian.
type UnitImpl struct {
	GameObjectImpl

	actedImpl      bool
	blueiumImpl    int64
	blueiumOreImpl int64
	healthImpl     int64
	jobImpl        newtonian.Job
	movesImpl      int64
	ownerImpl      newtonian.Player
	rediumImpl     int64
	rediumOreImpl  int64
	stunImmuneImpl int64
	stunTimeImpl   int64
	tileImpl       newtonian.Tile
}

// Acted returns whether or not this Unit has performed its action this
// turn.
func (unitImpl *UnitImpl) Acted() bool {
	return unitImpl.actedImpl
}

// Blueium returns the amount of blueium carried by this unit. (0 to job
// carry capacity - other carried items).
func (unitImpl *UnitImpl) Blueium() int64 {
	return unitImpl.blueiumImpl
}

// BlueiumOre returns the amount of blueium ore carried by this unit. (0 to
// job carry capacity - other carried items).
func (unitImpl *UnitImpl) BlueiumOre() int64 {
	return unitImpl.blueiumOreImpl
}

// Health returns the remaining health of a unit.
func (unitImpl *UnitImpl) Health() int64 {
	return unitImpl.healthImpl
}

// Job returns the Job this Unit has.
func (unitImpl *UnitImpl) Job() newtonian.Job {
	return unitImpl.jobImpl
}

// Moves returns the number of moves this unit has left this turn.
func (unitImpl *UnitImpl) Moves() int64 {
	return unitImpl.movesImpl
}

// Owner returns the Player that owns and can control this Unit.
func (unitImpl *UnitImpl) Owner() newtonian.Player {
	return unitImpl.ownerImpl
}

// Redium returns the amount of redium carried by this unit. (0 to job
// carry capacity - other carried items).
func (unitImpl *UnitImpl) Redium() int64 {
	return unitImpl.rediumImpl
}

// RediumOre returns the amount of redium ore carried by this unit. (0 to
// job carry capacity - other carried items).
func (unitImpl *UnitImpl) RediumOre() int64 {
	return unitImpl.rediumOreImpl
}

// StunImmune returns duration of stun immunity. (0 to timeImmune).
func (unitImpl *UnitImpl) StunImmune() int64 {
	return unitImpl.stunImmuneImpl
}

// StunTime returns duration the unit is stunned. (0 to the game constant
// stunTime).
func (unitImpl *UnitImpl) StunTime() int64 {
	return unitImpl.stunTimeImpl
}

// Tile returns the Tile this Unit is on.
func (unitImpl *UnitImpl) Tile() newtonian.Tile {
	return unitImpl.tileImpl
}

// Act runs logic that makes the unit do something to a machine or unit
// adjacent to its tile. Interns sabotage, physicists work. Interns stun
// physicist, physicist stuns manager, manager stuns intern.
func (unitImpl *UnitImpl) Act(tile newtonian.Tile) bool {
	return unitImpl.RunOnServer("act", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Attack runs logic that attacks a unit on an adjacent tile.
func (unitImpl *UnitImpl) Attack(tile newtonian.Tile) bool {
	return unitImpl.RunOnServer("attack", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Drop runs logic that drops materials at the units feet or adjacent tile.
func (unitImpl *UnitImpl) Drop(tile newtonian.Tile, amount int64, material string) bool {
	return unitImpl.RunOnServer("drop", map[string]interface{}{
		"tile": tile,
		"amount": amount,
		"material": material,
	}).(bool)
}

// Move runs logic that moves this Unit from its current Tile to an
// adjacent Tile.
func (unitImpl *UnitImpl) Move(tile newtonian.Tile) bool {
	return unitImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Pickup runs logic that picks up material at the units feet or adjacent
// tile.
func (unitImpl *UnitImpl) Pickup(tile newtonian.Tile, amount int64, material string) bool {
	return unitImpl.RunOnServer("pickup", map[string]interface{}{
		"tile": tile,
		"amount": amount,
		"material": material,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Unit.
func (unitImpl *UnitImpl) InitImplDefaults() {
	unitImpl.GameObjectImpl.InitImplDefaults()

	unitImpl.actedImpl = true
	unitImpl.blueiumImpl = 0
	unitImpl.blueiumOreImpl = 0
	unitImpl.healthImpl = 0
	unitImpl.jobImpl = nil
	unitImpl.movesImpl = 0
	unitImpl.ownerImpl = nil
	unitImpl.rediumImpl = 0
	unitImpl.rediumOreImpl = 0
	unitImpl.stunImmuneImpl = 0
	unitImpl.stunTimeImpl = 0
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

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "acted":
		unitImpl.actedImpl = newtonianDeltaMerge.Boolean(delta)
		return true, nil
	case "blueium":
		unitImpl.blueiumImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "blueiumOre":
		unitImpl.blueiumOreImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "health":
		unitImpl.healthImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "job":
		unitImpl.jobImpl = newtonianDeltaMerge.Job(delta)
		return true, nil
	case "moves":
		unitImpl.movesImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		unitImpl.ownerImpl = newtonianDeltaMerge.Player(delta)
		return true, nil
	case "redium":
		unitImpl.rediumImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "rediumOre":
		unitImpl.rediumOreImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "stunImmune":
		unitImpl.stunImmuneImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "stunTime":
		unitImpl.stunTimeImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "tile":
		unitImpl.tileImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

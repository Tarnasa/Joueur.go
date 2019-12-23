package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stumped"
)

// BeaverImpl is the struct that implements the container for Beaver
// instances in Stumped.
type BeaverImpl struct {
	GameObjectImpl

	actionsImpl         int64
	branchesImpl        int64
	foodImpl            int64
	healthImpl          int64
	jobImpl             stumped.Job
	movesImpl           int64
	ownerImpl           stumped.Player
	recruitedImpl       bool
	tileImpl            stumped.Tile
	turnsDistractedImpl int64
}

// Actions returns the number of actions remaining for the Beaver this
// turn.
func (beaverImpl *BeaverImpl) Actions() int64 {
	return beaverImpl.actionsImpl
}

// Branches returns the amount of branches this Beaver is holding.
func (beaverImpl *BeaverImpl) Branches() int64 {
	return beaverImpl.branchesImpl
}

// Food returns the amount of food this Beaver is holding.
func (beaverImpl *BeaverImpl) Food() int64 {
	return beaverImpl.foodImpl
}

// Health returns how much health this Beaver has left.
func (beaverImpl *BeaverImpl) Health() int64 {
	return beaverImpl.healthImpl
}

// Job returns the Job this Beaver was recruited to do.
func (beaverImpl *BeaverImpl) Job() stumped.Job {
	return beaverImpl.jobImpl
}

// Moves returns how many moves this Beaver has left this turn.
func (beaverImpl *BeaverImpl) Moves() int64 {
	return beaverImpl.movesImpl
}

// Owner returns the Player that owns and can control this Beaver.
func (beaverImpl *BeaverImpl) Owner() stumped.Player {
	return beaverImpl.ownerImpl
}

// Recruited returns true if the Beaver has finished being recruited and
// can do things, False otherwise.
func (beaverImpl *BeaverImpl) Recruited() bool {
	return beaverImpl.recruitedImpl
}

// Tile returns the Tile this Beaver is on.
//
// Value can be returned as a nil pointer.
func (beaverImpl *BeaverImpl) Tile() stumped.Tile {
	return beaverImpl.tileImpl
}

// TurnsDistracted returns number of turns this Beaver is distracted for (0
// means not distracted).
func (beaverImpl *BeaverImpl) TurnsDistracted() int64 {
	return beaverImpl.turnsDistractedImpl
}

// Attack runs logic that attacks another adjacent beaver.
func (beaverImpl *BeaverImpl) Attack(beaver stumped.Beaver) bool {
	return beaverImpl.RunOnServer("attack", map[string]interface{}{
		"beaver": beaver,
	}).(bool)
}

// BuildLodge runs logic that builds a lodge on the Beavers current Tile.
func (beaverImpl *BeaverImpl) BuildLodge() bool {
	return beaverImpl.RunOnServer("buildLodge", map[string]interface{}{}).(bool)
}

// Drop runs logic that drops some of the given resource on the beaver's
// Tile.
func (beaverImpl *BeaverImpl) Drop(tile stumped.Tile, resource string, amount int64) bool {
	return beaverImpl.RunOnServer("drop", map[string]interface{}{
		"tile":     tile,
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// Harvest runs logic that harvests the branches or food from a Spawner on
// an adjacent Tile.
func (beaverImpl *BeaverImpl) Harvest(spawner stumped.Spawner) bool {
	return beaverImpl.RunOnServer("harvest", map[string]interface{}{
		"spawner": spawner,
	}).(bool)
}

// Move runs logic that moves this Beaver from its current Tile to an
// adjacent Tile.
func (beaverImpl *BeaverImpl) Move(tile stumped.Tile) bool {
	return beaverImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Pickup runs logic that picks up some branches or food on the beaver's
// tile.
func (beaverImpl *BeaverImpl) Pickup(tile stumped.Tile, resource string, amount int64) bool {
	return beaverImpl.RunOnServer("pickup", map[string]interface{}{
		"tile":     tile,
		"resource": resource,
		"amount":   amount,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Beaver.
func (beaverImpl *BeaverImpl) InitImplDefaults() {
	beaverImpl.GameObjectImpl.InitImplDefaults()

	beaverImpl.actionsImpl = 0
	beaverImpl.branchesImpl = 0
	beaverImpl.foodImpl = 0
	beaverImpl.healthImpl = 0
	beaverImpl.jobImpl = nil
	beaverImpl.movesImpl = 0
	beaverImpl.ownerImpl = nil
	beaverImpl.recruitedImpl = true
	beaverImpl.tileImpl = nil
	beaverImpl.turnsDistractedImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Beaver.
func (beaverImpl *BeaverImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := beaverImpl.GameObjectImpl.DeltaMerge(
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
		beaverImpl.actionsImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "branches":
		beaverImpl.branchesImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "food":
		beaverImpl.foodImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "health":
		beaverImpl.healthImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "job":
		beaverImpl.jobImpl = stumpedDeltaMerge.Job(delta)
		return true, nil
	case "moves":
		beaverImpl.movesImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		beaverImpl.ownerImpl = stumpedDeltaMerge.Player(delta)
		return true, nil
	case "recruited":
		beaverImpl.recruitedImpl = stumpedDeltaMerge.Boolean(delta)
		return true, nil
	case "tile":
		beaverImpl.tileImpl = stumpedDeltaMerge.Tile(delta)
		return true, nil
	case "turnsDistracted":
		beaverImpl.turnsDistractedImpl = stumpedDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

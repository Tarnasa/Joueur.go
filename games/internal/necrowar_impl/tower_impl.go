package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/necrowar"
)

// TowerImpl is the struct that implements the container for Tower
// instances in Necrowar.
type TowerImpl struct {
	GameObjectImpl

	attackedImpl bool
	cooldownImpl int64
	healthImpl   int64
	jobImpl      necrowar.TowerJob
	ownerImpl    necrowar.Player
	tileImpl     necrowar.Tile
}

// Attacked returns whether this tower has attacked this turn or not.
func (towerImpl *TowerImpl) Attacked() bool {
	return towerImpl.attackedImpl
}

// Cooldown returns how many turns are left before it can fire again.
func (towerImpl *TowerImpl) Cooldown() int64 {
	return towerImpl.cooldownImpl
}

// Health returns how much remaining health this tower has.
func (towerImpl *TowerImpl) Health() int64 {
	return towerImpl.healthImpl
}

// Job returns what type of tower this is (it's job).
func (towerImpl *TowerImpl) Job() necrowar.TowerJob {
	return towerImpl.jobImpl
}

// Owner returns the player that built / owns this tower.
func (towerImpl *TowerImpl) Owner() necrowar.Player {
	return towerImpl.ownerImpl
}

// Tile returns the Tile this Tower is on.
func (towerImpl *TowerImpl) Tile() necrowar.Tile {
	return towerImpl.tileImpl
}

// Attack runs logic that attacks an enemy unit on an tile within it's
// range.
func (towerImpl *TowerImpl) Attack(tile necrowar.Tile) bool {
	return towerImpl.RunOnServer("attack", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Tower.
func (towerImpl *TowerImpl) InitImplDefaults() {
	towerImpl.GameObjectImpl.InitImplDefaults()

	towerImpl.attackedImpl = true
	towerImpl.cooldownImpl = 0
	towerImpl.healthImpl = 0
	towerImpl.jobImpl = nil
	towerImpl.ownerImpl = nil
	towerImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Tower.
func (towerImpl *TowerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := towerImpl.GameObjectImpl.DeltaMerge(
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
	case "attacked":
		towerImpl.attackedImpl = necrowarDeltaMerge.Boolean(delta)
		return true, nil
	case "cooldown":
		towerImpl.cooldownImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "health":
		towerImpl.healthImpl = necrowarDeltaMerge.Int(delta)
		return true, nil
	case "job":
		towerImpl.jobImpl = necrowarDeltaMerge.TowerJob(delta)
		return true, nil
	case "owner":
		towerImpl.ownerImpl = necrowarDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		towerImpl.tileImpl = necrowarDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// SpiderImpl is the struct that implements the container for Spider
// instances in Spiders.
type SpiderImpl struct {
	GameObjectImpl

	isDeadImpl bool
	nestImpl   spiders.Nest
	ownerImpl  spiders.Player
}

// IsDead returns if this Spider is dead and has been removed from the
// game.
func (spiderImpl *SpiderImpl) IsDead() bool {
	return spiderImpl.isDeadImpl
}

// Nest returns the Nest that this Spider is currently on. Nil when moving
// on a Web.
//
// Value can be returned as a nil pointer.
func (spiderImpl *SpiderImpl) Nest() spiders.Nest {
	return spiderImpl.nestImpl
}

// Owner returns the Player that owns this Spider, and can command it.
func (spiderImpl *SpiderImpl) Owner() spiders.Player {
	return spiderImpl.ownerImpl
}

// InitImplDefaults initializes safe defaults for all fields in Spider.
func (spiderImpl *SpiderImpl) InitImplDefaults() {
	spiderImpl.GameObjectImpl.InitImplDefaults()

	spiderImpl.isDeadImpl = true
	spiderImpl.nestImpl = nil
	spiderImpl.ownerImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Spider.
func (spiderImpl *SpiderImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := spiderImpl.GameObjectImpl.DeltaMerge(
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
	case "isDead":
		spiderImpl.isDeadImpl = spidersDeltaMerge.Boolean(delta)
		return true, nil
	case "nest":
		spiderImpl.nestImpl = spidersDeltaMerge.Nest(delta)
		return true, nil
	case "owner":
		spiderImpl.ownerImpl = spidersDeltaMerge.Player(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

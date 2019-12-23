package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// NestImpl is the struct that implements the container for Nest instances
// in Spiders.
type NestImpl struct {
	GameObjectImpl

	controllingPlayerImpl spiders.Player
	spidersImpl           []spiders.Spider
	websImpl              []spiders.Web
	xImpl                 int64
	yImpl                 int64
}

// ControllingPlayer returns the Player that 'controls' this Nest as they
// have the most Spiders on this nest.
//
// Value can be returned as a nil pointer.
func (nestImpl *NestImpl) ControllingPlayer() spiders.Player {
	return nestImpl.controllingPlayerImpl
}

// Spiders returns all the Spiders currently located on this Nest.
func (nestImpl *NestImpl) Spiders() []spiders.Spider {
	return nestImpl.spidersImpl
}

// Webs returns webs that connect to this Nest.
func (nestImpl *NestImpl) Webs() []spiders.Web {
	return nestImpl.websImpl
}

// X returns the X coordinate of the Nest. Used for distance calculations.
func (nestImpl *NestImpl) X() int64 {
	return nestImpl.xImpl
}

// Y returns the Y coordinate of the Nest. Used for distance calculations.
func (nestImpl *NestImpl) Y() int64 {
	return nestImpl.yImpl
}

// InitImplDefaults initializes safe defaults for all fields in Nest.
func (nestImpl *NestImpl) InitImplDefaults() {
	nestImpl.GameObjectImpl.InitImplDefaults()

	nestImpl.controllingPlayerImpl = nil
	nestImpl.spidersImpl = []spiders.Spider{}
	nestImpl.websImpl = []spiders.Web{}
	nestImpl.xImpl = 0
	nestImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Nest.
func (nestImpl *NestImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := nestImpl.GameObjectImpl.DeltaMerge(
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
	case "controllingPlayer":
		nestImpl.controllingPlayerImpl = spidersDeltaMerge.Player(delta)
		return true, nil
	case "spiders":
		nestImpl.spidersImpl = spidersDeltaMerge.ArrayOfSpider(&nestImpl.spidersImpl, delta)
		return true, nil
	case "webs":
		nestImpl.websImpl = spidersDeltaMerge.ArrayOfWeb(&nestImpl.websImpl, delta)
		return true, nil
	case "x":
		nestImpl.xImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "y":
		nestImpl.yImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

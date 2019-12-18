package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// WebImpl is the struct that implements the container for Web instances in
// Spiders.
type WebImpl struct {
	GameObjectImpl

	lengthImpl      float64
	loadImpl        int64
	nestAImpl       spiders.Nest
	nestBImpl       spiders.Nest
	spiderlingsImpl []spiders.Spiderling
	strengthImpl    int64
}

// Length returns how long this Web is, i.e., the distance between its
// nestA and nestB.
func (webImpl *WebImpl) Length() float64 {
	return webImpl.lengthImpl
}

// Load returns how much weight this Web currently has on it, which is the
// sum of all its Spiderlings weight.
func (webImpl *WebImpl) Load() int64 {
	return webImpl.loadImpl
}

// NestA returns the first Nest this Web is connected to.
func (webImpl *WebImpl) NestA() spiders.Nest {
	return webImpl.nestAImpl
}

// NestB returns the second Nest this Web is connected to.
func (webImpl *WebImpl) NestB() spiders.Nest {
	return webImpl.nestBImpl
}

// Spiderlings returns all the Spiderlings currently moving along this Web.
func (webImpl *WebImpl) Spiderlings() []spiders.Spiderling {
	return webImpl.spiderlingsImpl
}

// Strength returns how much weight this Web can take before snapping and
// destroying itself and all the Spiders on it.
func (webImpl *WebImpl) Strength() int64 {
	return webImpl.strengthImpl
}

// InitImplDefaults initializes safe defaults for all fields in Web.
func (webImpl *WebImpl) InitImplDefaults() {
	webImpl.GameObjectImpl.InitImplDefaults()

	webImpl.lengthImpl = 0
	webImpl.loadImpl = 0
	webImpl.nestAImpl = nil
	webImpl.nestBImpl = nil
	webImpl.spiderlingsImpl = []spiders.Spiderling{}
	webImpl.strengthImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Web.
func (webImpl *WebImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := webImpl.GameObjectImpl.DeltaMerge(
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
	case "length":
		webImpl.lengthImpl = spidersDeltaMerge.Float(delta)
		return true, nil
	case "load":
		webImpl.loadImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "nestA":
		webImpl.nestAImpl = spidersDeltaMerge.Nest(delta)
		return true, nil
	case "nestB":
		webImpl.nestBImpl = spidersDeltaMerge.Nest(delta)
		return true, nil
	case "spiderlings":
		webImpl.spiderlingsImpl = spidersDeltaMerge.ArrayOfSpiderling(&webImpl.spiderlingsImpl, delta)
		return true, nil
	case "strength":
		webImpl.strengthImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

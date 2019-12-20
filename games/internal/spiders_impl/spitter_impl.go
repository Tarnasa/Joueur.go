package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// SpitterImpl is the struct that implements the container for Spitter
// instances in Spiders.
type SpitterImpl struct {
	SpiderlingImpl

	spittingWebToNestImpl spiders.Nest
}

// SpittingWebToNest returns the Nest that this Spitter is creating a Web
// to spit at, thus connecting them. Nil if not spitting.
//
// Value can be returned as a nil pointer.
func (spitterImpl *SpitterImpl) SpittingWebToNest() spiders.Nest {
	return spitterImpl.spittingWebToNestImpl
}

// Spit runs logic that creates and spits a new Web from the Nest the
// Spitter is on to another Nest, connecting them.
func (spitterImpl *SpitterImpl) Spit(nest spiders.Nest) bool {
	return spitterImpl.RunOnServer("spit", map[string]interface{}{
		"nest": nest,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Spitter.
func (spitterImpl *SpitterImpl) InitImplDefaults() {
	spitterImpl.SpiderlingImpl.InitImplDefaults()

	spitterImpl.spittingWebToNestImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Spitter.
func (spitterImpl *SpitterImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := spitterImpl.SpiderlingImpl.DeltaMerge(
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
	case "spittingWebToNest":
		spitterImpl.spittingWebToNestImpl = spidersDeltaMerge.Nest(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

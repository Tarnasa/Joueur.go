package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// CutterImpl is the struct that implements the container for Cutter
// instances in Spiders.
type CutterImpl struct {
	SpiderlingImpl

	cuttingWebImpl spiders.Web
}

// CuttingWeb returns the Web that this Cutter is trying to cut. Nil if not
// cutting.
func (cutterImpl *CutterImpl) CuttingWeb() spiders.Web {
	return cutterImpl.cuttingWebImpl
}

// Cut runs logic that cuts a web, destroying it, and any Spiderlings on
// it.
func (cutterImpl *CutterImpl) Cut(web spiders.Web) bool {
	return cutterImpl.RunOnServer("cut", map[string]interface{}{
		"web": web,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Cutter.
func (cutterImpl *CutterImpl) InitImplDefaults() {
	cutterImpl.SpiderlingImpl.InitImplDefaults()

	cutterImpl.cuttingWebImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Cutter.
func (cutterImpl *CutterImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := cutterImpl.SpiderlingImpl.DeltaMerge(
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
	case "cuttingWeb":
		cutterImpl.cuttingWebImpl = spidersDeltaMerge.Web(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

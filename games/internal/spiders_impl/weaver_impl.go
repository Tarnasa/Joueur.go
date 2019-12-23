package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// WeaverImpl is the struct that implements the container for Weaver
// instances in Spiders.
type WeaverImpl struct {
	SpiderlingImpl

	strengtheningWebImpl spiders.Web
	weakeningWebImpl     spiders.Web
}

// StrengtheningWeb returns the Web that this Weaver is strengthening. Nil
// if not strengthening.
//
// Value can be returned as a nil pointer.
func (weaverImpl *WeaverImpl) StrengtheningWeb() spiders.Web {
	return weaverImpl.strengtheningWebImpl
}

// WeakeningWeb returns the Web that this Weaver is weakening. Nil if not
// weakening.
//
// Value can be returned as a nil pointer.
func (weaverImpl *WeaverImpl) WeakeningWeb() spiders.Web {
	return weaverImpl.weakeningWebImpl
}

// Strengthen runs logic that weaves more silk into an existing Web to
// strengthen it.
func (weaverImpl *WeaverImpl) Strengthen(web spiders.Web) bool {
	return weaverImpl.RunOnServer("strengthen", map[string]interface{}{
		"web": web,
	}).(bool)
}

// Weaken runs logic that weaves more silk into an existing Web to
// strengthen it.
func (weaverImpl *WeaverImpl) Weaken(web spiders.Web) bool {
	return weaverImpl.RunOnServer("weaken", map[string]interface{}{
		"web": web,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Weaver.
func (weaverImpl *WeaverImpl) InitImplDefaults() {
	weaverImpl.SpiderlingImpl.InitImplDefaults()

	weaverImpl.strengtheningWebImpl = nil
	weaverImpl.weakeningWebImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Weaver.
func (weaverImpl *WeaverImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := weaverImpl.SpiderlingImpl.DeltaMerge(
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
	case "strengtheningWeb":
		weaverImpl.strengtheningWebImpl = spidersDeltaMerge.Web(delta)
		return true, nil
	case "weakeningWeb":
		weaverImpl.weakeningWebImpl = spidersDeltaMerge.Web(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

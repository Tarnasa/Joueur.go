package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/spiders"
)

// BroodMotherImpl is the struct that implements the container for
// BroodMother instances in Spiders.
type BroodMotherImpl struct {
	SpiderImpl

	eggsImpl   int64
	healthImpl int64
}

// Eggs returns how many eggs the BroodMother has to spawn Spiderlings this
// turn.
func (broodMotherImpl *BroodMotherImpl) Eggs() int64 {
	return broodMotherImpl.eggsImpl
}

// Health returns how much health this BroodMother has left. When it
// reaches 0, she dies and her owner loses.
func (broodMotherImpl *BroodMotherImpl) Health() int64 {
	return broodMotherImpl.healthImpl
}

// Consume runs logic that consumes a Spiderling of the same owner to
// regain some eggs to spawn more Spiderlings.
func (broodMotherImpl *BroodMotherImpl) Consume(spiderling spiders.Spiderling) bool {
	return broodMotherImpl.RunOnServer("consume", map[string]interface{}{
		"spiderling": spiderling,
	}).(bool)
}

// Spawn runs logic that spawns a new Spiderling on the same Nest as this
// BroodMother, consuming an egg.
func (broodMotherImpl *BroodMotherImpl) Spawn(spiderlingType string) spiders.Spiderling {
	return broodMotherImpl.RunOnServer("spawn", map[string]interface{}{
		"spiderlingType": spiderlingType,
	}).(spiders.Spiderling)
}

// InitImplDefaults initializes safe defaults for all fields in BroodMother.
func (broodMotherImpl *BroodMotherImpl) InitImplDefaults() {
	broodMotherImpl.SpiderImpl.InitImplDefaults()

	broodMotherImpl.eggsImpl = 0
	broodMotherImpl.healthImpl = 0
}

// DeltaMerge merges the delta for a given attribute in BroodMother.
func (broodMotherImpl *BroodMotherImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := broodMotherImpl.SpiderImpl.DeltaMerge(
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
	case "eggs":
		broodMotherImpl.eggsImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	case "health":
		broodMotherImpl.healthImpl = spidersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// PoliceDepartmentImpl is the struct that implements the container for
// PoliceDepartment instances in Anarchy.
type PoliceDepartmentImpl struct {
	BuildingImpl

}

// Raid runs logic that bribe the police to raid a Warehouse, dealing
// damage equal based on the Warehouse's current exposure, and then
// resetting it to 0.
func (policeDepartmentImpl *PoliceDepartmentImpl) Raid(warehouse anarchy.Warehouse) int64 {
	return policeDepartmentImpl.RunOnServer("raid", map[string]interface{}{
		"warehouse": warehouse,
	}).(int64)
}

// InitImplDefaults initializes safe defaults for all fields in PoliceDepartment.
func (policeDepartmentImpl *PoliceDepartmentImpl) InitImplDefaults() {
	policeDepartmentImpl.BuildingImpl.InitImplDefaults()

}

// DeltaMerge merges the delta for a given attribute in PoliceDepartment.
func (policeDepartmentImpl *PoliceDepartmentImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := policeDepartmentImpl.BuildingImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	_, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'anarchy.impl.DeltaMerge'",
		)
	}

	return false, nil // no errors in delta merging
}

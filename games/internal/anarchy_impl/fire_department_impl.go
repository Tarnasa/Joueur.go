package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// FireDepartmentImpl is the struct that implements the container for
// FireDepartment instances in Anarchy.
type FireDepartmentImpl struct {
	BuildingImpl

	fireExtinguishedImpl int64
}

// FireExtinguished returns the amount of fire removed from a building when
// bribed to extinguish a building.
func (fireDepartmentImpl *FireDepartmentImpl) FireExtinguished() int64 {
	return fireDepartmentImpl.fireExtinguishedImpl
}

// Extinguish runs logic that bribes this FireDepartment to extinguish the
// some of the fire in a building.
func (fireDepartmentImpl *FireDepartmentImpl) Extinguish(building anarchy.Building) bool {
	return fireDepartmentImpl.RunOnServer("extinguish", map[string]interface{}{
		"building": building,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in FireDepartment.
func (fireDepartmentImpl *FireDepartmentImpl) InitImplDefaults() {
	fireDepartmentImpl.BuildingImpl.InitImplDefaults()

	fireDepartmentImpl.fireExtinguishedImpl = 0
}

// DeltaMerge merges the delta for a given attribute in FireDepartment.
func (fireDepartmentImpl *FireDepartmentImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := fireDepartmentImpl.BuildingImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	anarchyDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'anarchy.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "fireExtinguished":
		fireDepartmentImpl.fireExtinguishedImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

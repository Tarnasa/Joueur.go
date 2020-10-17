package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// WarehouseImpl is the struct that implements the container for Warehouse
// instances in Anarchy.
type WarehouseImpl struct {
	BuildingImpl

	exposureImpl  int64
	fireAddedImpl int64
}

// Exposure returns how exposed the anarchists in this warehouse are to
// PoliceDepartments. Raises when bribed to ignite buildings, and drops
// each turn if not bribed.
func (warehouseImpl *WarehouseImpl) Exposure() int64 {
	return warehouseImpl.exposureImpl
}

// FireAdded returns the amount of fire added to buildings when bribed to
// ignite a building. Headquarters add more fire than normal Warehouses.
func (warehouseImpl *WarehouseImpl) FireAdded() int64 {
	return warehouseImpl.fireAddedImpl
}

// Ignite runs logic that bribes the Warehouse to light a Building on fire.
// This adds this building's fireAdded to their fire, and then this
// building's exposure is increased based on the Manhattan distance between
// the two buildings.
func (warehouseImpl *WarehouseImpl) Ignite(building anarchy.Building) int64 {
	return warehouseImpl.RunOnServer("ignite", map[string]interface{}{
		"building": building,
	}).(int64)
}

// InitImplDefaults initializes safe defaults for all fields in Warehouse.
func (warehouseImpl *WarehouseImpl) InitImplDefaults() {
	warehouseImpl.BuildingImpl.InitImplDefaults()

	warehouseImpl.exposureImpl = 0
	warehouseImpl.fireAddedImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Warehouse.
func (warehouseImpl *WarehouseImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := warehouseImpl.BuildingImpl.DeltaMerge(
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
	case "exposure":
		warehouseImpl.exposureImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "fireAdded":
		warehouseImpl.fireAddedImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

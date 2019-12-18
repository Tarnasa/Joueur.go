package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// BuildingImpl is the struct that implements the container for Building
// instances in Anarchy.
type BuildingImpl struct {
	GameObjectImpl

	bribedImpl         bool
	buildingEastImpl   anarchy.Building
	buildingNorthImpl  anarchy.Building
	buildingSouthImpl  anarchy.Building
	buildingWestImpl   anarchy.Building
	fireImpl           int64
	healthImpl         int64
	isHeadquartersImpl bool
	ownerImpl          anarchy.Player
	xImpl              int64
	yImpl              int64
}

// Bribed returns when true this building has already been bribed this turn
// and cannot be bribed again this turn.
func (buildingImpl *BuildingImpl) Bribed() bool {
	return buildingImpl.bribedImpl
}

// BuildingEast returns the Building directly to the east of this building,
// or nil if not present.
func (buildingImpl *BuildingImpl) BuildingEast() anarchy.Building {
	return buildingImpl.buildingEastImpl
}

// BuildingNorth returns the Building directly to the north of this
// building, or nil if not present.
func (buildingImpl *BuildingImpl) BuildingNorth() anarchy.Building {
	return buildingImpl.buildingNorthImpl
}

// BuildingSouth returns the Building directly to the south of this
// building, or nil if not present.
func (buildingImpl *BuildingImpl) BuildingSouth() anarchy.Building {
	return buildingImpl.buildingSouthImpl
}

// BuildingWest returns the Building directly to the west of this building,
// or nil if not present.
func (buildingImpl *BuildingImpl) BuildingWest() anarchy.Building {
	return buildingImpl.buildingWestImpl
}

// Fire returns how much fire is currently burning the building, and thus
// how much damage it will take at the end of its owner's turn. 0 means no
// fire.
func (buildingImpl *BuildingImpl) Fire() int64 {
	return buildingImpl.fireImpl
}

// Health returns how much health this building currently has. When this
// reaches 0 the Building has been burned down.
func (buildingImpl *BuildingImpl) Health() int64 {
	return buildingImpl.healthImpl
}

// IsHeadquarters returns true if this is the Headquarters of the owning
// player, false otherwise. Burning this down wins the game for the other
// Player.
func (buildingImpl *BuildingImpl) IsHeadquarters() bool {
	return buildingImpl.isHeadquartersImpl
}

// Owner returns the player that owns this building. If it burns down
// (health reaches 0) that player gets an additional bribe(s).
func (buildingImpl *BuildingImpl) Owner() anarchy.Player {
	return buildingImpl.ownerImpl
}

// X returns the location of the Building along the x-axis.
func (buildingImpl *BuildingImpl) X() int64 {
	return buildingImpl.xImpl
}

// Y returns the location of the Building along the y-axis.
func (buildingImpl *BuildingImpl) Y() int64 {
	return buildingImpl.yImpl
}

// InitImplDefaults initializes safe defaults for all fields in Building.
func (buildingImpl *BuildingImpl) InitImplDefaults() {
	buildingImpl.GameObjectImpl.InitImplDefaults()

	buildingImpl.bribedImpl = true
	buildingImpl.buildingEastImpl = nil
	buildingImpl.buildingNorthImpl = nil
	buildingImpl.buildingSouthImpl = nil
	buildingImpl.buildingWestImpl = nil
	buildingImpl.fireImpl = 0
	buildingImpl.healthImpl = 0
	buildingImpl.isHeadquartersImpl = true
	buildingImpl.ownerImpl = nil
	buildingImpl.xImpl = 0
	buildingImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Building.
func (buildingImpl *BuildingImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := buildingImpl.GameObjectImpl.DeltaMerge(
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
	case "bribed":
		buildingImpl.bribedImpl = anarchyDeltaMerge.Boolean(delta)
		return true, nil
	case "buildingEast":
		buildingImpl.buildingEastImpl = anarchyDeltaMerge.Building(delta)
		return true, nil
	case "buildingNorth":
		buildingImpl.buildingNorthImpl = anarchyDeltaMerge.Building(delta)
		return true, nil
	case "buildingSouth":
		buildingImpl.buildingSouthImpl = anarchyDeltaMerge.Building(delta)
		return true, nil
	case "buildingWest":
		buildingImpl.buildingWestImpl = anarchyDeltaMerge.Building(delta)
		return true, nil
	case "fire":
		buildingImpl.fireImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "health":
		buildingImpl.healthImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "isHeadquarters":
		buildingImpl.isHeadquartersImpl = anarchyDeltaMerge.Boolean(delta)
		return true, nil
	case "owner":
		buildingImpl.ownerImpl = anarchyDeltaMerge.Player(delta)
		return true, nil
	case "x":
		buildingImpl.xImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "y":
		buildingImpl.yImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

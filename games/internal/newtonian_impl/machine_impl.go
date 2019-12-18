package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// MachineImpl is the struct that implements the container for Machine
// instances in Newtonian.
type MachineImpl struct {
	GameObjectImpl

	oreTypeImpl      string
	refineInputImpl  int64
	refineOutputImpl int64
	refineTimeImpl   int64
	tileImpl         newtonian.Tile
	workedImpl       int64
}

// OreType returns what type of ore the machine takes it. Also determines
// the type of material it outputs. (redium or blueium).
// Literal Values: "redium" or "blueium"
func (machineImpl *MachineImpl) OreType() string {
	return machineImpl.oreTypeImpl
}

// RefineInput returns the amount of ore that needs to be inputted into the
// machine for it to be worked.
func (machineImpl *MachineImpl) RefineInput() int64 {
	return machineImpl.refineInputImpl
}

// RefineOutput returns the amount of refined ore that is returned after
// the machine has been fully worked.
func (machineImpl *MachineImpl) RefineOutput() int64 {
	return machineImpl.refineOutputImpl
}

// RefineTime returns the number of times this machine needs to be worked
// to refine ore.
func (machineImpl *MachineImpl) RefineTime() int64 {
	return machineImpl.refineTimeImpl
}

// Tile returns the Tile this Machine is on.
func (machineImpl *MachineImpl) Tile() newtonian.Tile {
	return machineImpl.tileImpl
}

// Worked returns tracks how many times this machine has been worked. (0 to
// refineTime).
func (machineImpl *MachineImpl) Worked() int64 {
	return machineImpl.workedImpl
}

// InitImplDefaults initializes safe defaults for all fields in Machine.
func (machineImpl *MachineImpl) InitImplDefaults() {
	machineImpl.GameObjectImpl.InitImplDefaults()

	machineImpl.oreTypeImpl = ""
	machineImpl.refineInputImpl = 0
	machineImpl.refineOutputImpl = 0
	machineImpl.refineTimeImpl = 0
	machineImpl.tileImpl = nil
	machineImpl.workedImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Machine.
func (machineImpl *MachineImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := machineImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "oreType":
		machineImpl.oreTypeImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "refineInput":
		machineImpl.refineInputImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "refineOutput":
		machineImpl.refineOutputImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "refineTime":
		machineImpl.refineTimeImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "tile":
		machineImpl.tileImpl = newtonianDeltaMerge.Tile(delta)
		return true, nil
	case "worked":
		machineImpl.workedImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

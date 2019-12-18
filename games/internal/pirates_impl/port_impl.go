package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// PortImpl is the struct that implements the container for Port instances
// in Pirates.
type PortImpl struct {
	GameObjectImpl

	goldImpl       int64
	investmentImpl int64
	ownerImpl      pirates.Player
	tileImpl       pirates.Tile
}

// Gold returns for players, how much more gold this Port can spend this
// turn. For merchants, how much gold this Port has accumulated (it will
// spawn a ship when the Port can afford one).
func (portImpl *PortImpl) Gold() int64 {
	return portImpl.goldImpl
}

// Investment returns (Merchants only) How much gold was invested into this
// Port. Investment determines the strength and value of the next ship.
func (portImpl *PortImpl) Investment() int64 {
	return portImpl.investmentImpl
}

// Owner returns the owner of this Port, or nil if owned by merchants.
func (portImpl *PortImpl) Owner() pirates.Player {
	return portImpl.ownerImpl
}

// Tile returns the Tile this Port is on.
func (portImpl *PortImpl) Tile() pirates.Tile {
	return portImpl.tileImpl
}

// Spawn runs logic that spawn a Unit on this port.
func (portImpl *PortImpl) Spawn(typeVar string) bool {
	return portImpl.RunOnServer("spawn", map[string]interface{}{
		"type": typeVar,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Port.
func (portImpl *PortImpl) InitImplDefaults() {
	portImpl.GameObjectImpl.InitImplDefaults()

	portImpl.goldImpl = 0
	portImpl.investmentImpl = 0
	portImpl.ownerImpl = nil
	portImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Port.
func (portImpl *PortImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := portImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	piratesDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'pirates.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "gold":
		portImpl.goldImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "investment":
		portImpl.investmentImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		portImpl.ownerImpl = piratesDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		portImpl.tileImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

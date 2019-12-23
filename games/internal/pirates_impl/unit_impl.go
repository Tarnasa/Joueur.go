package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/pirates"
)

// UnitImpl is the struct that implements the container for Unit instances
// in Pirates.
type UnitImpl struct {
	GameObjectImpl

	actedImpl      bool
	crewImpl       int64
	crewHealthImpl int64
	goldImpl       int64
	movesImpl      int64
	ownerImpl      pirates.Player
	pathImpl       []pirates.Tile
	shipHealthImpl int64
	stunTurnsImpl  int64
	targetPortImpl pirates.Port
	tileImpl       pirates.Tile
}

// Acted returns whether this Unit has performed its action this turn.
func (unitImpl *UnitImpl) Acted() bool {
	return unitImpl.actedImpl
}

// Crew returns how many crew are on this Tile. This number will always be
// <= crewHealth.
func (unitImpl *UnitImpl) Crew() int64 {
	return unitImpl.crewImpl
}

// CrewHealth returns how much total health the crew on this Tile have.
func (unitImpl *UnitImpl) CrewHealth() int64 {
	return unitImpl.crewHealthImpl
}

// Gold returns how much gold this Unit is carrying.
func (unitImpl *UnitImpl) Gold() int64 {
	return unitImpl.goldImpl
}

// Moves returns how many more times this Unit may move this turn.
func (unitImpl *UnitImpl) Moves() int64 {
	return unitImpl.movesImpl
}

// Owner returns the Player that owns and can control this Unit, or nil if
// the Unit is neutral.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Owner() pirates.Player {
	return unitImpl.ownerImpl
}

// Path returns (Merchants only) The path this Unit will follow. The first
// element is the Tile this Unit will move to next.
func (unitImpl *UnitImpl) Path() []pirates.Tile {
	return unitImpl.pathImpl
}

// ShipHealth returns if a ship is on this Tile, how much health it has
// remaining. 0 for no ship.
func (unitImpl *UnitImpl) ShipHealth() int64 {
	return unitImpl.shipHealthImpl
}

// StunTurns returns (Merchants only) The number of turns this merchant
// ship won't be able to move. They will still attack. Merchant ships are
// stunned when they're attacked.
func (unitImpl *UnitImpl) StunTurns() int64 {
	return unitImpl.stunTurnsImpl
}

// TargetPort returns (Merchants only) The Port this Unit is moving to.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) TargetPort() pirates.Port {
	return unitImpl.targetPortImpl
}

// Tile returns the Tile this Unit is on.
//
// Value can be returned as a nil pointer.
func (unitImpl *UnitImpl) Tile() pirates.Tile {
	return unitImpl.tileImpl
}

// Attack runs logic that attacks either the 'crew' or 'ship' on a Tile in
// range.
func (unitImpl *UnitImpl) Attack(tile pirates.Tile, target string) bool {
	return unitImpl.RunOnServer("attack", map[string]interface{}{
		"tile": tile,
		"target": target,
	}).(bool)
}

// Bury runs logic that buries gold on this Unit's Tile. Gold must be a
// certain distance away for it to get interest (Game.minInterestDistance).
func (unitImpl *UnitImpl) Bury(amount int64) bool {
	return unitImpl.RunOnServer("bury", map[string]interface{}{
		"amount": amount,
	}).(bool)
}

// Deposit runs logic that puts gold into an adjacent Port. If that Port is
// the Player's port, the gold is added to that Player. If that Port is
// owned by merchants, it adds to that Port's investment.
func (unitImpl *UnitImpl) Deposit(amount int64) bool {
	return unitImpl.RunOnServer("deposit", map[string]interface{}{
		"amount": amount,
	}).(bool)
}

// Dig runs logic that digs up gold on this Unit's Tile.
func (unitImpl *UnitImpl) Dig(amount int64) bool {
	return unitImpl.RunOnServer("dig", map[string]interface{}{
		"amount": amount,
	}).(bool)
}

// Move runs logic that moves this Unit from its current Tile to an
// adjacent Tile. If this Unit merges with another one, the other Unit will
// be destroyed and its tile will be set to nil. Make sure to check that
// your Unit's tile is not nil before doing things with it.
func (unitImpl *UnitImpl) Move(tile pirates.Tile) bool {
	return unitImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Rest runs logic that regenerates this Unit's health. Must be used in
// range of a port.
func (unitImpl *UnitImpl) Rest() bool {
	return unitImpl.RunOnServer("rest", map[string]interface{}{
	}).(bool)
}

// Split runs logic that moves a number of crew from this Unit to the given
// Tile. This will consume a move from those crew.
func (unitImpl *UnitImpl) Split(tile pirates.Tile, amount int64, gold int64) bool {
	return unitImpl.RunOnServer("split", map[string]interface{}{
		"tile": tile,
		"amount": amount,
		"gold": gold,
	}).(bool)
}

// Withdraw runs logic that takes gold from the Player. You can only
// withdraw from your own Port.
func (unitImpl *UnitImpl) Withdraw(amount int64) bool {
	return unitImpl.RunOnServer("withdraw", map[string]interface{}{
		"amount": amount,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Unit.
func (unitImpl *UnitImpl) InitImplDefaults() {
	unitImpl.GameObjectImpl.InitImplDefaults()

	unitImpl.actedImpl = true
	unitImpl.crewImpl = 0
	unitImpl.crewHealthImpl = 0
	unitImpl.goldImpl = 0
	unitImpl.movesImpl = 0
	unitImpl.ownerImpl = nil
	unitImpl.pathImpl = []pirates.Tile{}
	unitImpl.shipHealthImpl = 0
	unitImpl.stunTurnsImpl = 0
	unitImpl.targetPortImpl = nil
	unitImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Unit.
func (unitImpl *UnitImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := unitImpl.GameObjectImpl.DeltaMerge(
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
	case "acted":
		unitImpl.actedImpl = piratesDeltaMerge.Boolean(delta)
		return true, nil
	case "crew":
		unitImpl.crewImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "crewHealth":
		unitImpl.crewHealthImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "gold":
		unitImpl.goldImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "moves":
		unitImpl.movesImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		unitImpl.ownerImpl = piratesDeltaMerge.Player(delta)
		return true, nil
	case "path":
		unitImpl.pathImpl = piratesDeltaMerge.ArrayOfTile(&unitImpl.pathImpl, delta)
		return true, nil
	case "shipHealth":
		unitImpl.shipHealthImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "stunTurns":
		unitImpl.stunTurnsImpl = piratesDeltaMerge.Int(delta)
		return true, nil
	case "targetPort":
		unitImpl.targetPortImpl = piratesDeltaMerge.Port(delta)
		return true, nil
	case "tile":
		unitImpl.tileImpl = piratesDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

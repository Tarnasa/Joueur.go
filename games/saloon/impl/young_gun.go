package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// YoungGunImpl is the struct that implements the container for YoungGun
// instances in Saloon.
type YoungGunImpl struct {
	GameObjectImpl

	callInTileImpl saloon.Tile
	canCallInImpl  bool
	ownerImpl      saloon.Player
	tileImpl       saloon.Tile
}

// CallInTile returns the Tile that a Cowboy will be called in on if this
// YoungGun calls in a Cowboy.
func (youngGunImpl *YoungGunImpl) CallInTile() saloon.Tile {
	return youngGunImpl.callInTileImpl
}

// CanCallIn returns true if the YoungGun can call in a Cowboy, false
// otherwise.
func (youngGunImpl *YoungGunImpl) CanCallIn() bool {
	return youngGunImpl.canCallInImpl
}

// Owner returns the Player that owns and can control this YoungGun.
func (youngGunImpl *YoungGunImpl) Owner() saloon.Player {
	return youngGunImpl.ownerImpl
}

// Tile returns the Tile this YoungGun is currently on.
func (youngGunImpl *YoungGunImpl) Tile() saloon.Tile {
	return youngGunImpl.tileImpl
}

// CallIn runs logic that tells the YoungGun to call in a new Cowboy of the
// given job to the open Tile nearest to them.
func (youngGunImpl *YoungGunImpl) CallIn(job string) saloon.Cowboy {
	return youngGunImpl.RunOnServer("callIn", map[string]interface{}{
		"job": job,
	}).(saloon.Cowboy)
}

// InitImplDefaults initializes safe defaults for all fields in YoungGun.
func (youngGunImpl *YoungGunImpl) InitImplDefaults() {
	youngGunImpl.GameObjectImpl.InitImplDefaults()

	youngGunImpl.callInTileImpl = nil
	youngGunImpl.canCallInImpl = true
	youngGunImpl.ownerImpl = nil
	youngGunImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in YoungGun.
func (youngGunImpl *YoungGunImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := youngGunImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	saloonDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: "+
			"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "callInTile":
		youngGunImpl.callInTileImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "canCallIn":
		youngGunImpl.canCallInImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "owner":
		youngGunImpl.ownerImpl = saloonDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		youngGunImpl.tileImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

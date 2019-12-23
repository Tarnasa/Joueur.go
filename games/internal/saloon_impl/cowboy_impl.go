package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/saloon"
)

// CowboyImpl is the struct that implements the container for Cowboy
// instances in Saloon.
type CowboyImpl struct {
	GameObjectImpl

	canMoveImpl        bool
	drunkDirectionImpl string
	focusImpl          int64
	healthImpl         int64
	isDeadImpl         bool
	isDrunkImpl        bool
	jobImpl            string
	ownerImpl          saloon.Player
	tileImpl           saloon.Tile
	toleranceImpl      int64
	turnsBusyImpl      int64
}

// CanMove returns if the Cowboy can be moved this turn via its owner.
func (cowboyImpl *CowboyImpl) CanMove() bool {
	return cowboyImpl.canMoveImpl
}

// DrunkDirection returns the direction this Cowboy is moving while drunk.
// Will be 'North', 'East', 'South', or 'West' when drunk; or '' (empty
// string) when not drunk.
//
// Literal Values: "", "North", "East", "South", or "West"
func (cowboyImpl *CowboyImpl) DrunkDirection() string {
	return cowboyImpl.drunkDirectionImpl
}

// Focus returns how much focus this Cowboy has. Different Jobs do
// different things with their Cowboy's focus.
func (cowboyImpl *CowboyImpl) Focus() int64 {
	return cowboyImpl.focusImpl
}

// Health returns how much health this Cowboy currently has.
func (cowboyImpl *CowboyImpl) Health() int64 {
	return cowboyImpl.healthImpl
}

// IsDead returns if this Cowboy is dead and has been removed from the
// game.
func (cowboyImpl *CowboyImpl) IsDead() bool {
	return cowboyImpl.isDeadImpl
}

// IsDrunk returns if this Cowboy is drunk, and will automatically walk.
func (cowboyImpl *CowboyImpl) IsDrunk() bool {
	return cowboyImpl.isDrunkImpl
}

// Job returns the job that this Cowboy does, and dictates how they fight
// and interact within the Saloon.
//
// Literal Values: "Bartender", "Brawler", or "Sharpshooter"
func (cowboyImpl *CowboyImpl) Job() string {
	return cowboyImpl.jobImpl
}

// Owner returns the Player that owns and can control this Cowboy.
func (cowboyImpl *CowboyImpl) Owner() saloon.Player {
	return cowboyImpl.ownerImpl
}

// Tile returns the Tile that this Cowboy is located on.
//
// Value can be returned as a nil pointer.
func (cowboyImpl *CowboyImpl) Tile() saloon.Tile {
	return cowboyImpl.tileImpl
}

// Tolerance returns how many times this unit has been drunk before taking
// their siesta and reseting this to 0.
func (cowboyImpl *CowboyImpl) Tolerance() int64 {
	return cowboyImpl.toleranceImpl
}

// TurnsBusy returns how many turns this unit has remaining before it is no
// longer busy and can `act()` or `play()` again.
func (cowboyImpl *CowboyImpl) TurnsBusy() int64 {
	return cowboyImpl.turnsBusyImpl
}

// Act runs logic that does their job's action on a Tile.
func (cowboyImpl *CowboyImpl) Act(tile saloon.Tile, drunkDirection string) bool {
	return cowboyImpl.RunOnServer("act", map[string]interface{}{
		"tile":           tile,
		"drunkDirection": drunkDirection,
	}).(bool)
}

// Move runs logic that moves this Cowboy from its current Tile to an
// adjacent Tile.
func (cowboyImpl *CowboyImpl) Move(tile saloon.Tile) bool {
	return cowboyImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Play runs logic that sits down and plays a piano.
func (cowboyImpl *CowboyImpl) Play(piano saloon.Furnishing) bool {
	return cowboyImpl.RunOnServer("play", map[string]interface{}{
		"piano": piano,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Cowboy.
func (cowboyImpl *CowboyImpl) InitImplDefaults() {
	cowboyImpl.GameObjectImpl.InitImplDefaults()

	cowboyImpl.canMoveImpl = true
	cowboyImpl.drunkDirectionImpl = ""
	cowboyImpl.focusImpl = 0
	cowboyImpl.healthImpl = 0
	cowboyImpl.isDeadImpl = true
	cowboyImpl.isDrunkImpl = true
	cowboyImpl.jobImpl = ""
	cowboyImpl.ownerImpl = nil
	cowboyImpl.tileImpl = nil
	cowboyImpl.toleranceImpl = 0
	cowboyImpl.turnsBusyImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Cowboy.
func (cowboyImpl *CowboyImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := cowboyImpl.GameObjectImpl.DeltaMerge(
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
			"deltaMerge is not the expected type of: " +
				"'saloon.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "canMove":
		cowboyImpl.canMoveImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "drunkDirection":
		cowboyImpl.drunkDirectionImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "focus":
		cowboyImpl.focusImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "health":
		cowboyImpl.healthImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "isDead":
		cowboyImpl.isDeadImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "isDrunk":
		cowboyImpl.isDrunkImpl = saloonDeltaMerge.Boolean(delta)
		return true, nil
	case "job":
		cowboyImpl.jobImpl = saloonDeltaMerge.String(delta)
		return true, nil
	case "owner":
		cowboyImpl.ownerImpl = saloonDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		cowboyImpl.tileImpl = saloonDeltaMerge.Tile(delta)
		return true, nil
	case "tolerance":
		cowboyImpl.toleranceImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	case "turnsBusy":
		cowboyImpl.turnsBusyImpl = saloonDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

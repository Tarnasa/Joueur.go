package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/galapagos"
)

// CreatureImpl is the struct that implements the container for Creature
// instances in Galapagos.
type CreatureImpl struct {
	GameObjectImpl

	canBiteImpl       bool
	canBreedImpl      bool
	carnivorismImpl   int64
	currentHealthImpl int64
	defenseImpl       int64
	enduranceImpl     int64
	herbivorismImpl   int64
	isEggImpl         bool
	maxHealthImpl     int64
	movementLeftImpl  int64
	ownerImpl         galapagos.Player
	parentsImpl       []galapagos.Creature
	speedImpl         int64
	tileImpl          galapagos.Tile
}

// CanBite returns indicates whether or not this creature can bite this
// turn.
func (creatureImpl *CreatureImpl) CanBite() bool {
	return creatureImpl.canBiteImpl
}

// CanBreed returns indicates whether or not this creature can breed this
// turn.
func (creatureImpl *CreatureImpl) CanBreed() bool {
	return creatureImpl.canBreedImpl
}

// Carnivorism returns the carnivore level of the creature. This increases
// damage to other other creatures and health restored on kill.
func (creatureImpl *CreatureImpl) Carnivorism() int64 {
	return creatureImpl.carnivorismImpl
}

// CurrentHealth returns the current amount of health that this creature
// has.
func (creatureImpl *CreatureImpl) CurrentHealth() int64 {
	return creatureImpl.currentHealthImpl
}

// Defense returns the defense of the creature. This reduces the amount of
// damage this creature takes from being eaten.
func (creatureImpl *CreatureImpl) Defense() int64 {
	return creatureImpl.defenseImpl
}

// Endurance returns the endurance level of the creature. This increases
// the max health a creature can have.
func (creatureImpl *CreatureImpl) Endurance() int64 {
	return creatureImpl.enduranceImpl
}

// Herbivorism returns the herbivore level of the creature. This increases
// health restored from eating plants.
func (creatureImpl *CreatureImpl) Herbivorism() int64 {
	return creatureImpl.herbivorismImpl
}

// IsEgg returns indicates whether or not this creature is still in an egg
// and cannot bite, breed, or be bitten.
func (creatureImpl *CreatureImpl) IsEgg() bool {
	return creatureImpl.isEggImpl
}

// MaxHealth returns the maximum amount of health this creature can have.
func (creatureImpl *CreatureImpl) MaxHealth() int64 {
	return creatureImpl.maxHealthImpl
}

// MovementLeft returns the amount of moves this creature has left this
// turn.
func (creatureImpl *CreatureImpl) MovementLeft() int64 {
	return creatureImpl.movementLeftImpl
}

// Owner returns the owner of the creature.
//
// Value can be returned as a nil pointer.
func (creatureImpl *CreatureImpl) Owner() galapagos.Player {
	return creatureImpl.ownerImpl
}

// Parents returns the creatures that gave birth to this one.
func (creatureImpl *CreatureImpl) Parents() []galapagos.Creature {
	return creatureImpl.parentsImpl
}

// Speed returns the speed of the creature. This determines how many times
// a creature can move in one turn.
func (creatureImpl *CreatureImpl) Speed() int64 {
	return creatureImpl.speedImpl
}

// Tile returns the Tile this Creature occupies.
//
// Value can be returned as a nil pointer.
func (creatureImpl *CreatureImpl) Tile() galapagos.Tile {
	return creatureImpl.tileImpl
}

// Bite runs logic that command a creature to bite a plant or creature on
// the specified tile.
func (creatureImpl *CreatureImpl) Bite(tile galapagos.Tile) bool {
	return creatureImpl.RunOnServer("bite", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// Breed runs logic that command a creature to breed with an adjacent
// creature.
func (creatureImpl *CreatureImpl) Breed(mate galapagos.Creature) galapagos.Creature {
	return creatureImpl.RunOnServer("breed", map[string]interface{}{
		"mate": mate,
	}).(galapagos.Creature)
}

// Move runs logic that command a creature to move to a specified adjacent
// tile.
func (creatureImpl *CreatureImpl) Move(tile galapagos.Tile) bool {
	return creatureImpl.RunOnServer("move", map[string]interface{}{
		"tile": tile,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Creature.
func (creatureImpl *CreatureImpl) InitImplDefaults() {
	creatureImpl.GameObjectImpl.InitImplDefaults()

	creatureImpl.canBiteImpl = true
	creatureImpl.canBreedImpl = true
	creatureImpl.carnivorismImpl = 0
	creatureImpl.currentHealthImpl = 0
	creatureImpl.defenseImpl = 0
	creatureImpl.enduranceImpl = 0
	creatureImpl.herbivorismImpl = 0
	creatureImpl.isEggImpl = true
	creatureImpl.maxHealthImpl = 0
	creatureImpl.movementLeftImpl = 0
	creatureImpl.ownerImpl = nil
	creatureImpl.parentsImpl = []galapagos.Creature{}
	creatureImpl.speedImpl = 0
	creatureImpl.tileImpl = nil
}

// DeltaMerge merges the delta for a given attribute in Creature.
func (creatureImpl *CreatureImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := creatureImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	galapagosDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'galapagos.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "canBite":
		creatureImpl.canBiteImpl = galapagosDeltaMerge.Boolean(delta)
		return true, nil
	case "canBreed":
		creatureImpl.canBreedImpl = galapagosDeltaMerge.Boolean(delta)
		return true, nil
	case "carnivorism":
		creatureImpl.carnivorismImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "currentHealth":
		creatureImpl.currentHealthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "defense":
		creatureImpl.defenseImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "endurance":
		creatureImpl.enduranceImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "herbivorism":
		creatureImpl.herbivorismImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "isEgg":
		creatureImpl.isEggImpl = galapagosDeltaMerge.Boolean(delta)
		return true, nil
	case "maxHealth":
		creatureImpl.maxHealthImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "movementLeft":
		creatureImpl.movementLeftImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		creatureImpl.ownerImpl = galapagosDeltaMerge.Player(delta)
		return true, nil
	case "parents":
		creatureImpl.parentsImpl = galapagosDeltaMerge.ArrayOfCreature(&creatureImpl.parentsImpl, delta)
		return true, nil
	case "speed":
		creatureImpl.speedImpl = galapagosDeltaMerge.Int(delta)
		return true, nil
	case "tile":
		creatureImpl.tileImpl = galapagosDeltaMerge.Tile(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

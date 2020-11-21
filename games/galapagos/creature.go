package galapagos

// Creature is a Creature in the game.
type Creature interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// CanBite is indicates whether or not this creature can bite this
	// turn.
	CanBite() bool

	// CanBreed is indicates whether or not this creature can breed
	// this turn.
	CanBreed() bool

	// Carnivorism is the carnivore level of the creature. This
	// increases damage to other other creatures and health restored
	// on kill.
	Carnivorism() int64

	// CurrentHealth is the current amount of health that this
	// creature has.
	CurrentHealth() int64

	// Defense is the defense of the creature. This reduces the amount
	// of damage this creature takes from being eaten.
	Defense() int64

	// Endurance is the endurance level of the creature. This
	// increases the max health a creature can have.
	Endurance() int64

	// Herbivorism is the herbivore level of the creature. This
	// increases health restored from eating plants.
	Herbivorism() int64

	// IsEgg is indicates whether or not this creature is still in an
	// egg and cannot bite, breed, or be bitten.
	IsEgg() bool

	// MaxHealth is the maximum amount of health this creature can
	// have.
	MaxHealth() int64

	// MovementLeft is the amount of moves this creature has left this
	// turn.
	MovementLeft() int64

	// Owner is the owner of the creature.
	//
	// Value can be returned as a nil pointer.
	Owner() Player

	// Parents is the creatures that gave birth to this one.
	Parents() []Creature

	// Speed is the speed of the creature. This determines how many
	// times a creature can move in one turn.
	Speed() int64

	// Tile is the Tile this Creature occupies.
	//
	// Value can be returned as a nil pointer.
	Tile() Tile

	// -- Methods -- \\

	// Bite command a creature to bite a plant or creature on the
	// specified tile.
	Bite(Tile) bool

	// Breed command a creature to breed with an adjacent creature.
	Breed(Creature) Creature

	// Move command a creature to move to a specified adjacent tile.
	Move(Tile) bool
}

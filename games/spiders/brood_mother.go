package spiders

// BroodMother is the Spider Queen. She alone can spawn Spiderlings for
// each Player, and if she dies the owner loses.
type BroodMother interface {
	// Parent interfaces
	Spider

	// -- Attributes -- \\

	// Eggs is how many eggs the BroodMother has to spawn Spiderlings
	// this turn.
	Eggs() int64

	// Health is how much health this BroodMother has left. When it
	// reaches 0, she dies and her owner loses.
	Health() int64

	// -- Methods -- \\

	// Consume consumes a Spiderling of the same owner to regain some
	// eggs to spawn more Spiderlings.
	Consume(Spiderling) bool

	// Spawn spawns a new Spiderling on the same Nest as this
	// BroodMother, consuming an egg.
	Spawn(string) Spiderling
}

package saloon

// YoungGun is an eager young person that wants to join your gang, and will
// call in the veteran Cowboys you need to win the brawl in the saloon.
type YoungGun interface {
	// Parent interfaces
	GameObject

	// -- Attributes -- \\

	// CallInTile is the Tile that a Cowboy will be called in on if
	// this YoungGun calls in a Cowboy.
	CallInTile() Tile

	// CanCallIn is true if the YoungGun can call in a Cowboy, false
	// otherwise.
	CanCallIn() bool

	// Owner is the Player that owns and can control this YoungGun.
	Owner() Player

	// Tile is the Tile this YoungGun is currently on.
	Tile() Tile

	// -- Methods -- \\

	// CallIn tells the YoungGun to call in a new Cowboy of the given
	// job to the open Tile nearest to them.
	CallIn(string) Cowboy
}

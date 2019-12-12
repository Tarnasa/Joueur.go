package base

// Player is the base interface all Player GameObjects within the game must implement.
type Player interface {
	GameObject

	/** If the player won the game or not. */
	Won() bool

	/** If the player lost the game or not. */
	Lost() bool

	/** The reason why the player won the game. */
	ReasonWon() string

	/** The reason why the player lost the game. */
	ReasonLost() string
}

func asPlayer(gameObject GameObject) Player {
	if player, isPlayer := gameObject.(Player); isPlayer {
		return player
	}

	return nil
}

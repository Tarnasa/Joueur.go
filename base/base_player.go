package base

type BasePlayer interface {
	BaseGameObject

	/** If the player won the game or not. */
	Won() bool

	/** If the player lost the game or not. */
	Lost() bool

	/** The reason why the player won the game. */
	ReasonWon() string

	/** The reason why the player lost the game. */
	ReasonLost() string
}

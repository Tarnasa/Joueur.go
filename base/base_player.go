package base

type BasePlayer struct {
	/** The name of the player. */
	Name string

	/**
	* What type of client this is,
	* For potential data mining purposes.
	*
	* @example "Python", "JavaScript", or some other language.
	 */
	ClientType string

	/** The amount of time (in ns) remaining for this AI to send commands. */
	TimeRemaining float64

	/** If the player won the game or not. */
	Won bool

	/** If the player lost the game or not. */
	Lost bool

	/** The reason why the player won the game. */
	ReasonWon string

	/** The reason why the player lost the game. */
	ReasonLost string
}

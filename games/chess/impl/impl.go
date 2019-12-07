// Package impl contains all the Chess implimentation
// logic and structures required by aa client to play with a game server.
// To start coding your AI open ./ai.go
package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/chess"
)

// -- Game -- \\

// GameImpl is the struct that implements the container for Game instances in Chess.
type GameImpl struct {
	base.GameImpl
	FenImpl         string
	GameObjectsImpl map[string]chess.GameObject
	HistoryImpl     []string
	PlayersImpl     []chess.Player
	SessionImpl     string
}

// Fen returns forsyth-Edwards Notation (fen), a notation that describes the game board state.
func (gameImpl *GameImpl) Fen() string {
	return gameImpl.FenImpl
}

// GameObjects returns a mapping of every game object's ID to the actual game object. Primarily used by the server and client to easily refer to the game objects via ID.
func (gameImpl *GameImpl) GameObjects() map[string]chess.GameObject {
	return gameImpl.GameObjectsImpl
}

// History returns the array of [known] moves that have occurred in the game, in Standard Algebraic Notation (SAN) format. The first element is the first move, with the last being the most recent.
func (gameImpl *GameImpl) History() []string {
	return gameImpl.HistoryImpl
}

// Players returns array of all the players in the game.
func (gameImpl *GameImpl) Players() []chess.Player {
	return gameImpl.PlayersImpl
}

// Session returns a unique identifier for the game instance that is being played.
func (gameImpl *GameImpl) Session() string {
	return gameImpl.SessionImpl
}

// InitImplDefaults initializes safe defaults for all fields in Game.
func (gameImpl *GameImpl) InitImplDefaults() {
	gameImpl.GameImpl.InitImplDefaults()

	gameImpl.FenImpl = ""
	gameImpl.GameObjectsImpl = make(map[string]chess.GameObject)
	gameImpl.HistoryImpl = make([]string, 0)
	gameImpl.PlayersImpl = make([]chess.Player, 0)
	gameImpl.SessionImpl = ""
}

// -- GameObject -- \\

// GameObjectImpl is the struct that implements the container for GameObject instances in Chess.
type GameObjectImpl struct {
	base.GameObjectImpl
	game               *GameImpl
	GameObjectNameImpl string
	LogsImpl           []string
}

// Game returns a pointer to the Chess Game instance
func (gameObjectImpl *GameObjectImpl) Game() chess.Game {
	return gameObjectImpl.game
}

// GameObjectName returns string representing the top level Class that this game object is an instance of. Used for reflection to create new instances on clients, but exposed for convenience should AIs want this data.
func (gameObjectImpl *GameObjectImpl) GameObjectName() string {
	return gameObjectImpl.GameObjectNameImpl
}

// Logs returns any strings logged will be stored here. Intended for debugging.
func (gameObjectImpl *GameObjectImpl) Logs() []string {
	return gameObjectImpl.LogsImpl
}

// Log runs logic that adds a message to this GameObject's logs. Intended for your own debugging purposes, as strings stored here are saved in the gamelog.
func (gameObjectImpl *GameObjectImpl) Log(message string) {
	gameObjectImpl.RunOnServer("log", map[string]interface{}{
		"message": message,
	})
}

// InitImplDefaults initializes safe defaults for all fields in GameObject.
func (gameObjectImpl *GameObjectImpl) InitImplDefaults() {
	gameObjectImpl.GameObjectImpl.InitImplDefaults()

	gameObjectImpl.GameObjectNameImpl = ""
	gameObjectImpl.IdImpl = ""
	gameObjectImpl.LogsImpl = make([]string, 0)
}

// -- Player -- \\

// PlayerImpl is the struct that implements the container for Player instances in Chess.
type PlayerImpl struct {
	GameObjectImpl
	ClientTypeImpl    string
	ColorImpl         string
	LostImpl          bool
	NameImpl          string
	OpponentImpl      chess.Player
	ReasonLostImpl    string
	ReasonWonImpl     string
	TimeRemainingImpl float64
	WonImpl           bool
}

// ClientType returns what type of client this is, e.g. 'Python', 'JavaScript', or some other language. For potential data mining purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.ClientTypeImpl
}

// Color returns the color (side) of this player. Either 'white' or 'black', with the 'white' player having the first move.
func (playerImpl *PlayerImpl) Color() string {
	return playerImpl.ColorImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.LostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.NameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() chess.Player {
	return playerImpl.OpponentImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.ReasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.ReasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.TimeRemainingImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.WonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.ClientTypeImpl = ""
	playerImpl.ColorImpl = ""
	playerImpl.LostImpl = false
	playerImpl.NameImpl = ""
	playerImpl.OpponentImpl = nil
	playerImpl.ReasonLostImpl = ""
	playerImpl.ReasonWonImpl = ""
	playerImpl.TimeRemainingImpl = 0
	playerImpl.WonImpl = false
}

func (playerImpl *PlayerImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) {
	switch(attribute) {
	case "":
		(*playerImpl.clientTypeImpl) = deltaMerge.String(delta)
		break
	case "":
		(*playerImpl.colorImpl) = deltaMerge.String(delta)
		break
	case "":
		(*playerImpl.lostImpl) = deltaMerge.Boolean(delta)
		break
	case "":
		(*playerImpl.nameImpl) = deltaMerge.String(delta)
		break
	case "":
		(*playerImpl.opponentImpl) = deltaMerge.Player(delta)
		break
	case "":
		(*playerImpl.reasonLostImpl) = deltaMerge.String(delta)
		break
	case "":
		(*playerImpl.reasonWonImpl) = deltaMerge.String(delta)
		break
	case "":
		(*playerImpl.timeRemainingImpl) = deltaMerge.Float(delta)
		break
	case "":
		(*playerImpl.wonImpl) = deltaMerge.Boolean(delta)
		break
	}
}

// -- Namespace -- \\

// ChessNamespace is the collection of implimentation logic for the Chess game.
type ChessNamespace struct {}

// Name returns the name of the Chess game.
func (*ChessNamespace) Name() string {
	return "Chess"
}

// Version returns the current version hash as last generated for the Chess game.
func (*ChessNamespace) Version() string {
	return "cfa5f5c1685087ce2899229c04c26e39f231e897ecc8fe036b44bc22103ef801"
}

// PlayerName returns the desired name of the AI in the Chess game.
func (*ChessNamespace) PlayerName() string {
	return chess.PlayerName()
}

// CreateGameObject is the factory method for all GameObject instances in the Chess game.
func (*ChessNamespace) CreateGameObject(gameObjectName string) (base.GameObject, *base.DeltaMergeableImpl, error) {
	switch (gameObjectName) {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InitImplDefaults()
		return &newGameObject, &(newGameObject.GameObjectImpl.DeltaMergeableImpl), nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InitImplDefaults()
		return &newPlayer, &(newPlayer.GameObjectImpl.DeltaMergeableImpl), nil
	}
	return nil, nil, errors.New("No game object named " + gameObjectName + " for game Chess")
}

// CreateGame is the factory method for Game the Chess game.
func (*ChessNamespace) CreateGame() (base.Game, *base.DeltaMergeableImpl) {
	game := GameImpl{}
	game.InitImplDefaults()
	return &game, &(game.GameImpl.DeltaMergeableImpl)
}

// CreateAI is the factory method for the AI in the Chess game.
func (*ChessNamespace) CreateAI() (base.AI, *base.AIImpl) {
	ai := chess.AI{}
	return &ai, &ai.AIImpl
}

// OrderAI handles an order for the AI in the Chess game.
func (*ChessNamespace) OrderAI(baseAI base.AI, functionName string, args []interface{}) (interface{}, error) {
	ai, validAI := baseAI.(*chess.AI)
	if !validAI {
		return nil, errors.New("AI is not a valid chess.AI to order!")
	}
	switch (functionName) {
	case "makeMove":
		return (*ai).MakeMove(), nil
	}

	return nil, errors.New("Cannot find functionName "+functionName+" to run in S{game_name} AI")
}

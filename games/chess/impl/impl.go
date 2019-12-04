// This package contains all the structs, methods, and the AI required as
// a client to play the Chess with a game server.
// To start coding your AI open ./ai.go
package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/chess"
)

// -- Game -- \\

type GameImpl struct {
	base.BaseGameImpl
}

func (this GameImpl) Fen() string {
	return this.InternalDataMap["fen"].(string)
}

func (this GameImpl) GameObjects() map[string](chess.GameObject) {
	return this.InternalDataMap["gameObjects"].(map[string](chess.GameObject))
}

func (this GameImpl) History() []string {
	return this.InternalDataMap["history"].([]string)
}

func (this GameImpl) Players() [](chess.Player) {
	return this.InternalDataMap["players"].([](chess.Player))
}

func (this GameImpl) Session() string {
	return this.InternalDataMap["session"].(string)
}

func defaultInternalDataMapForGame() map[string]interface{} {
	data := make(map[string]interface{})
	data["fen"] = ""
	data["gameObjects"] = make(map[string](chess.GameObject))
	data["history"] = make([]string, 0)
	data["players"] = make([](chess.Player), 0)
	data["session"] = ""

	return data
}

// -- GameObject -- \\

type GameObjectImpl struct {
	base.BaseGameObjectImpl
	game *chess.Game
}

func (this GameObjectImpl) Game() *chess.Game {
	return this.game
}

func (this GameObjectImpl) GameObjectName() string {
	return this.InternalDataMap["gameObjectName"].(string)
}

func (this GameObjectImpl) Id() string {
	return this.InternalDataMap["id"].(string)
}

func (this GameObjectImpl) Logs() []string {
	return this.InternalDataMap["logs"].([]string)
}

func (this GameObjectImpl) Log(message string) {
	args := make(map[string]interface{})
	args["message"] = message
	this.RunOnServer("log", args)
}

func defaultInternalDataMapForGameObject() map[string]interface{} {
	data := make(map[string]interface{})
	data["gameObjectName"] = ""
	data["id"] = ""
	data["logs"] = make([]string, 0)

	return data
}

// -- Player -- \\

type PlayerImpl struct {
	GameObjectImpl
}

func (this PlayerImpl) ClientType() string {
	return this.InternalDataMap["clientType"].(string)
}

func (this PlayerImpl) Color() string {
	return this.InternalDataMap["color"].(string)
}

func (this PlayerImpl) Lost() bool {
	return this.InternalDataMap["lost"].(bool)
}

func (this PlayerImpl) Name() string {
	return this.InternalDataMap["name"].(string)
}

func (this PlayerImpl) Opponent() (chess.Player) {
	return this.InternalDataMap["opponent"].((chess.Player))
}

func (this PlayerImpl) ReasonLost() string {
	return this.InternalDataMap["reasonLost"].(string)
}

func (this PlayerImpl) ReasonWon() string {
	return this.InternalDataMap["reasonWon"].(string)
}

func (this PlayerImpl) TimeRemaining() float64 {
	return this.InternalDataMap["timeRemaining"].(float64)
}

func (this PlayerImpl) Won() bool {
	return this.InternalDataMap["won"].(bool)
}

func defaultInternalDataMapForPlayer() map[string]interface{} {
	data := make(map[string]interface{})
	parentData0 := defaultInternalDataMapForGameObject()
	for key, value := range parentData0 {
		data[key] = value
	}
	data["clientType"] = ""
	data["color"] = ""
	data["lost"] = false
	data["name"] = ""
	data["opponent"] = nil
	data["reasonLost"] = ""
	data["reasonWon"] = ""
	data["timeRemaining"] = 0
	data["won"] = false

	return data
}

// -- Namespace -- \
type ChessNamespace struct {}

func (_ ChessNamespace) Name() string {
	return "Chess"
}

func (_ ChessNamespace) Version() string {
	return "cfa5f5c1685087ce2899229c04c26e39f231e897ecc8fe036b44bc22103ef801"
}

func (_ ChessNamespace) PlayerName() string {
	return chess.PlayerName()
}

func (_ ChessNamespace) CreateGameObject(gameObjectName string) (base.BaseGameObject, *base.BaseDeltaMergeableImpl, error) {
	switch (gameObjectName) {
	case "GameObject":
		newGameObject := GameObjectImpl{}
		newGameObject.InternalDataMap = defaultInternalDataMapForGameObject()
		return &newGameObject, &(newGameObject.BaseGameObjectImpl.BaseDeltaMergeableImpl), nil
	case "Player":
		newPlayer := PlayerImpl{}
		newPlayer.InternalDataMap = defaultInternalDataMapForPlayer()
		return &newPlayer, &(newPlayer.BaseGameObjectImpl.BaseDeltaMergeableImpl), nil
	}
	return nil, nil, errors.New("No game object named " + gameObjectName + " for game Chess")
}

func (_ ChessNamespace) CreateGame() (base.BaseGame, *base.BaseDeltaMergeableImpl) {
	game := GameImpl{}
	return &game, &(game.BaseGameImpl.BaseDeltaMergeableImpl)
}

func (_ ChessNamespace) CreateAI() (base.BaseAI, *base.BaseAIImpl) {
	ai := chess.AI{}
	return &ai, &ai.BaseAIImpl
}

func (_ ChessNamespace) OrderAI(baseAI base.BaseAI, functionName string, args []interface{}) (interface{}, error) {
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

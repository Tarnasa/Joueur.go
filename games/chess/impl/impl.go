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
	GameObjects map[string]*chess.GameObject
	data map[string]interface{}
}

func (this GameImpl) Fen() string {
	return this.data["fen"].(string)
}

func (this GameImpl) History() []string {
	return this.data["history"].([]string)
}

func (this GameImpl) Players() []*(chess.Player) {
	return this.data["players"].([]*(chess.Player))
}

func (this GameImpl) Session() string {
	return this.data["session"].(string)
}

// -- GameObject -- \\

type GameObjectImpl struct {
	base.BaseGameObjectImpl
	game *Game
	data map[string]interface{}
}

func (this GameObjectImpl) Game() *chess.Game {
	return this.game
}

func (this GameObjectImpl) GameObjectName() string {
	return this.data["gameObjectName"].(string)
}

func (this GameObjectImpl) Id() string {
	return this.data["id"].(string)
}

func (this GameObjectImpl) Logs() []string {
	return this.data["logs"].([]string)
}

func (this GameObjectImpl) Log(message string) {
	this.RunOnServer(make(map[string]interface{
		"message": message,
	}))
}

// -- Player -- \\

type PlayerImpl struct {
	GameObject
	data map[string]interface{}
}

func (this PlayerImpl) ClientType() string {
	return this.data["clientType"].(string)
}

func (this PlayerImpl) Color() string {
	return this.data["color"].(string)
}

func (this PlayerImpl) Lost() bool {
	return this.data["lost"].(bool)
}

func (this PlayerImpl) Name() string {
	return this.data["name"].(string)
}

func (this PlayerImpl) Opponent() *(chess.Player) {
	return this.data["opponent"].(*(chess.Player))
}

func (this PlayerImpl) ReasonLost() string {
	return this.data["reasonLost"].(string)
}

func (this PlayerImpl) ReasonWon() string {
	return this.data["reasonWon"].(string)
}

func (this PlayerImpl) TimeRemaining() float64 {
	return this.data["timeRemaining"].(float64)
}

func (this PlayerImpl) Won() bool {
	return this.data["won"].(bool)
}

// Factory functions

func CreateGameObject(gameObjectName string) (*chess.GameObject, error) {
	switch (gameObjectName) {
	case "GameObject":
		return &(GameObjectImpl{}), nil
	case "Player":
		return &(PlayerImpl{}), nil
	}
	return nil, errors.New("No game object named " + gameObjectName + " for game Chess")
}

func CreateGame() *chess.Game {
	return &(GameImpl{})
}

func CreateAI() *chess.AI {
	return &(chess.AI{})
}

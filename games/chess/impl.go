// This package contains all the structs, methods, and the AI required as
// a client to play the Chess with a game server.
// To start coding your AI open ./ai.go
package chess

import (
	"errors"
	"joueur/base"
)

type GameObjectImpl struct {
	base.BaseGameObjectImpl
	game *Game
	data map[string]interface{}
}

func (this GameObjectImpl) Game() *Game {
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

func (this PlayerImpl) Opponent() *Player {
	return this.data["opponent"].(*Player)
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

func CreateGameObject(gameObjectName string) (*GameObject, error) {
	switch (gameObjectName) {
	case "GameObject":
		return &(GameObjectImpl{}), nil
	case "Player":
		return &(PlayerImpl{}), nil
	}
	return nil, errors.New("No game object named " + gameObjectName + " for game Chess")
}

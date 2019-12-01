package client

import (
	"fmt"
	"joueur/runtime/errorhandler"
	"net"
	"sync"
)

type Client struct {
	conn    *net.Conn
	printIO bool
}

// End of transmission char code is 4
const eotChar = byte(4)

var instance *Client
var once sync.Once

var eventDeltaHandler = func(map[string]interface{}) {}
var eventOverHandler = func() {}

func RegisterEventDeltaHandler(handler func(map[string]interface{})) {
	eventDeltaHandler = handler
}

func RegisterEventOverHandler(handler func()) {
	eventOverHandler = handler
}

func Setup(printIO bool) *Client {
	once.Do(func() {
		instance = &Client{
			printIO: printIO,
		}

		errorhandler.RegisterErrorHandler(func() {
			Disconnect()
		})
	})
	return instance
}

func Connect(address string) error {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	instance.conn = &conn

	return nil
}

func Disconnect() {
	if instance.conn != nil {
		fmt.Println("disconnecting...")
		(*instance.conn).Close()
	}
}

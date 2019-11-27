package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"joueur/base/errorcodes"
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Client struct {
	Coolbeans string
	conn      net.Conn
	printIO   bool
}

var instance *Client
var once sync.Once

func Setup() *Client {
	once.Do(func() {
		instance = &Client{
			Coolbeans: "lololol",
		}
	})
	return instance
}

func Connect(server string, port string) error {
	address := server + ":" + port
	fmt.Println("gonna connect to", address)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	instance.conn = conn

	return nil
}

func Disconnect() {
	if instance.conn != nil {
		instance.conn.Close()
	}
}

func SendEventDataString(event string, data) {
	json.Marshal(interface{
		event: event,
		data: data,
		sentTime: time.Now(),
	})
}

func sendRaw(bytes []byte) error {
	/**
	 * Sends a raw string to the game server
	 * @param str The string to send.
	 */
	if instance.conn == nil {
		return errors.New("Cannot write to socket that has not been initialized")
	}

	if instance.printIO {
		color.Magenta("TO SERVER <-- " + string(bytes))
	}

	_, err := instance.conn.Write((bytes))
	if err != nil {
		HandleError(
			errorcodes.DisconnectedUnexpectedly,
			err,
			"Could not send string through server.",
		)
	}

	return nil
}

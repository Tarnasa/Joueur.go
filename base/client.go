package base

import (
	"errors"
	"fmt"
	"joueur/base/errorcodes"
	"net"
	"sync"

	"github.com/fatih/color"
)

type Client struct {
	Coolbeans string
	conn      net.Conn
	printIO   bool
}

var instance *Client
var once sync.Once

func GetClient() *Client {
	once.Do(func() {
		instance = &Client{
			Coolbeans: "lololol",
		}
	})
	return instance
}

func (client Client) Connect(server string, port string) error {
	address := server + ":" + port
	fmt.Println("gonna connect to", address)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	client.conn = conn

	return nil
}

func (client Client) Disconnect() {
	if client.conn != nil {
		client.conn.Close()
	}
}

func (client Client) sendRaw(bytes []byte) error {
	/**
	 * Sends a raw string to the game server
	 * @param str The string to send.
	 */
	if client.conn == nil {
		return errors.New("Cannot write to socket that has not been initialized")
	}

	if client.printIO {
		color.Magenta("TO SERVER <-- " + string(bytes))
	}

	_, err := client.conn.Write((bytes))
	if err != nil {
		HandleError(
			errorcodes.DisconnectedUnexpectedly,
			err,
			"Could not send string through server.",
		)
	}

	return nil
}

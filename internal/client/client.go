// Package client implements game client basics to with a Cadre game server.
package client

import (
	"joueur/internal/errorhandler"
	"net"
	"sync"
)

// Client is the singleton instance wrapper to talk to a Cadre server
type Client struct {
	conn    *net.TCPConn
	printIO bool
}

// End of transmission char code is 4
const eotChar = byte(4)

var instance *Client
var once sync.Once

// Setup sets up the client singleton instance and returns it
func Setup(printIO bool) *Client {
	once.Do(func() {
		instance = &Client{
			printIO: printIO,
		}

		errorhandler.ErrorHandler = func() {
			Disconnect()
		}
	})
	return instance
}

// Connect connects the instance to a given tcp address
func Connect(address string) error {
	addr, addrError := net.ResolveTCPAddr("tcp", address)
	if addrError != nil {
		return addrError
	}

	conn, dialError := net.DialTCP("tcp", nil, addr)
	if dialError != nil {
		return dialError
	}

	conn.SetNoDelay(true) // disable nagle for optimal latency
	instance.conn = conn

	return nil
}

// Disconnect disconnects the connection from the Cadre server
func Disconnect() {
	if instance.conn != nil {
		(*instance.conn).Close()
	}
}

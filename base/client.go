package base

import (
	"fmt"
	"sync"
)

type Client struct {
	Coolbeans string
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

func (client Client) Connect(server string, port int) error {
	fmt.Println("gonna connect to", server, port)
	return nil
}

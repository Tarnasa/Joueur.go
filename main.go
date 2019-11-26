package main

import (
	"fmt"
	"os"

	"github.com/JacobFischer/argparse"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("joueur.go", "Run the Go lang client with options to connect to a gameserver. Must provide a game name to play.")
	// Create string flag
	parser.GetArgs()

	gameNameArray := parser.PosString("game", &argparse.Options{
		// Required: true, never parses a result, wait till PR merged/fixed to add back
		Help: "the name of the game you want to play on the server",
	})

	server := parser.String("-s", "--server", &argparse.Options{
		Default: "127.0.0.1",
		Help:    "the hostname or the server you want to connect to e.g. locahost:3000",
	})

	port := parser.Int("-p", "--port", &argparse.Options{
		Default: 3000,
		Help:    "the port to connect to on the server. Can be defined on the server arg via server:port",
	})

	playerName := parser.String("-n", "--name", &argparse.Options{
		Help: "the name you want to use as your AI\"s player name",
	})

	index := parser.Int("-i", "--index", &argparse.Options{
		Help: "the player number you want to be, with 0 being the first player",
	})

	password := parser.String("-w", "--password", &argparse.Options{
		Help: "the password required for authentication on official servers",
	})

	session := parser.String("-r", "--session", &argparse.Options{
		Default: "*",
		Help:    "the requested game session you want to play on the server",
	})

	gameSettings := parser.String("", "--gameSettings", &argparse.Options{
		Help: "Any settings for the game server to force. Must be url parms formatted (key=value&otherKey=otherValue)",
	})

	aiSettings := parser.String("", "--aiSettings", &argparse.Options{
		Help: "Any settings for the AI. Delimit pairs by an ampersand (key=value&otherKey=otherValue)",
	})

	printoIO := parser.Flag("", "--printIO", &argparse.Options{
		Default: false,
		Help:    "(debugging) print IO through the TCP socket to the terminal",
	})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil || len(*gameNameArray) != 1 {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	} else {
		// Finally print the collected string
		fmt.Println((*gameNameArray)[0], *server, *port, *playerName, *index, *password, *session, *gameSettings, *aiSettings, *printoIO)
	}
}

// Package joueur implements the base logic required to interact with the
// Cadre AI framework to facilitate playing AI base games.
package main

import (
	"fmt"
	"joueur/internal"
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
		Help: "The name of the game you want to play on the server",
	})

	server := parser.String("s", "server", &argparse.Options{
		Help: "The hostname or the server you want to connect to e.g. locahost:3000",
	})

	port := parser.String("p", "port", &argparse.Options{
		Help: "The port to connect to on the server. Can be defined on the server arg via server:port",
	})

	playerName := parser.String("n", "name", &argparse.Options{
		Help: "The name you want to use as your AI\"s player name",
	})

	index := parser.String("i", "index", &argparse.Options{
		Help: "The player number you want to be, with 0 being the first player",
	})

	password := parser.String("w", "password", &argparse.Options{
		Help: "The password required for authentication on official servers",
	})

	session := parser.String("r", "session", &argparse.Options{
		Help: "The requested game session you want to play on the server",
	})

	gameSettings := parser.String("", "gameSettings", &argparse.Options{
		Help: "Any settings for the game server to force. Must be url parms formatted (key=value&otherKey=otherValue)",
	})

	aiSettings := parser.String("", "aiSettings", &argparse.Options{
		Help: "Any settings for the AI. Delimit pairs by an ampersand (key=value&otherKey=otherValue)",
	})

	printIO := parser.Flag("", "printIO", &argparse.Options{
		Help: "(debugging) print IO through the TCP socket to the terminal",
	})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil || len(*gameNameArray) != 1 {
		// In case of error print error and print usage
		// This can also be done by passing -h or help flags
		fmt.Print(parser.Usage(err))
	} else {
		// Finally print the collected string
		internal.Run(internal.RunArgs{
			Server:       *server,
			Port:         *port,
			GameName:     (*gameNameArray)[0],
			AISettings:   *aiSettings,
			PlayerName:   *playerName,
			Password:     *password,
			Session:      *session,
			Index:        *index,
			GameSettings: *gameSettings,
			PrintIO:      *printIO,
		})
	}
}

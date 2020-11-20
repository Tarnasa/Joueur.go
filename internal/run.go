// Package runtime contains the runtime logic of this client to talk to and
// interact with a game server.
package internal

import (
	"errors"
	"fmt"
	"joueur/games"
	"joueur/internal/client"
	"joueur/internal/errorhandler"
	"joueur/internal/gamemanager"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

// RunArgs are the expected arguments needed to actually run this Go Cadre
// game client.
type RunArgs struct {
	Server string
	Port   string

	GameName string

	AISettings   string
	PlayerName   string
	Password     string
	Session      string
	Index        string
	GameSettings string

	PrintIO bool
}

// SetupInterruptHandler creates a 'listener' on a new goroutine which will
// notify the program if it receives an interrupt from the OS. We then handle
// this by calling our clean up procedure and exiting the program.
func SetupInterruptHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n- Ctrl+C pressed in Terminal")
		errorhandler.StopHandlingErrors()
		client.Disconnect()
		os.Exit(0)
	}()
}

// Run actually runs the client, connecting to the game server, then
// playing with the AI and game objects.
func Run(args RunArgs) error {
	SetupInterruptHandler()

	splitServer := strings.Split(args.Server, ":")
	args.Server = splitServer[0]
	if len(splitServer) == 2 {
		args.Port = splitServer[1]
	}

	if args.Port == "" {
		args.Port = "3000"
	}

	if args.Server == "" {
		args.Server = "localhost"
	}

	client.Setup(args.PrintIO)

	playerIndex := -1
	if args.Index != "" {
		i, err := strconv.Atoi(args.Index)
		if err != nil {
			errorhandler.HandleError(
				errorhandler.InvalidArgs,
				err,
				"Cannot convert "+args.Index+" for a number for player index.",
			)
		}
		playerIndex = i
	}

	address := args.Server + ":" + args.Port
	color.Cyan("Connecting to: " + address)
	err := client.Connect(address)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.CouldNotConnect,
			err,
			"Error connecting to "+address,
		)
	}

	client.SendEventAlias(args.GameName)
	gameName := client.WaitForEventNamed()

	gameNamespace, err := games.Get(gameName)
	if gameNamespace == nil {
		err = errors.New("No GameNamespace for " + gameName)
	}

	if err != nil {
		return errorhandler.HandleError(
			errorhandler.GameNotFound,
			err,
			"Cannot find Game "+gameName,
		)
	}

	playerName := gameNamespace.PlayerName()
	if playerName == "" {
		playerName = "Go Player"
	}

	if args.PlayerName != "" {
		playerName = args.PlayerName
	}

	client.SendEventPlay(client.EventPlayData{
		ClientType:       "Go",
		GameName:         gameName,
		GameSettings:     args.GameSettings,
		Password:         args.Password,
		PlayerIndex:      playerIndex,
		PlayerName:       playerName,
		RequestedSession: args.Session,
	})
	lobbiedData := client.WaitForEventLobbied()
	color.Cyan("In lobby for game " + lobbiedData.GameName + " in session " + lobbiedData.GameSession)

	ourGameVersion := gameNamespace.Version()
	if lobbiedData.GameVersion != ourGameVersion {
		color.Yellow(
			`WARNING: Game versions do not match.
-> Your local game version is:     %s
-> Game Server's game version is:  %s

Version mismatch means that unexpected crashes may happen due to differing game structures!`,
			ourGameVersion[:8],
			lobbiedData.GameVersion[:8],
		)
	}

	gameManager := gamemanager.New(gameNamespace, args.AISettings, lobbiedData.Constants)
	if gameManager == nil {
		errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not create GameManager for game "+gameName),
		)
	}

	startData := client.WaitForEventStart()

	color.Green("Game is starting.")

	gameManager.Start(startData.PlayerID)

	// The client will now wait for order(s) asynchronously.
	// The process will exit when "over" is sent from the game server.

	overData := client.WaitForEventOver()

	won := gameManager.Player.Won()
	reason := gameManager.Player.ReasonWon()
	didWin := "I Won!"
	if !won {
		reason = gameManager.Player.ReasonLost()
		didWin = "I Lost :("
	}
	color.Green("Game is over. " + didWin + " because: " + reason)

	gameManager.AI.Ended(won, reason)

	if overData.Message != "" {
		color.Cyan(strings.Replace(overData.Message, "__HOSTNAME__", args.Server, 1))
	}

	client.Disconnect()

	return nil // should end the process with return code 0
}

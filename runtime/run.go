// Package base
// A pretty cool package
package runtime

import (
	"errors"
	"fmt"
	"joueur/base"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"
	"joueur/runtime/gamemanager"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

// hio
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

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
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

/**
 * Invoked to actually run the client, connecting to the game server, then
 * playing with the AI and game objects
 * @param args the command line args already parsed into a key/value dict
 */
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
		if err == nil {
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

	aiType := gameNamespace.AIType
	ai := reflect.New(aiType)
	if !ai.IsValid() {
		return errorhandler.HandleError(
			errorhandler.ReflectionFailed,
			errors.New("Could not create AI struct via reflect"),
		)
	}
	bai := ai.Elem().Interface().(base.InterfaceAI)

	playerName := bai.GetPlayerName()
	if playerName == "" {
		playerName = "Go Player"
	}

	if args.PlayerName != "" {
		playerName = args.PlayerName
	}

	client.SendEventPlay(client.EventPlay{
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

	if lobbiedData.GameVersion != (*gameNamespace).Version {
		color.Yellow(
			`WARNING: Game versions do not match.
-> Your local game version is:     %s
-> Game Server's game version is:  %s

Version mismatch means that unexpected crashes may happen due to differing game structures!`,
			(*gameNamespace).Version[:8],
			lobbiedData.GameVersion[:8],
		)
	}

	gameManager := gamemanager.New(&gamemanager.GameManager{
		ServerConstants: lobbiedData.Constants,
		GameNamespace:   gameNamespace,
		InterfaceAI:     &bai,
		ReflectAI:       &ai,
	}, args.AISettings)

	startData := client.WaitForEventStart()

	(*gameManager).Start(startData.PlayerID)

	color.Green("Game is starting.")

	/*
			// player is readonly but that's so competitors don't change it,
			// so cast to any here so we can set it
			(ai as any).player = game.gameObjects[startData.playerID];
			try {
				await ai.start();
				await ai.gameUpdated();
			} catch (err) {
				handleError(
					ErrorCode.AI_ERRORED,
					err,
					"AI errored when game initially started.",
				);
			}

			client.acceptOrders();

			// The client will now wait for order(s) asynchronously.
			// The process will exit when "end" is sent from the game server.
	*/

	client.Disconnect()

	fmt.Println("done!")

	return nil
}

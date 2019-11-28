// Package base
// A pretty cool package
package runtime

import (
	"errors"
	"fmt"
	"joueur/games"
	"joueur/runtime/client"
	"joueur/runtime/errorhandler"
	"reflect"
	"strings"
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
	Index        *int
	GameSettings string

	PrintIO bool
}

/**
 * Invoked to actually run the client, connecting to the game server, then
 * playing with the AI and game objects
 * @param args the command line args already parsed into a key/value dict
 */
func Run(args RunArgs) error {
	fmt.Println(args)

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

	address := args.Server + ":" + args.Port
	err := client.Connect(address)
	if err != nil {
		errorhandler.HandleError(
			errorhandler.CouldNotConnect,
			err,
			"Error connecting to "+address,
		)
	}

	// client.SendEventAlias(args.GameName)
	// gameName := client.WaitForEventNamed()
	gameName := "Chess"
	fmt.Println("gameName", string(gameName))

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

	gameType := (*gameNamespace).Types.Game
	val := reflect.New(gameType)
	fmt.Println("game type:", val)

	/*
			const gameNamespace: IGameNamespace | undefined = imported.namespace;

			if (!gameNamespace) {
				return handleError(
					ErrorCode.GAME_NOT_FOUND,
					`Game namespace for '${gameName}' is empty.`,
				);
			}

			if (!gameNamespace.AI
				|| !gameNamespace.Game
				|| !gameNamespace.GameObjectClasses
			) {
				return handleError(
					ErrorCode.GAME_NOT_FOUND,
					`Game namespace malformed for '${gameName}'.`,
				);
			}

			let game: BaseGame | undefined;
			try {
				game = new gameNamespace.Game();
			} catch (err) {
				return handleError(
					ErrorCode.REFLECTION_FAILED,
					err,
					`Error constructing the Game for '${gameName}'.`,
				);
			}

			let ai: BaseAI | undefined;
			try {
				ai = new gameNamespace.AI(game);
			} catch (err) {
				return handleError(
					ErrorCode.REFLECTION_FAILED,
					err,
					`Error constructing the AI for '${gameName}'.`,
				);
			}

			const gameManager = new BaseGameManager(
				game,
				gameNamespace.GameObjectClasses,
			);

			client.setup(ai, game, gameManager);

			setAISettings(ai, args.aiSettings || "");

			client.send("play", {
				clientType: "TypeScript",
				gameName,
				gameSettings: args.gameSettings,
				password: args.password,
				playerIndex: args.index,
				playerName: args.playerName
					|| ai.getName()
					|| "TypeScript Player",
				requestedSession: args.session,
			});

			const lobbyData = await client.waitForEvent("lobbied");

			if (lobbyData.gameVersion !== gameNamespace.gameVersion) {
				// tslint:disable-next-line:no-console
				console.warn(chalk.yellow(
		`WARNING: Game versions do not match.
		-> Your local game version is:	 ${gameNamespace.gameVersion.substr(0, 8)}
		-> Game Server's game version is:  ${lobbyData.gameVersion.substr(0, 8)}

		Version mismatch means that unexpected crashes may happen due to differing game structures!`));
			}

			// tslint:disable-next-line:no-console
			console.log(chalk.cyan(
				`In lobby for game '${lobbyData.gameName}' in`
				+ ` session '${lobbyData.gameSession}'.`,
			));

			gameManager.serverConstants = lobbyData.constants;

			// NOTE: if we try to use async/await syntax here it does NOT work
			// instead the order will execute before control is returned after this
			// waitForEvent("start") to resolve... for some reason...

			const startData = await client.waitForEvent("start");

			// tslint:disable-next-line:no-console
			console.log(chalk.green(`Game is starting.`));

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

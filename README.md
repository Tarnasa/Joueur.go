# Go Joueur Client

This is the client for the [Cadre] AI framework. It can play multiple
different games, though you will probably only be interested in one at a time.

In general, try to stay out of the `base/`, `runtime/`, and
`games/game_name/impl/` directories; they do most of the heavy lifting to play
on our game servers.

Each AI, and the game objects it manipulates are all in `games/game_name/`,
with your very own AI living in `games/game_name/ai.go` for you to make
smarter.

The interfaces for all the game objects also live in `games/game_name/`. The
implimentation logic for the game client ⮂ server communications live in the
`games/game_name/impl/` directory and are especially important to leave alone.

## How to Run

This client has been tested and confirmed to work on Linux, Mac and Windows
computers with the appropriate requirements installed.

### Requirements

The only requirement is [Go] version 1.13, which should install the `go`
binary on your system.

### Building

For all operating systems, you should only need to do the following:

```bash
go install
go build -o joueur
```

For any subsequent builds `go build -o joueur` should suffice. as `go install` should
download the base dependencies we need.

### Running

By default, the binary should output to `joueur`. Just run that binary with
the appropriate CLI args.

```bash
./joueur Game_name -s some.server.io
```

or

```bash
./run Game_name
```

## Dependencies

By default we include some dependencies in your `go.mod`. You are free to add
your own. However please don't the ones we included as otherwise your project
won't build!

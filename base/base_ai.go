package base

type BaseAI struct {
	game   BaseGame
	player BasePlayer
}

func (player BasePlayer) getPlayerName() string {
	return "Go Player"
}

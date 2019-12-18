package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/newtonian"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Newtonian.
type PlayerImpl struct {
	GameObjectImpl

	clientTypeImpl     string
	generatorTilesImpl []newtonian.Tile
	heatImpl           int64
	internSpawnImpl    int64
	lostImpl           bool
	managerSpawnImpl   int64
	nameImpl           string
	opponentImpl       newtonian.Player
	physicistSpawnImpl int64
	pressureImpl       int64
	reasonLostImpl     string
	reasonWonImpl      string
	spawnTilesImpl     []newtonian.Tile
	timeRemainingImpl  float64
	unitsImpl          []newtonian.Unit
	wonImpl            bool
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// GeneratorTiles returns every generator Tile owned by this Player.
// (arrayed from the outer edges inward, from top to bottom).
func (playerImpl *PlayerImpl) GeneratorTiles() []newtonian.Tile {
	return playerImpl.generatorTilesImpl
}

// Heat returns the amount of heat this Player has.
func (playerImpl *PlayerImpl) Heat() int64 {
	return playerImpl.heatImpl
}

// InternSpawn returns the time left till a intern spawns. (0 to
// spawnTime).
func (playerImpl *PlayerImpl) InternSpawn() int64 {
	return playerImpl.internSpawnImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// ManagerSpawn returns the time left till a manager spawns. (0 to
// spawnTime).
func (playerImpl *PlayerImpl) ManagerSpawn() int64 {
	return playerImpl.managerSpawnImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() newtonian.Player {
	return playerImpl.opponentImpl
}

// PhysicistSpawn returns the time left till a physicist spawns. (0 to
// spawnTime).
func (playerImpl *PlayerImpl) PhysicistSpawn() int64 {
	return playerImpl.physicistSpawnImpl
}

// Pressure returns the amount of pressure this Player has.
func (playerImpl *PlayerImpl) Pressure() int64 {
	return playerImpl.pressureImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// SpawnTiles returns all the tiles this Player's units can spawn on.
// (arrayed from the outer edges inward, from top to bottom).
func (playerImpl *PlayerImpl) SpawnTiles() []newtonian.Tile {
	return playerImpl.spawnTilesImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Units returns every Unit owned by this Player.
func (playerImpl *PlayerImpl) Units() []newtonian.Unit {
	return playerImpl.unitsImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.clientTypeImpl = ""
	playerImpl.generatorTilesImpl = []newtonian.Tile{}
	playerImpl.heatImpl = 0
	playerImpl.internSpawnImpl = 0
	playerImpl.lostImpl = true
	playerImpl.managerSpawnImpl = 0
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.physicistSpawnImpl = 0
	playerImpl.pressureImpl = 0
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.spawnTilesImpl = []newtonian.Tile{}
	playerImpl.timeRemainingImpl = 0
	playerImpl.unitsImpl = []newtonian.Unit{}
	playerImpl.wonImpl = true
}

// DeltaMerge merges the delta for a given attribute in Player.
func (playerImpl *PlayerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := playerImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	newtonianDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'newtonian.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "clientType":
		playerImpl.clientTypeImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "generatorTiles":
		playerImpl.generatorTilesImpl = newtonianDeltaMerge.ArrayOfTile(&playerImpl.generatorTilesImpl, delta)
		return true, nil
	case "heat":
		playerImpl.heatImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "internSpawn":
		playerImpl.internSpawnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = newtonianDeltaMerge.Boolean(delta)
		return true, nil
	case "managerSpawn":
		playerImpl.managerSpawnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = newtonianDeltaMerge.Player(delta)
		return true, nil
	case "physicistSpawn":
		playerImpl.physicistSpawnImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "pressure":
		playerImpl.pressureImpl = newtonianDeltaMerge.Int(delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = newtonianDeltaMerge.String(delta)
		return true, nil
	case "spawnTiles":
		playerImpl.spawnTilesImpl = newtonianDeltaMerge.ArrayOfTile(&playerImpl.spawnTilesImpl, delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = newtonianDeltaMerge.Float(delta)
		return true, nil
	case "units":
		playerImpl.unitsImpl = newtonianDeltaMerge.ArrayOfUnit(&playerImpl.unitsImpl, delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = newtonianDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

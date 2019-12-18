package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/anarchy"
)

// PlayerImpl is the struct that implements the container for Player
// instances in Anarchy.
type PlayerImpl struct {
	GameObjectImpl

	bribesRemainingImpl   int64
	buildingsImpl         []anarchy.Building
	clientTypeImpl        string
	fireDepartmentsImpl   []anarchy.FireDepartment
	headquartersImpl      anarchy.Warehouse
	lostImpl              bool
	nameImpl              string
	opponentImpl          anarchy.Player
	policeDepartmentsImpl []anarchy.PoliceDepartment
	reasonLostImpl        string
	reasonWonImpl         string
	timeRemainingImpl     float64
	warehousesImpl        []anarchy.Warehouse
	weatherStationsImpl   []anarchy.WeatherStation
	wonImpl               bool
}

// BribesRemaining returns how many bribes this player has remaining to use
// during their turn. Each action a Building does costs 1 bribe. Any unused
// bribes are lost at the end of the player's turn.
func (playerImpl *PlayerImpl) BribesRemaining() int64 {
	return playerImpl.bribesRemainingImpl
}

// Buildings returns all the buildings owned by this player.
func (playerImpl *PlayerImpl) Buildings() []anarchy.Building {
	return playerImpl.buildingsImpl
}

// ClientType returns what type of client this is, e.g. 'Python',
// 'JavaScript', or some other language. For potential data mining
// purposes.
func (playerImpl *PlayerImpl) ClientType() string {
	return playerImpl.clientTypeImpl
}

// FireDepartments returns all the FireDepartments owned by this player.
func (playerImpl *PlayerImpl) FireDepartments() []anarchy.FireDepartment {
	return playerImpl.fireDepartmentsImpl
}

// Headquarters returns the Warehouse that serves as this player's
// headquarters and has extra health. If this gets destroyed they lose.
func (playerImpl *PlayerImpl) Headquarters() anarchy.Warehouse {
	return playerImpl.headquartersImpl
}

// Lost returns if the player lost the game or not.
func (playerImpl *PlayerImpl) Lost() bool {
	return playerImpl.lostImpl
}

// Name returns the name of the player.
func (playerImpl *PlayerImpl) Name() string {
	return playerImpl.nameImpl
}

// Opponent returns this player's opponent in the game.
func (playerImpl *PlayerImpl) Opponent() anarchy.Player {
	return playerImpl.opponentImpl
}

// PoliceDepartments returns all the PoliceDepartments owned by this
// player.
func (playerImpl *PlayerImpl) PoliceDepartments() []anarchy.PoliceDepartment {
	return playerImpl.policeDepartmentsImpl
}

// ReasonLost returns the reason why the player lost the game.
func (playerImpl *PlayerImpl) ReasonLost() string {
	return playerImpl.reasonLostImpl
}

// ReasonWon returns the reason why the player won the game.
func (playerImpl *PlayerImpl) ReasonWon() string {
	return playerImpl.reasonWonImpl
}

// TimeRemaining returns the amount of time (in ns) remaining for this AI
// to send commands.
func (playerImpl *PlayerImpl) TimeRemaining() float64 {
	return playerImpl.timeRemainingImpl
}

// Warehouses returns all the warehouses owned by this player. Includes the
// Headquarters.
func (playerImpl *PlayerImpl) Warehouses() []anarchy.Warehouse {
	return playerImpl.warehousesImpl
}

// WeatherStations returns all the WeatherStations owned by this player.
func (playerImpl *PlayerImpl) WeatherStations() []anarchy.WeatherStation {
	return playerImpl.weatherStationsImpl
}

// Won returns if the player won the game or not.
func (playerImpl *PlayerImpl) Won() bool {
	return playerImpl.wonImpl
}

// InitImplDefaults initializes safe defaults for all fields in Player.
func (playerImpl *PlayerImpl) InitImplDefaults() {
	playerImpl.GameObjectImpl.InitImplDefaults()

	playerImpl.bribesRemainingImpl = 0
	playerImpl.buildingsImpl = []anarchy.Building{}
	playerImpl.clientTypeImpl = ""
	playerImpl.fireDepartmentsImpl = []anarchy.FireDepartment{}
	playerImpl.headquartersImpl = nil
	playerImpl.lostImpl = true
	playerImpl.nameImpl = ""
	playerImpl.opponentImpl = nil
	playerImpl.policeDepartmentsImpl = []anarchy.PoliceDepartment{}
	playerImpl.reasonLostImpl = ""
	playerImpl.reasonWonImpl = ""
	playerImpl.timeRemainingImpl = 0
	playerImpl.warehousesImpl = []anarchy.Warehouse{}
	playerImpl.weatherStationsImpl = []anarchy.WeatherStation{}
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

	anarchyDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'anarchy.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "bribesRemaining":
		playerImpl.bribesRemainingImpl = anarchyDeltaMerge.Int(delta)
		return true, nil
	case "buildings":
		playerImpl.buildingsImpl = anarchyDeltaMerge.ArrayOfBuilding(&playerImpl.buildingsImpl, delta)
		return true, nil
	case "clientType":
		playerImpl.clientTypeImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "fireDepartments":
		playerImpl.fireDepartmentsImpl = anarchyDeltaMerge.ArrayOfFireDepartment(&playerImpl.fireDepartmentsImpl, delta)
		return true, nil
	case "headquarters":
		playerImpl.headquartersImpl = anarchyDeltaMerge.Warehouse(delta)
		return true, nil
	case "lost":
		playerImpl.lostImpl = anarchyDeltaMerge.Boolean(delta)
		return true, nil
	case "name":
		playerImpl.nameImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "opponent":
		playerImpl.opponentImpl = anarchyDeltaMerge.Player(delta)
		return true, nil
	case "policeDepartments":
		playerImpl.policeDepartmentsImpl = anarchyDeltaMerge.ArrayOfPoliceDepartment(&playerImpl.policeDepartmentsImpl, delta)
		return true, nil
	case "reasonLost":
		playerImpl.reasonLostImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "reasonWon":
		playerImpl.reasonWonImpl = anarchyDeltaMerge.String(delta)
		return true, nil
	case "timeRemaining":
		playerImpl.timeRemainingImpl = anarchyDeltaMerge.Float(delta)
		return true, nil
	case "warehouses":
		playerImpl.warehousesImpl = anarchyDeltaMerge.ArrayOfWarehouse(&playerImpl.warehousesImpl, delta)
		return true, nil
	case "weatherStations":
		playerImpl.weatherStationsImpl = anarchyDeltaMerge.ArrayOfWeatherStation(&playerImpl.weatherStationsImpl, delta)
		return true, nil
	case "won":
		playerImpl.wonImpl = anarchyDeltaMerge.Boolean(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

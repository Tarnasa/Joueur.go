package impl
import (
	"errors"
	"joueur/base"
	"joueur/games/checkers"
)

// CheckerImpl is the struct that implements the container for Checker
// instances in Checkers.
type CheckerImpl struct {
	GameObjectImpl

	kingedImpl bool
	ownerImpl  checkers.Player
	xImpl      int64
	yImpl      int64
}

// Kinged returns if the checker has been kinged and can move backwards.
func (checkerImpl *CheckerImpl) Kinged() bool {
	return checkerImpl.kingedImpl
}

// Owner returns the player that controls this Checker.
func (checkerImpl *CheckerImpl) Owner() checkers.Player {
	return checkerImpl.ownerImpl
}

// X returns the x coordinate of the checker.
func (checkerImpl *CheckerImpl) X() int64 {
	return checkerImpl.xImpl
}

// Y returns the y coordinate of the checker.
func (checkerImpl *CheckerImpl) Y() int64 {
	return checkerImpl.yImpl
}

// IsMine runs logic that returns if the checker is owned by your player or
// not.
func (checkerImpl *CheckerImpl) IsMine() bool {
	return checkerImpl.RunOnServer("isMine", map[string]interface{}{
	}).(bool)
}

// Move runs logic that moves the checker from its current location to the
// given (x, y).
func (checkerImpl *CheckerImpl) Move(x int64, y int64) checkers.Checker {
	return checkerImpl.RunOnServer("move", map[string]interface{}{
		"x": x,
		"y": y,
	}).(checkers.Checker)
}

// InitImplDefaults initializes safe defaults for all fields in Checker.
func (checkerImpl *CheckerImpl) InitImplDefaults() {
	checkerImpl.GameObjectImpl.InitImplDefaults()

	checkerImpl.kingedImpl = true
	checkerImpl.ownerImpl = nil
	checkerImpl.xImpl = 0
	checkerImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Checker.
func (checkerImpl *CheckerImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := checkerImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	checkersDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
			"'checkers.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "kinged":
		checkerImpl.kingedImpl = checkersDeltaMerge.Boolean(delta)
		return true, nil
	case "owner":
		checkerImpl.ownerImpl = checkersDeltaMerge.Player(delta)
		return true, nil
	case "x":
		checkerImpl.xImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	case "y":
		checkerImpl.yImpl = checkersDeltaMerge.Int(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

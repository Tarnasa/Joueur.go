package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/stardash"
)

// BodyImpl is the struct that implements the container for Body instances
// in Stardash.
type BodyImpl struct {
	GameObjectImpl

	amountImpl       int64
	bodyTypeImpl     string
	materialTypeImpl string
	ownerImpl        stardash.Player
	radiusImpl       float64
	xImpl            float64
	yImpl            float64
}

// Amount returns the amount of material the object has, or energy if it is
// a planet.
func (bodyImpl *BodyImpl) Amount() int64 {
	return bodyImpl.amountImpl
}

// BodyType returns the type of celestial body it is. Either 'planet',
// 'asteroid', or 'sun'.
//
// Literal Values: "planet", "asteroid", or "sun"
func (bodyImpl *BodyImpl) BodyType() string {
	return bodyImpl.bodyTypeImpl
}

// MaterialType returns the type of material the celestial body has. Either
// 'none', 'genarium', 'rarium', 'legendarium', or 'mythicite'.
//
// Literal Values: "none", "genarium", "rarium", "legendarium", or
// "mythicite"
func (bodyImpl *BodyImpl) MaterialType() string {
	return bodyImpl.materialTypeImpl
}

// Owner returns the Player that owns and can control this Body.
//
// Value can be returned as a nil pointer.
func (bodyImpl *BodyImpl) Owner() stardash.Player {
	return bodyImpl.ownerImpl
}

// Radius returns the radius of the circle that this body takes up.
func (bodyImpl *BodyImpl) Radius() float64 {
	return bodyImpl.radiusImpl
}

// X returns the x value this celestial body is on.
func (bodyImpl *BodyImpl) X() float64 {
	return bodyImpl.xImpl
}

// Y returns the y value this celestial body is on.
func (bodyImpl *BodyImpl) Y() float64 {
	return bodyImpl.yImpl
}

// NextX runs logic that the x value of this body a number of turns from
// now. (0-how many you want).
func (bodyImpl *BodyImpl) NextX(num int64) int64 {
	return bodyImpl.RunOnServer("nextX", map[string]interface{}{
		"num": num,
	}).(int64)
}

// NextY runs logic that the x value of this body a number of turns from
// now. (0-how many you want).
func (bodyImpl *BodyImpl) NextY(num int64) int64 {
	return bodyImpl.RunOnServer("nextY", map[string]interface{}{
		"num": num,
	}).(int64)
}

// Spawn runs logic that spawn a unit on some value of this celestial body.
func (bodyImpl *BodyImpl) Spawn(x float64, y float64, title string) bool {
	return bodyImpl.RunOnServer("spawn", map[string]interface{}{
		"x":     x,
		"y":     y,
		"title": title,
	}).(bool)
}

// InitImplDefaults initializes safe defaults for all fields in Body.
func (bodyImpl *BodyImpl) InitImplDefaults() {
	bodyImpl.GameObjectImpl.InitImplDefaults()

	bodyImpl.amountImpl = 0
	bodyImpl.bodyTypeImpl = ""
	bodyImpl.materialTypeImpl = ""
	bodyImpl.ownerImpl = nil
	bodyImpl.radiusImpl = 0
	bodyImpl.xImpl = 0
	bodyImpl.yImpl = 0
}

// DeltaMerge merges the delta for a given attribute in Body.
func (bodyImpl *BodyImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := bodyImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	stardashDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'stardash.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "amount":
		bodyImpl.amountImpl = stardashDeltaMerge.Int(delta)
		return true, nil
	case "bodyType":
		bodyImpl.bodyTypeImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "materialType":
		bodyImpl.materialTypeImpl = stardashDeltaMerge.String(delta)
		return true, nil
	case "owner":
		bodyImpl.ownerImpl = stardashDeltaMerge.Player(delta)
		return true, nil
	case "radius":
		bodyImpl.radiusImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "x":
		bodyImpl.xImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	case "y":
		bodyImpl.yImpl = stardashDeltaMerge.Float(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

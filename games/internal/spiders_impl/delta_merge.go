package impl

import (
	"joueur/base"
	"joueur/games/spiders"
)

// DeltaMerge is the set of functions that can delta merge a
// Spiders game.
type DeltaMerge interface {
	base.DeltaMerge

	BroodMother(interface{}) spiders.BroodMother
	Cutter(interface{}) spiders.Cutter
	GameObject(interface{}) spiders.GameObject
	Nest(interface{}) spiders.Nest
	Player(interface{}) spiders.Player
	Spider(interface{}) spiders.Spider
	Spiderling(interface{}) spiders.Spiderling
	Spitter(interface{}) spiders.Spitter
	Weaver(interface{}) spiders.Weaver
	Web(interface{}) spiders.Web

	ArrayOfNest(*[]spiders.Nest, interface{}) []spiders.Nest
	ArrayOfPlayer(*[]spiders.Player, interface{}) []spiders.Player
	ArrayOfSpider(*[]spiders.Spider, interface{}) []spiders.Spider
	ArrayOfSpiderling(*[]spiders.Spiderling, interface{}) []spiders.Spiderling
	ArrayOfString(*[]string, interface{}) []string
	ArrayOfWeb(*[]spiders.Web, interface{}) []spiders.Web
	MapOfStringToGameObject(*map[string]spiders.GameObject, interface{}) map[string]spiders.GameObject
}

// DeltaMergeImpl is the implimentation struct for the DeltaMerge interface.
type DeltaMergeImpl struct {
	base.DeltaMergeImpl
}

// -- Spiders Game Object Deltas -- \\

// BroodMother attempts to return the instance of BroodMother
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) BroodMother(delta interface{}) spiders.BroodMother {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asBroodMother, isBroodMother := baseGameObject.(spiders.BroodMother)
	if !isBroodMother {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.BroodMother", delta)
	}

	return asBroodMother
}

// Cutter attempts to return the instance of Cutter
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Cutter(delta interface{}) spiders.Cutter {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asCutter, isCutter := baseGameObject.(spiders.Cutter)
	if !isCutter {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Cutter", delta)
	}

	return asCutter
}

// GameObject attempts to return the instance of GameObject
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) GameObject(delta interface{}) spiders.GameObject {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asGameObject, isGameObject := baseGameObject.(spiders.GameObject)
	if !isGameObject {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.GameObject", delta)
	}

	return asGameObject
}

// Nest attempts to return the instance of Nest
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Nest(delta interface{}) spiders.Nest {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asNest, isNest := baseGameObject.(spiders.Nest)
	if !isNest {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Nest", delta)
	}

	return asNest
}

// Player attempts to return the instance of Player
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Player(delta interface{}) spiders.Player {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asPlayer, isPlayer := baseGameObject.(spiders.Player)
	if !isPlayer {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Player", delta)
	}

	return asPlayer
}

// Spider attempts to return the instance of Spider
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Spider(delta interface{}) spiders.Spider {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asSpider, isSpider := baseGameObject.(spiders.Spider)
	if !isSpider {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Spider", delta)
	}

	return asSpider
}

// Spiderling attempts to return the instance of Spiderling
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Spiderling(delta interface{}) spiders.Spiderling {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asSpiderling, isSpiderling := baseGameObject.(spiders.Spiderling)
	if !isSpiderling {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Spiderling", delta)
	}

	return asSpiderling
}

// Spitter attempts to return the instance of Spitter
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Spitter(delta interface{}) spiders.Spitter {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asSpitter, isSpitter := baseGameObject.(spiders.Spitter)
	if !isSpitter {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Spitter", delta)
	}

	return asSpitter
}

// Weaver attempts to return the instance of Weaver
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Weaver(delta interface{}) spiders.Weaver {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asWeaver, isWeaver := baseGameObject.(spiders.Weaver)
	if !isWeaver {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Weaver", delta)
	}

	return asWeaver
}

// Web attempts to return the instance of Web
// for the given Delta.
func (deltaMergeImpl *DeltaMergeImpl) Web(delta interface{}) spiders.Web {
	baseGameObject := (*deltaMergeImpl).BaseGameObject(delta)
	if baseGameObject == nil {
		return nil
	}

	asWeb, isWeb := baseGameObject.(spiders.Web)
	if !isWeb {
		(*deltaMergeImpl).CannotConvertDeltaTo("spiders.Web", delta)
	}

	return asWeb
}

// -- Deep Deltas -- \\

// ArrayOfNest delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfNest(state *[]spiders.Nest, delta interface{}) []spiders.Nest {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]spiders.Nest, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Nest(deltaValue)
	}
	return newArray
}

// ArrayOfPlayer delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfPlayer(state *[]spiders.Player, delta interface{}) []spiders.Player {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]spiders.Player, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Player(deltaValue)
	}
	return newArray
}

// ArrayOfSpider delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfSpider(state *[]spiders.Spider, delta interface{}) []spiders.Spider {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]spiders.Spider, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Spider(deltaValue)
	}
	return newArray
}

// ArrayOfSpiderling delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfSpiderling(state *[]spiders.Spiderling, delta interface{}) []spiders.Spiderling {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]spiders.Spiderling, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Spiderling(deltaValue)
	}
	return newArray
}

// ArrayOfString delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfString(state *[]string, delta interface{}) []string {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]string, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.String(deltaValue)
	}
	return newArray
}

// ArrayOfWeb delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) ArrayOfWeb(state *[]spiders.Web, delta interface{}) []spiders.Web {
	deltaList, listLength := (*deltaMergeImpl).ToDeltaArray(delta)
	newArray := make([]spiders.Web, listLength) // resize array with new copy
	copy(newArray, *state)
	for deltaIndex, deltaValue := range deltaList {
		newArray[deltaIndex] = deltaMergeImpl.Web(deltaValue)
	}
	return newArray
}

// MapOfStringToGameObject delta attempts to merge
// deep structures of this type.
func (deltaMergeImpl *DeltaMergeImpl) MapOfStringToGameObject(state *map[string]spiders.GameObject, delta interface{}) map[string]spiders.GameObject {
	deltaMap := (*deltaMergeImpl).ToDeltaMap(delta)
	for deltaKey, deltaValue := range deltaMap {
		if (*deltaMergeImpl).IsDeltaRemoved(deltaValue) {
			delete(*state, deltaKey)
		} else {
			(*state)[deltaKey] = deltaMergeImpl.GameObject(deltaValue)
		}
	}
	return *state
}

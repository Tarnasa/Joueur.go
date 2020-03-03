package gamemanager

import (
	"errors"
	"fmt"
	"joueur/internal/errorhandler"
)

// applyDeltaState take the root (game) state and delta updates all the
// structs within the game it is managing.
func (gameManager *GameManager) applyDeltaState(delta map[string]interface{}) {
	gameObjectsRaw, gameObjectsExist := delta["gameObjects"]
	if gameObjectsExist {
		gameObjectsDeltaRaw, gameObjectsAreMapRaw := gameObjectsRaw.(map[string]interface{})
		if !gameObjectsAreMapRaw {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("cannot merge delta when 'gameObjects' property is not a map: "+fmt.Sprintf("%v", gameObjectsRaw)),
			)
		}

		gameObjectsDeltas := make(map[string]map[string]interface{})
		for id, rawGameObjectDelta := range gameObjectsDeltaRaw {
			gameObjectsDeltas[id] = gameManager.deltaMerge.ToDeltaMap(rawGameObjectDelta)
		}

		// game objects delta looks solid, create all new game objects we have not seen yet before meging
		gameManager.initGameObjects(&gameObjectsDeltas)

		// now all new game objects should be initialize so we can delta merge as normal
		for id, gameObjectDelta := range gameObjectsDeltas {
			gameObject, gameObjectExists := gameManager.gameObjects[id]
			if !gameObjectExists {
				errorhandler.HandleError(
					errorhandler.DeltaMergeFailure,
					errors.New("cannot merge delta state of game object #"+id+" with no game object for given id"),
				)
			}
			for gameObjectAttribute, gameObjectAttributeDelta := range gameObjectDelta {
				_, err := gameObject.DeltaMerge(gameManager.deltaMerge, gameObjectAttribute, gameObjectAttributeDelta)
				if err != nil {
					errorhandler.HandleError(
						errorhandler.DeltaMergeFailure,
						err,
						"Error merging '"+gameObjectAttribute+"' in game object #"+id,
					)
				}
			}
		}
	}

	// now all new game objects should be delta merged, only thing remaining is the game itself's delta state
	for gameAttribute, gameAttributeDelta := range delta {
		if gameAttribute == "gameObjects" {
			continue // we already updated gameObject above
		}

		gameManager.Game.DeltaMerge(gameManager.deltaMerge, gameAttribute, gameAttributeDelta)
	}
}

// initGameObjects initializes all new game objects found in a gameObjects
// delta. This should be done before any other delta merges, so that if a
// reference to a new game object is found, it will exist to be hooked up.
func (gameManager *GameManager) initGameObjects(gameObjectsDeltas *map[string]map[string]interface{}) {
	for id, gameObjectDelta := range *gameObjectsDeltas {
		_, gameObjectAlreadyExists := gameManager.gameObjects[id]
		if gameObjectAlreadyExists {
			continue
		}

		rawGameObjectName, rawNameOk := gameObjectDelta["gameObjectName"]
		gameObjectName, nameIsString := rawGameObjectName.(string)
		if !rawNameOk || !nameIsString || gameObjectName == "" {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("field 'gameObjectName' not a string on new game object #"+id),
			)
		}

		newGameObject, creationError := gameManager.GameNamespace.CreateGameObject(gameObjectName)
		if creationError != nil {
			errorhandler.HandleError(
				errorhandler.DeltaMergeFailure,
				errors.New("creation error on new game object "+gameObjectName+" #"+id),
			)
		}

		gameManager.gameObjects[id] = newGameObject
		gameManager.Game.AddGameObject(id, newGameObject)
	}
}

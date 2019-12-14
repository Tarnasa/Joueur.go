package base

// GameNamespace is the base interface all games must implement to be able to
// play via our soft reflection
type GameNamespace interface {
	Name() string
	Version() string
	PlayerName() string
	CreateAI() (AI, *AIImpl)
	CreateDeltaMerge(*DeltaMergeImpl) DeltaMerge
	CreateGame() DeltaMergeableGame
	CreateGameObject(string) (DeltaMergeableGameObject, error)
	OrderAI(AI, string, []interface{}) (interface{}, error)
}

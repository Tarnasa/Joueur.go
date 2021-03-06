package base

// DeltaMergeable is an interface to something in the game that will receive and merge delta states
type DeltaMergeable interface {
	DeltaMerge(DeltaMerge, string, interface{}) (bool, error)
	InitImplDefaults()
}

// DeltaMergeableImpl is the implimentation of a struct that can be
// delta merged
type DeltaMergeableImpl struct{}

// DeltaMerge will merge the given delta into itself
func (DeltaMergeableImpl) DeltaMerge(deltaMerge DeltaMerge, attribute string, delta interface{}) (bool, error) {
	// this is up to the game impl to actually implement so we can avoid golang reflect
	return false, nil
}

// InitImplDefaults should initialize any defaults for delta mergable fields.
func InitImplDefaults() {}

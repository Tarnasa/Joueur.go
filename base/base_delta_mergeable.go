package base

type DeltaMergeable interface {
	DeltaMerge(DeltaMerge, interface{})
}

// DeltaMergeableImpl is the implimentation of a struct that can be
// delta merged
type DeltaMergeableImpl struct {
	InternalDataMap map[string]interface{}
}

func (*DeltaMergeableImpl) DeltaMerge(deltaMerge DeltaMerge, delta interface{}) {
	// this is up to the game impl to actually impliment so we can avoid golang reflect
}

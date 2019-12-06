package base

// DeltaMergeableImpl is the implimentation of a struct that can be
// delta merged
type DeltaMergeableImpl struct {
	InternalDataMap map[string]interface{}
}

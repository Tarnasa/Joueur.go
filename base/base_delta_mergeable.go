package base

type DeltaMerge interface {
	String(*string, interface{})
	Int(*int64, interface{})
	Float(*float64, interface{})
	Boolean(*bool, interface{})

	Array(*[]interface{}, interface{}, func(...interface{}))
	Map(*map[interface{}]interface{}, interface{}, func(...interface{}), func(...interface{}))
	GameObject(*GameObject, interface{}, func(GameObject) bool)
}

// DeltaMergeableImpl is the implimentation of a struct that can be
// delta merged
type DeltaMergeableImpl struct {
	InternalDataMap map[string]interface{}
}

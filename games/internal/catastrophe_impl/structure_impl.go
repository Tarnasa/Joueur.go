package impl

import (
	"errors"
	"joueur/base"
	"joueur/games/catastrophe"
)

// StructureImpl is the struct that implements the container for Structure
// instances in Catastrophe.
type StructureImpl struct {
	GameObjectImpl

	effectRadiusImpl int64
	materialsImpl    int64
	ownerImpl        catastrophe.Player
	tileImpl         catastrophe.Tile
	typeImpl         string
}

// EffectRadius returns the range of this Structure's effect. For example,
// a radius of 1 means this Structure affects a 3x3 square centered on this
// Structure.
func (structureImpl *StructureImpl) EffectRadius() int64 {
	return structureImpl.effectRadiusImpl
}

// Materials returns the number of materials in this Structure. Once this
// number reaches 0, this Structure is destroyed.
func (structureImpl *StructureImpl) Materials() int64 {
	return structureImpl.materialsImpl
}

// Owner returns the owner of this Structure if any, otherwise nil.
//
// Value can be returned as a nil pointer.
func (structureImpl *StructureImpl) Owner() catastrophe.Player {
	return structureImpl.ownerImpl
}

// Tile returns the Tile this Structure is on.
//
// Value can be returned as a nil pointer.
func (structureImpl *StructureImpl) Tile() catastrophe.Tile {
	return structureImpl.tileImpl
}

// Type returns the type of Structure this is ('shelter', 'monument',
// 'wall', 'road', 'neutral').
//
// Literal Values: "neutral", "shelter", "monument", "wall", or "road"
func (structureImpl *StructureImpl) Type() string {
	return structureImpl.typeImpl
}

// InitImplDefaults initializes safe defaults for all fields in Structure.
func (structureImpl *StructureImpl) InitImplDefaults() {
	structureImpl.GameObjectImpl.InitImplDefaults()

	structureImpl.effectRadiusImpl = 0
	structureImpl.materialsImpl = 0
	structureImpl.ownerImpl = nil
	structureImpl.tileImpl = nil
	structureImpl.typeImpl = ""
}

// DeltaMerge merges the delta for a given attribute in Structure.
func (structureImpl *StructureImpl) DeltaMerge(
	deltaMerge base.DeltaMerge,
	attribute string,
	delta interface{},
) (bool, error) {
	merged, err := structureImpl.GameObjectImpl.DeltaMerge(
		deltaMerge,
		attribute,
		delta,
	)
	if merged || err != nil {
		return merged, err
	}

	catastropheDeltaMerge, ok := deltaMerge.(DeltaMerge)
	if !ok {
		return false, errors.New(
			"deltaMerge is not the expected type of: " +
				"'catastrophe.impl.DeltaMerge'",
		)
	}

	switch attribute {
	case "effectRadius":
		structureImpl.effectRadiusImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "materials":
		structureImpl.materialsImpl = catastropheDeltaMerge.Int(delta)
		return true, nil
	case "owner":
		structureImpl.ownerImpl = catastropheDeltaMerge.Player(delta)
		return true, nil
	case "tile":
		structureImpl.tileImpl = catastropheDeltaMerge.Tile(delta)
		return true, nil
	case "type":
		structureImpl.typeImpl = catastropheDeltaMerge.String(delta)
		return true, nil
	}

	return false, nil // no errors in delta merging
}

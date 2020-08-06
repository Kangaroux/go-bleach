package bleach

import (
	"strings"
)

type trimType int

const (
	trimLeft trimType = iota
	trimRight
	trimBoth
)

// TrimMutator is a mutator which trims characters contained in a cutset from the beginning and/or
// end of an input.
type TrimMutator struct {
	cutset   string
	trimType trimType
}

var _ Mutator = (*TrimMutator)(nil)

// Mutate trims any values contained in the cutset from the beginning and/or end of the input value.
func (m *TrimMutator) Mutate(in interface{}) interface{} {
	val, _ := in.(string)

	if m.trimType == trimLeft {
		return strings.TrimLeft(val, m.cutset)
	} else if m.trimType == trimRight {
		return strings.TrimRight(val, m.cutset)
	} else {
		return strings.Trim(val, m.cutset)
	}
}

// Trim returns a new TrimMutator which will trim any characters contained in cutset from the beginning
// and end of a string. This works the same way as strings.Trim()
func Trim(cutset string) *TrimMutator {
	return &TrimMutator{
		cutset:   cutset,
		trimType: trimBoth,
	}
}

// TrimLeft returns a new TrimMutator which trims any characters contained in cutset from the beginning
// of a string. This works the same way as strings.TrimLeft()
func TrimLeft(cutset string) *TrimMutator {
	return &TrimMutator{
		cutset:   cutset,
		trimType: trimLeft,
	}
}

// TrimRight returns a new TrimMutator which trims any characters contained in cutset from the end
// of a string. This works the same way as strings.TrimRight()
func TrimRight(cutset string) *TrimMutator {
	return &TrimMutator{
		cutset:   cutset,
		trimType: trimRight,
	}
}

// TrimSpace returns a new TrimMutator which will trim any whitespace characters. This works the
// same way as strings.TrimSpace()
func TrimSpace() *TrimMutator {
	return &TrimMutator{
		cutset:   strings.TrimSpace("\t\n\v\f\r \u0085\u00A0"),
		trimType: trimBoth,
	}
}

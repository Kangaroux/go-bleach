package bleach

import (
	"strings"
)

// Trim returns a new Mutator which will trim any characters contained in cutset from the beginning
// and end of a string.
func Trim(cutset string) Mutator {
	fn := func(in interface{}) interface{} {
		val, _ := in.(string)
		return strings.Trim(val, cutset)
	}

	return MutatorFunc(fn)
}

// TrimLeft returns a new Mutator which trims any characters contained in cutset from the beginning
// of a string.
func TrimLeft(cutset string) Mutator {
	fn := func(in interface{}) interface{} {
		val, _ := in.(string)
		return strings.TrimLeft(val, cutset)
	}

	return MutatorFunc(fn)
}

// TrimRight returns a new Mutator which trims any characters contained in cutset from the end
// of a string.
func TrimRight(cutset string) Mutator {
	fn := func(in interface{}) interface{} {
		val, _ := in.(string)
		return strings.TrimRight(val, cutset)
	}

	return MutatorFunc(fn)
}

// TrimSpace returns a new Mutator which will trim any whitespace characters.
func TrimSpace() Mutator {
	fn := func(in interface{}) interface{} {
		val, _ := in.(string)
		return strings.TrimSpace(val)
	}

	return MutatorFunc(fn)
}

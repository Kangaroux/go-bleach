package bleach

// CheckFunc is a function which is used to check and validate a value. If the value fails the check,
// the CheckFunc should return an error with the reason why.
type CheckFunc func(interface{}) error

// Checker is an interface for objects which can be used to check values.
type Checker interface {
	Check(interface{}) error
}

// MutatorFunc is a function which accepts a value, modifies it, and then returns it.
type MutatorFunc func(interface{}) interface{}

// Mutator is an interface for objects which can mutate a value.
type Mutator interface {
	Mutate(interface{}) interface{}
}

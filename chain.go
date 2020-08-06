package bleach

// Chain represents a series of mutations/checks.
type Chain struct {
	steps []interface{}
}

// Run passes an input value through the chain and returns the mutated value and any accumulated
// errors. The returned value is nil if there are any errors. Run will normally complete the full
// chain even if it encounters errors. The exception to this is if a Checker returns a CancelError.
// A CancelError signals the chain to immediately stop and return.
func (c *Chain) Run(in interface{}) (interface{}, []error) {
	errors := []error{}
	val := in

	// Perform each step in the chain. Checks and mutations are not mutually exclusive. A step may
	// be both a Checker and a Mutator. In that scenario, the check is performed before the mutation.
	for _, step := range c.steps {
		// Perform the check.
		if checker, ok := step.(Checker); ok {
			err := checker.Check(val)

			if err != nil {
				errors = append(errors, err)

				if _, ok := err.(CancelError); ok {
					break
				}
			}
		}

		// Perform the mutation.
		if mutator, ok := step.(Mutator); ok {
			val = mutator.Mutate(val)
		}
	}

	if len(errors) == 0 {
		return val, nil
	}

	return nil, errors
}

// NewChain creates a new Chain of steps. Each step is either a Checker, a Mutator, or both. In the
// case that a step is both a Checker and a Mutator, the check is performed before the mutation.
// Panics if a step is neither a Checker nor a Mutator.
func NewChain(steps ...interface{}) *Chain {
	chain := &Chain{
		steps: make([]interface{}, len(steps)),
	}

	// Verify each step is a Checker and/or a Mutator, otherwise panic.
	for _, step := range steps {
		switch step.(type) {
		case Checker, Mutator:
			chain.steps = append(chain.steps, step)
		default:
			panic("expected chain steps to be a Checker and/or Mutator")
		}
	}

	return chain
}

package bleach

// Chain represents a series of mutations/checks.
type Chain struct {
	steps []interface{}
}

// NewChain creates a new Chain that is composed of the given steps.
func NewChain(steps ...interface{}) *Chain {
	chain := &Chain{
		steps: make([]interface{}, len(steps)),
	}

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

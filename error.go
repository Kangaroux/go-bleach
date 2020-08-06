package bleach

import "errors"

// CancelError is a special kind of error that can be returned by a Checker. If a Checker returns a
// CancelError and the Checker is part of a chain, the chain will immediately stop and return.
type CancelError error

// NewCancelError creates a CancelError using the provided error message.
func NewCancelError(msg string) CancelError {
	return (CancelError)(errors.New(msg))
}

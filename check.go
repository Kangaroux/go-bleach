package validation

import (
	"errors"
	"fmt"
)

type CheckFunc func(interface{}) error

type Checker interface {
	Check(interface{}) error
}

type LengthCheck struct {
	Min     int
	Max     int
	Message error
}

func (c *LengthCheck) Check(val interface{}) error {
	length := len(val.(string))

	if (c.Min > 0 && length < c.Min) || (c.Max > 0 && length > c.Max) {
		return c.Message
	}

	return nil
}

func (c *LengthCheck) Throws(msg string) *LengthCheck {
	c.Message = errors.New(msg)
	return c
}

func Length(min, max int) *LengthCheck {
	check := &LengthCheck{
		Min: min,
		Max: max,
	}

	if min > 0 && max > 0 {
		check.Message = fmt.Errorf("length must be between %d and %d characters", min, max)
	} else if min > 0 && max == 0 {
		check.Message = fmt.Errorf("cannot be shorter than %d characters", min)
	} else if min == 0 && max > 0 {
		check.Message = fmt.Errorf("cannot be longer than %d characters", max)
	}

	return check
}

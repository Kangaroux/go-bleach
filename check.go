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
	min     int
	max     int
	message error
}

func (c *LengthCheck) Check(val interface{}) error {
	length := len(val.(string))

	if (c.min > 0 && length < c.min) || (c.max > 0 && length > c.max) {
		return c.message
	}

	return nil
}

func (c *LengthCheck) Throws(msg string) *LengthCheck {
	c.message = errors.New(msg)
	return c
}

func Length(min, max int) *LengthCheck {
	check := &LengthCheck{
		min: min,
		max: max,
	}

	if min > 0 && max > 0 {
		check.message = fmt.Errorf(i18n.get(CheckLengthOutOfRange), min, max)
	} else if min > 0 && max == 0 {
		check.message = fmt.Errorf(i18n.get(CheckLengthTooShort), min)
	} else if min == 0 && max > 0 {
		check.message = fmt.Errorf(i18n.get(CheckLengthTooLong), max)
	}

	return check
}

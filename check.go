package bleach

import (
	"errors"
	"fmt"
)

// CheckFunc is a function which is used to check and validate a value. If the value fails the check,
// the CheckFunc should return an error with the reason why.
type CheckFunc func(interface{}) error

// Checker is an interface for objects which can be used to check values.
type Checker interface {
	Check(interface{}) error
}

// LengthChecker checks a string's length.
type LengthChecker struct {
	min     int
	max     int
	message error
}

var _ Checker = (*LengthChecker)(nil)

// Check checks a string's length.
func (c *LengthChecker) Check(in interface{}) error {
	val, _ := in.(string)
	length := len(val)

	if (c.min > 0 && length < c.min) || (c.max > 0 && length > c.max) {
		return c.message
	}

	return nil
}

// Throws is a chaining method for setting a custom error message for a LengthCheck.
//
//		Length(0, 10).Throws("too long!")
func (c *LengthChecker) Throws(msg string) *LengthChecker {
	c.message = errors.New(msg)
	return c
}

// Length returns a new LengthCheck. The min and max parameters correspond to the minimum and maximum
// lengths allowed. If max is zero then only the minimum length is checked. Length panics if min or max
// are negative.
//
// 		// Length must be between [5, 10] characters.
// 		Length(5, 10)
//
// 		// Length cannot be more than 10 characters.
//		Length(0, 10)
//
// 		// Length must be at least 5 characters.
//		Length(5, 0)
func Length(min int, max int) *LengthChecker {
	if min < 0 || max < 0 {
		panic("min and max cannot be negative")
	}

	check := &LengthChecker{
		min: min,
		max: max,
	}

	if min > 0 && max > 0 {
		check.message = fmt.Errorf(i18n.get(i18nCheckLengthOutOfRange), min, max)
	} else if min > 0 && max == 0 {
		check.message = fmt.Errorf(i18n.get(i18nCheckLengthTooShort), min)
	} else if min == 0 && max > 0 {
		check.message = fmt.Errorf(i18n.get(i18nCheckLengthTooLong), max)
	}

	return check
}

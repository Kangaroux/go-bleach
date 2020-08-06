package bleach

import (
	"errors"
	"fmt"
	"reflect"
)

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

// Throws is a chaining method for setting a custom error message for a LengthChecker.
//
//		Length(0, 10).Throws("too long!")
func (c *LengthChecker) Throws(msg string) *LengthChecker {
	c.message = errors.New(msg)
	return c
}

// Length returns a new LengthChecker. The min and max parameters correspond to the minimum and maximum
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

// TypeChecker checks that an input value can be converted to a specified type.
type TypeChecker struct {
	t       reflect.Type
	strict  bool
	message error
}

// Check checks that the input value can be converted to the specified type.
func (c *TypeChecker) Check(in interface{}) error {
	if c.strict {
		if !c.t.AssignableTo(reflect.TypeOf(in)) {
			return c.message
		}
	} else {
		if !c.t.ConvertibleTo(reflect.TypeOf(in)) {
			return c.message
		}
	}

	return nil
}

// Throws is a chaining method for setting a custom error message for a TypeChecker.
//
//		IsType(reflect.TypeOf("")).Throws("must be a string")
func (c *TypeChecker) Throws(msg string) *TypeChecker {
	c.message = errors.New(msg)
	return c
}

// IsType returns a new TypeChecker. The type parameter t describes what type is allowed for an input.
// This form of type checking allows values that are convertible to the provided type.
func IsType(t reflect.Type) *TypeChecker {
	return &TypeChecker{
		t:       t,
		strict:  false,
		message: fmt.Errorf(i18n.get(i18nCheckTypeNotConvertible), t.Kind),
	}
}

// IsTypeStrict returns a new TypeChecker. The type parameter t describes what type is allowed
// for an input. Strict type checking only allows input values that match the type exactly.
func IsTypeStrict(t reflect.Type) *TypeChecker {
	return &TypeChecker{
		t:       t,
		strict:  true,
		message: fmt.Errorf(i18n.get(i18nCheckTypeStrictBadType), t.Kind),
	}
}

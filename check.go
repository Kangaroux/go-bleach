package bleach

import (
	"errors"
	"fmt"
	"reflect"
)

// lengthChecker checks a string's length.
type lengthChecker struct {
	min     int
	max     int
	message error
}

var _ CheckerThrower = (*lengthChecker)(nil)

// Check checks a string's length.
func (c *lengthChecker) Check(in interface{}) error {
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
func (c *lengthChecker) Throws(msg string) CheckerThrower {
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
func Length(min int, max int) CheckerThrower {
	if min < 0 || max < 0 {
		panic("min and max cannot be negative")
	}

	check := &lengthChecker{
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

// typeChecker checks that an input value can be converted to a specified type.
type typeChecker struct {
	t       reflect.Type
	strict  bool
	message error
}

var _ CheckerThrower = (*typeChecker)(nil)

// Check checks that the input value can be converted to the specified type.
func (c *typeChecker) Check(in interface{}) error {
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
func (c *typeChecker) Throws(msg string) CheckerThrower {
	c.message = errors.New(msg)
	return c
}

// IsType returns a new Checker that checks whether the input is a valid type. This uses weak type
// checking which allows the input type to be different so long as it's convertible.
func IsType(t reflect.Type) CheckerThrower {
	return &typeChecker{
		t:       t,
		strict:  false,
		message: fmt.Errorf(i18n.get(i18nCheckTypeNotConvertible), t.Kind),
	}
}

// IsTypeStrict returns a new Checker that checks whether the input is a valid type. This uses strict
// type checking which only allows input types that are the same as the provided type.
func IsTypeStrict(t reflect.Type) CheckerThrower {
	return &typeChecker{
		t:       t,
		strict:  true,
		message: fmt.Errorf(i18n.get(i18nCheckTypeStrictBadType), t.Kind),
	}
}

package bleach

import (
	"fmt"
	"reflect"
	"strconv"
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

// ToString returns a new Mutator which converts the input to a string.
//
// 		bool  -> "true" | "false"
//		nil   -> "null"
func ToString() Mutator {
	fn := func(in interface{}) interface{} {
		if in == nil {
			return "null"
		}

		return fmt.Sprintf("%v", in)
	}

	return MutatorFunc(fn)
}

// ToInt returns a new Mutator which converts the input to an int64. Invalid types (such as nil)
// are converted to zero. If the input is a string, ToInt will attempt to parse it as a number.
//
//		bool	-> 1 | 0
//		nil		-> 0
//		3.14	-> 3
//		"3.14"	-> 3
//		-3.99	-> -3
func ToInt() Mutator {
	fn := func(in interface{}) interface{} {
		if in == nil {
			return int64(0)
		}

		if val, ok := in.(int64); ok {
			return val
		}

		switch reflect.TypeOf(in).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return reflect.ValueOf(in).Int()

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int64(reflect.ValueOf(in).Uint())

		case reflect.Float32, reflect.Float64:
			return int64(in.(float64))

		case reflect.String:
			str := in.(string)

			if val, err := strconv.ParseInt(str, 10, 64); err == nil {
				return val
			} else if val, err := strconv.ParseFloat(str, 10); err == nil {
				return int64(val)
			}

		case reflect.Bool:
			val := in.(bool)

			if val {
				return int64(1)
			}

			return int64(0)
		}

		return int64(0)
	}

	return MutatorFunc(fn)
}

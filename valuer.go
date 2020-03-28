package conf

import (
	"fmt"
	"reflect"
	"strconv"
)

// ParseValue parses str in to value.
func ParseValue(kind reflect.Kind, str string) (v reflect.Value, err error) {
	valuer, ok := valuers[kind]
	if !ok {
		err = ErrUnsupportedType
		return
	}
	v, err = valuer(str)
	return
}

type valuer func(string) (reflect.Value, error)

var valuers = map[reflect.Kind]valuer{
	reflect.String:  strValuer,
	reflect.Int:     intValuerGenerator(func(i int64) interface{} { return int(i) }),         // nolint: lll
	reflect.Int8:    intValuerGenerator(func(i int64) interface{} { return int8(i) }),        // nolint: lll
	reflect.Int16:   intValuerGenerator(func(i int64) interface{} { return int16(i) }),       // nolint: lll
	reflect.Int32:   intValuerGenerator(func(i int64) interface{} { return int32(i) }),       // nolint: lll
	reflect.Int64:   intValuerGenerator(func(i int64) interface{} { return i }),              // nolint: lll
	reflect.Uint:    uintValuerGenerator(func(i uint64) interface{} { return uint(i) }),      // nolint: lll
	reflect.Uint8:   uintValuerGenerator(func(i uint64) interface{} { return uint8(i) }),     // nolint: lll
	reflect.Uint16:  uintValuerGenerator(func(i uint64) interface{} { return uint16(i) }),    // nolint: lll
	reflect.Uint32:  uintValuerGenerator(func(i uint64) interface{} { return uint32(i) }),    // nolint: lll
	reflect.Uint64:  uintValuerGenerator(func(i uint64) interface{} { return i }),            // nolint: lll
	reflect.Float32: floatValuerGenerator(func(f float64) interface{} { return float32(f) }), // nolint: lll
	reflect.Float64: floatValuerGenerator(func(f float64) interface{} { return f }),          // nolint: lll
	reflect.Bool:    boolValuer,
	// nolint: godox
	// TODO: support more kinds.
	// reflect.Complex64
	// reflect.Complex128
	// reflect.Array
	// reflect.Chan
	// reflect.Slice
}

func strValuer(str string) (reflect.Value, error) {
	return reflect.ValueOf(str), nil
}

func intValuerGenerator(fn func(i int64) interface{}) valuer {
	return func(str string) (reflect.Value, error) {
		i, err := strconv.ParseInt(str, 0, 0)
		if err != nil {
			err = fmt.Errorf(
				"failed to parse int from \"%s\". %v", str, ErrInvalidValue)
			return reflect.Value{}, err
		}
		return reflect.ValueOf(fn(i)), nil
	}
}

func uintValuerGenerator(fn func(i uint64) interface{}) valuer {
	return func(str string) (reflect.Value, error) {
		ui, err := strconv.ParseUint(str, 0, 0)
		if err != nil {
			err = fmt.Errorf(
				"failed to parse uint from \"%s\". %v", str, ErrInvalidValue)
			return reflect.Value{}, err
		}
		return reflect.ValueOf(fn(ui)), nil
	}
}

func floatValuerGenerator(fn func(i float64) interface{}) valuer {
	return func(str string) (reflect.Value, error) {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			err = fmt.Errorf(
				"failed to parse float from \"%s\". %v", str, ErrInvalidValue)
			return reflect.Value{}, err
		}
		return reflect.ValueOf(fn(f)), nil
	}
}

func boolValuer(str string) (reflect.Value, error) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		err = fmt.Errorf(
			"failed to parse bool for from \"%s\". %w", str, ErrInvalidValue)
		return reflect.Value{}, err
	}
	return reflect.ValueOf(b), nil
}

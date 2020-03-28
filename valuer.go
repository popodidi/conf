package conf

import (
	"fmt"
	"reflect"
	"strconv"
)

// ParseValue parses str in to value.
func ParseValue(kind reflect.Kind, str string) (v reflect.Value, err error) {
	// parse value
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
	reflect.Int:     intValuer,
	reflect.Float64: float64Valuer,
	reflect.Bool:    boolValuer,
	// nolint: godox
	// TODO: support more kinds.
	// reflect.Int8
	// reflect.Int16
	// reflect.Int32
	// reflect.Int64
	// reflect.Uint
	// reflect.Uint8
	// reflect.Uint16
	// reflect.Uint32
	// reflect.Uint64
	// reflect.Uintptr
	// reflect.Float32
	// reflect.Complex64
	// reflect.Complex128
	// reflect.Array
	// reflect.Chan
	// reflect.Func
	// reflect.Interface
	// reflect.Map
	// reflect.Ptr
	// reflect.Slice
	// reflect.Struct
}

func strValuer(str string) (reflect.Value, error) {
	return reflect.ValueOf(str), nil
}

func intValuer(str string) (reflect.Value, error) {
	i, err := strconv.ParseInt(str, 0, 0)
	if err != nil {
		err = fmt.Errorf(
			"failed to parse int from \"%s\". %v", str, ErrInvalidValue)
		return reflect.Value{}, err
	}
	return reflect.ValueOf(int(i)), nil
}

func float64Valuer(str string) (reflect.Value, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		err = fmt.Errorf(
			"failed to parse float from \"%s\". %v", str, ErrInvalidValue)
		return reflect.Value{}, err
	}
	return reflect.ValueOf(f), nil
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

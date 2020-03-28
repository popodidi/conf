package conf

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func recordInMap(m map[string]struct{}) func(field reflect.Value,
	typeField reflect.StructField, key string) error {
	return func(
		field reflect.Value, typeField reflect.StructField, key string) error {
		m[key] = struct{}{}
		return nil
	}
}

func TestIter(t *testing.T) {
	var c testCfg
	// Test err
	err := iterFields(
		reflect.ValueOf(c),
		"prefix",
		nil,
	)
	require.True(t, errors.Is(err, ErrConfigNotPtr), err)

	// Test success
	m := make(map[string]struct{})
	iterFunc := recordInMap(m)
	require.NoError(t, iterFields(
		reflect.ValueOf(&c),
		"prefix",
		iterFunc,
	))
	exp := map[string]struct{}{
		"PREFIX_HEY_YOYO": struct{}{},
		"PREFIX_HEY_YO":   struct{}{},
		"PREFIX_HI":       struct{}{},
		"PREFIX_HIEMPTY":  struct{}{},
		"PREFIX_QQ":       struct{}{},
	}
	require.Equal(t, exp, m)
}

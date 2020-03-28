package conf

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	var c TestCfg
	// Test err
	err := iterFields(reflect.ValueOf(c), nil, nil)
	require.True(t, errors.Is(err, ErrConfigNotPtr), err)

	// Test success
	m := make(Map)
	require.NoError(t, iterFields(reflect.ValueOf(&c), nil, mapRecorder(m)))
	fmt.Println(m)

	expected := [][]string{
		[]string{"bool", "Hi"},
		[]string{"bool", "HiEmpty"},
		[]string{"string", "QQ"},
		[]string{"int", "YO", "Hey"},
		[]string{"int", "YOYO", "Hey"},
	}
	for _, path := range expected {
		require.Equal(t, path[0], m.In(path[2:]...).Get(path[1]), path)
	}
}

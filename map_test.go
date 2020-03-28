package conf

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	m := make(Map)

	// In
	require.Nil(t, m.In("a", "b", "c"))
	m["a"] = make(Map)
	require.Equal(t, m["a"], m.MustIn("a"))
	require.Nil(t, m.ValidateStrMap())

	// MustIn
	require.Equal(t, m["a"].(Map)["b"].(Map)["c"], m.MustIn("a", "b", "c"))
	require.Equal(t, m["a"].(Map)["b"], m.MustIn("a", "b"))
	require.Equal(t, m["a"], m.MustIn("a"))
	require.Nil(t, m.ValidateStrMap())

	// Set
	require.Equal(t, ErrNilMap, m.In("d").Set("x", "x"))
	require.NoError(t, m.Set("x", "x"))
	require.Equal(t, "x", m["x"])
	require.Nil(t, m.ValidateStrMap())

	// Get
	require.NoError(t, m.MustIn("d").Set("y", "y"))
	require.NoError(t, m.MustIn("d", "x").Set("y", "y"))
	require.Nil(t, m.Get("e"))                // nil for nothing
	require.Nil(t, m.Get("d"))                // nil for a Map
	require.Nil(t, m.In("d").Get("x"))        // nil for a Map
	require.Equal(t, "y", m.In("d").Get("y")) // nil for a Map
	require.Nil(t, m.ValidateStrMap())

	// ValidateStrMap
	badMap := Map(map[string]interface{}{"hello": 1})
	require.True(t, errors.Is(badMap.ValidateStrMap(), ErrValueType))

	// Iter
	m = make(Map)
	require.NoError(t, m.Set("a", "b"))
	require.NoError(t, m.MustIn("c").Set("d", "e"))
	require.NoError(t, m.MustIn("f", "g").Set("h", "i"))
	require.NoError(t, m.MustIn("j", "k").Set("l", "m"))
	clone := make(Map)
	m.Iter(func(key string, val interface{}, path ...string) bool {
		require.NoError(t, clone.MustIn(path...).Set(key, val))
		return true
	})
	require.Equal(t, m, clone)
}

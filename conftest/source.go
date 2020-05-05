package conftest

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/popodidi/conf"
)

// MockConf is the sample map for test.
var MockConf = make(conf.Map)

func init() { // nolint: gochecknoinits
	must(MockConf.Set("a", "b"))
	must(MockConf.MustIn("c").Set("d", "e"))
	must(MockConf.MustIn("f", "g").Set("h", "i"))
	must(MockConf.MustIn("j", "k").Set("str", "m"))
	must(MockConf.MustIn("j", "k").Set("int", 1))
	must(MockConf.MustIn("j", "k").Set("float", 1.234))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// ReaderTest tests the reader.
func ReaderTest(t *testing.T, reader conf.Reader) {
	MockConf.Iter(func(key string, val interface{}, path ...string) bool {
		value, ok := reader.Read(key, path...)
		require.True(t, ok, "%v: %s not found", path, key)
		kind := reflect.ValueOf(val).Kind()
		parsed, err := conf.ParseValue(kind, value)
		require.NoError(t, err)
		require.Equal(t, val, parsed.Interface(), "%v!=%v", val, parsed.Interface())
		return true
	})
	{
		value, ok := reader.Read("val", "doesn't", "not", "exist")
		require.Empty(t, value)
		require.False(t, ok)
	}
}

// ExporterTest tests the exporter.
func ExporterTest(t *testing.T, exp string, exporter conf.Exporter) {
	var b strings.Builder
	require.NoError(t, exporter.Export(MockConf, &b))
	require.Equal(t, exp, b.String())
}

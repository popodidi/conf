package flag

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/popodidi/conf"
	"github.com/popodidi/conf/conftest"
)

func TestReader(t *testing.T) {
	reader := New()
	var args []string
	conftest.MockConf.Iter(
		func(key string, val interface{}, path ...string) (next bool) {
			name := reader.flagName(key, path...)
			args = append(args, fmt.Sprintf("--%s=%v", name, val))
			tag := conf.FieldTag{
				Default: nil,
				Usage:   "...",
			}
			require.NoError(t, reader.Configure(
				reflect.ValueOf(val).Type(), tag, key, path...))
			return true
		})
	require.NoError(t, reader.Parse(args))
	conftest.ReaderTest(t, reader)
}

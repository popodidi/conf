package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/popodidi/conf/conftest"
)

func TestVarName(t *testing.T) {
	cases := map[string][]string{
		"B_C_A":  []string{"a", "b", "c"},
		"B_C_AA": []string{"aa", "b", "c"},
		"AA":     []string{"aa"},
	}
	e := &env{}
	for exp, c := range cases {
		require.Equal(t, exp, e.varName(c[0], c[1:]...))
	}
}

func TestEnv(t *testing.T) {
	prepareEnv(t)
	defer clearEnv(t)
	conftest.ReaderTest(t, New())
}

func prepareEnv(t *testing.T) {
	e := &env{}
	conftest.MockConf.Iter(
		func(key string, val interface{}, path ...string) bool {
			require.NoError(t,
				os.Setenv(e.varName(key, path...), fmt.Sprintf("%v", val)))
			return true
		},
	)
}
func clearEnv(t *testing.T) {
	e := &env{}
	conftest.MockConf.Iter(
		func(key string, val interface{}, path ...string) bool {
			require.NoError(t, os.Unsetenv(e.varName(key, path...)))
			return true
		},
	)
}

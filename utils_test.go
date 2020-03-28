package conf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToOldKey(t *testing.T) {
	cases := map[string][]string{
		"B_C_A":  []string{"a", "b", "c"},
		"B_C_AA": []string{"aa", "b", "c"},
		"AA":     []string{"aa"},
	}
	for exp, c := range cases {
		require.Equal(t, exp, toOldKey(c[0], c[1:]...))
	}
}

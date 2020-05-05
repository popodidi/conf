package conf

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTag(t *testing.T) {
	yoyo := "yoyo"
	cases := map[string]FieldTag{
		"":                              FieldTag{},
		"default:" + yoyo:               FieldTag{Default: &yoyo},
		"usage:wo":                      FieldTag{Usage: "wo"},
		"default:" + yoyo + ",usage:wo": FieldTag{Default: &yoyo, Usage: "wo"},
	}
	for str, exp := range cases {
		act, err := parseTag(str)
		require.NoError(t, err)
		require.Equal(t, exp, act)
	}
}

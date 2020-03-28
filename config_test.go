package conf

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfLoad(t *testing.T) {
	var c TestCfg

	cfg := New(&c)
	assert.True(t,
		errors.Is(cfg.Load([]Reader{InvalidSrc}), ErrInvalidValue))
	assert.True(t,
		errors.Is(cfg.Load([]Reader{MissingSrc}), ErrConfigNotFound))
	assert.NoError(t, cfg.Load([]Reader{MockSrc}))
	assert.Equal(t, true, c.Hi)
	assert.Equal(t, true, c.HiEmpty)
	assert.Equal(t, "str", c.QQ)
	assert.Equal(t, 87, c.Hey.YO)
	assert.Equal(t, 1, c.Hey.YOYO)

	var invalidValueC InvalidValueCfg
	ivCfg := New(&invalidValueC)
	assert.True(t,
		errors.Is(ivCfg.Load([]Reader{MockSrc}), ErrUnsupportedType))

	var invalidTagC InvalidTagCfg
	itCfg := New(&invalidTagC)
	assert.True(t, errors.Is(itCfg.Load([]Reader{MockSrc}), ErrInvalidTag))
}

func TestConfMap(t *testing.T) {
	var c TestCfg
	cfg := New(&c)
	_, err := cfg.Map()
	assert.Error(t, err, ErrConfigNotLoaded)
	assert.NoError(t, cfg.Load([]Reader{MockSrc}))
	m, err := cfg.Map()
	assert.NoError(t, err)
	assert.Equal(t, true, m["HI"])
	assert.Equal(t, true, m["HIEMPTY"])
	assert.Equal(t, "str", m["QQ"])
	assert.Equal(t, 87, m["HEY_YO"])
	assert.Equal(t, 1, m["HEY_YOYO"])
}

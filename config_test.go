package conf

import (
	"errors"
	"strings"
	"testing"

	"github.com/popodidi/conf/source"
	"github.com/popodidi/conf/source/mock"
	"github.com/popodidi/conf/source/yaml"
	"github.com/stretchr/testify/assert"
)

type testSubCfg struct {
	YO   int
	YOYO int `conf:"default:1"`
}

type testCfg struct {
	Hi      bool
	HiEmpty bool `conf:"default:true"`
	QQ      string
	Hey     testSubCfg
}

type invalidValueCfg struct {
	Hi func()
}

type invalidTagCfg struct {
	Hi bool `conf:"hello"`
}

var mockSrc = mock.New(
	map[string]string{
		"HI":     "true",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

var missingSrc = mock.New(
	map[string]string{
		// "HI":     "true",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

var invalidSrc = mock.New(
	map[string]string{
		"HI":     "hello",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

func TestLoad(t *testing.T) {
	var c testCfg

	cfg := New(&c)
	assert.True(t,
		errors.Is(cfg.Load([]source.Reader{invalidSrc}), ErrInvalidValue))
	assert.True(t,
		errors.Is(cfg.Load([]source.Reader{missingSrc}), ErrConfigNotFound))
	assert.NoError(t, cfg.Load([]source.Reader{mockSrc}))
	assert.Equal(t, true, c.Hi)
	assert.Equal(t, true, c.HiEmpty)
	assert.Equal(t, "str", c.QQ)
	assert.Equal(t, 87, c.Hey.YO)
	assert.Equal(t, 1, c.Hey.YOYO)

	var invalidValueC invalidValueCfg
	ivCfg := New(&invalidValueC)
	assert.True(t,
		errors.Is(ivCfg.Load([]source.Reader{mockSrc}), ErrUnsupportedType))

	var invalidTagC invalidTagCfg
	itCfg := New(&invalidTagC)
	assert.True(t, errors.Is(itCfg.Load([]source.Reader{mockSrc}), ErrInvalidTag))
}

func TestMap(t *testing.T) {
	var c testCfg
	cfg := New(&c)
	_, err := cfg.Map()
	assert.Error(t, err, ErrConfigNotLoaded)
	assert.NoError(t, cfg.Load([]source.Reader{mockSrc}))
	m, err := cfg.Map()
	assert.NoError(t, err)
	assert.Equal(t, true, m["HI"])
	assert.Equal(t, true, m["HIEMPTY"])
	assert.Equal(t, "str", m["QQ"])
	assert.Equal(t, 87, m["HEY_YO"])
	assert.Equal(t, 1, m["HEY_YOYO"])
}

func TestExport(t *testing.T) {
	var c testCfg
	var b strings.Builder
	cfg := New(&c)
	assert.NoError(t, cfg.Load([]source.Reader{mockSrc}))
	assert.NoError(t, cfg.Export(yaml.NewExporter(), &b))
	expected := "HEY_YO: 87\nHEY_YOYO: 1\nHI: true\nHIEMPTY: true\nQQ: str\n"
	assert.Equal(t, expected, b.String())
}

func TestTemplate(t *testing.T) {
	var c testCfg
	cfg := New(&c)
	assert.NoError(t, cfg.Load([]source.Reader{mockSrc}))
	tmpl, err := cfg.Template(yaml.NewExporter())
	assert.NoError(t, err)
	expected := "HEY_YO: int\n" +
		"HEY_YOYO: int\nHI: bool\nHIEMPTY: bool\nQQ: string\n"
	assert.Equal(t, expected, tmpl)
}

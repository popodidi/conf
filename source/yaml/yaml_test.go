package yaml

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/popodidi/conf"
)

func TestExport(t *testing.T) {
	var c conf.TestCfg
	var b strings.Builder
	cfg := conf.New(&c)
	assert.NoError(t, cfg.Load([]conf.Reader{conf.MockSrc}))
	assert.NoError(t, cfg.Export(NewExporter(), &b))
	expected := "HEY_YO: 87\nHEY_YOYO: 1\nHI: true\nHIEMPTY: true\nQQ: str\n"
	assert.Equal(t, expected, b.String())
}

func TestTemplate(t *testing.T) {
	var c conf.TestCfg
	cfg := conf.New(&c)
	assert.NoError(t, cfg.Load([]conf.Reader{conf.MockSrc}))
	tmpl, err := cfg.Template(NewExporter())
	assert.NoError(t, err)
	expected := "HEY_YO: int\n" +
		"HEY_YOYO: int\nHI: bool\nHIEMPTY: bool\nQQ: string\n"
	assert.Equal(t, expected, tmpl)
}

package yaml

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/popodidi/conf/conftest"
)

func TestReader(t *testing.T) {
	ymlPath, clean := prepareYml(t)
	defer clean()

	conftest.ReaderTest(t, New(ymlPath))
}

func TestExport(t *testing.T) {
	expected := "a: b\nc:\n  d: e\nf:\n  g:\n    h: i\nj:\n  k:\n    float: 1.234\n    int: 1\n    str: m\n" // nolint: lll
	conftest.ExporterTest(t, expected, &exporter{})
}

func prepareYml(t *testing.T) (ymlPath string, clean func()) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	ymlPath = filepath.Join(dir,
		fmt.Sprintf("test_conf_%d.yml", time.Now().UnixNano()))
	f, err := os.Create(ymlPath)
	require.NoError(t, err)
	enc := yaml.NewEncoder(f)
	require.NoError(t, enc.Encode(conftest.MockConf))
	clean = func() {
		require.NoError(t, os.Remove(ymlPath))
	}
	return
}

package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/popodidi/conf/conftest"
)

func TestReader(t *testing.T) {
	ymlPath, clean := prepareYml(t)
	defer clean()

	conftest.ReaderTest(t, New(ymlPath))
}

func TestExport(t *testing.T) {
	expected := "{\"a\":\"b\",\"c\":{\"d\":\"e\"},\"f\":{\"g\":{\"h\":\"i\"}},\"j\":{\"k\":{\"float\":1.234,\"int\":1,\"str\":\"m\"}}}" // nolint: lll
	conftest.ExporterTest(t, expected, &exporter{})
}

func prepareYml(t *testing.T) (ymlPath string, clean func()) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	ymlPath = filepath.Join(dir,
		fmt.Sprintf("test_conf_%d.json", time.Now().UnixNano()))
	f, err := os.Create(ymlPath)
	require.NoError(t, err)
	enc := json.NewEncoder(f)
	require.NoError(t, enc.Encode(conftest.MockConf))
	clean = func() {
		require.NoError(t, os.Remove(ymlPath))
	}
	return
}

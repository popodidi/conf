package memory

import (
	"testing"

	"github.com/popodidi/conf/conftest"
)

func TestReader(t *testing.T) {
	conftest.ReaderTest(t, New(conftest.MockConf))
}

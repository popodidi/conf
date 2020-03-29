package conf

import (
	"fmt"
	"io"
	"strings"
)

// nolint: golint
type TestSubCfg struct {
	YO   int
	YOYO int `conf:"default:1"`
}

// nolint: golint
type TestCfg struct {
	Hi      bool
	HiEmpty bool `conf:"default:true"`
	QQ      string
	Hey     TestSubCfg
}

// nolint: golint
type InvalidValueCfg struct {
	Hi func()
}

// nolint: golint
type InvalidTagCfg struct {
	Hi bool `conf:"hello"`
}

// nolint: golint
var MockSrc = NewMock(
	map[string]interface{}{
		"Hi": "true",
		"QQ": "str",
		"Hey": Map(map[string]interface{}{
			"YO": "87",
		}),
	},
)

// nolint: golint
var MissingSrc = NewMock(
	map[string]interface{}{
		// "HI":     "true",
		"QQ": "str",
		"Hey": Map(map[string]interface{}{
			"YO": "87",
		}),
	},
)

// nolint: golint
var InvalidSrc = NewMock(
	map[string]interface{}{
		"Hi": "hello",
		"QQ": "str",
		"Hey": Map(map[string]interface{}{
			"YO": "87",
		}),
	},
)

// New returns a mock source.
func NewMock(m Map) Source {
	return mock(m)
}

type mock Map

func (s mock) Read(key string, path ...string) (value string, exists bool) {
	val := Map(s).In(path...).Get(key)
	if val == nil {
		return
	}
	value = fmt.Sprintf("%v", val)
	exists = true
	return
}

func (s mock) Export(m Map, writer io.Writer) error {
	var err error
	m.Iter(func(key string, val interface{}, path ...string) (next bool) {
		if len(path) == 0 {
			_, err = writer.Write([]byte(fmt.Sprintf("%s: %v", key, val)))
		} else {
			_, err = writer.Write([]byte(
				fmt.Sprintf("%s.%s: %v", strings.Join(path, "."), key, val)))
		}
		return err == nil
	})
	return err
}

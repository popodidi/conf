package conf

import (
	"fmt"
	"io"
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
	map[string]string{
		"HI":     "true",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

// nolint: golint
var MissingSrc = NewMock(
	map[string]string{
		// "HI":     "true",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

// nolint: golint
var InvalidSrc = NewMock(
	map[string]string{
		"HI":     "hello",
		"QQ":     "str",
		"HEY_YO": "87",
	},
)

// NewMock returns a mock source with m.
func NewMock(m map[string]string) Source {
	return &mock{m}
}

type mock struct {
	m map[string]string
}

func (s *mock) Read(key string, path ...string) (value string, exists bool) {
	value, exists = s.m[key]
	return
}

func (s *mock) Export(m Map, writer io.Writer) error {
	_, err := writer.Write([]byte(fmt.Sprintf("%v", m)))
	return err
}

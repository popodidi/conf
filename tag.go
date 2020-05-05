package conf

import (
	"strings"
)

const (
	tagDefaultPrefix = "default:"
	tagUsagePrefix   = "usage:"
)

// FieldTag is the conf tag of a struct field
type FieldTag struct {
	Default *string
	Usage   string
}

func parseTag(tagStr string) (ft FieldTag, err error) {
	if len(tagStr) == 0 {
		return
	}
	tags := strings.Split(tagStr, ",")
	for _, t := range tags {
		switch {
		case strings.HasPrefix(t, tagDefaultPrefix):
			defaultValue := t[8:]
			ft.Default = &defaultValue
		case strings.HasPrefix(t, tagUsagePrefix):
			ft.Usage = t[6:]
		default:
			err = ErrInvalidTag
			return
		}
	}
	return
}

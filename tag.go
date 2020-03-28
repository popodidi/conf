package conf

import (
	"strings"
)

const (
	tagDefaultPrefix = "default:"
)

type fieldTag struct {
	hasDefault   bool
	defaultValue string
}

func parseTag(tagStr string) (ft fieldTag, err error) {
	if len(tagStr) == 0 {
		return
	}
	tags := strings.Split(tagStr, ",")
	for _, t := range tags {
		// XXX: This should be a switch-case. However, since we only support one
		// `default` tag at the moment, an if-else will do the work.
		if strings.HasPrefix(t, tagDefaultPrefix) {
			ft.hasDefault = true
			ft.defaultValue = t[8:]
			continue
		}
		err = ErrInvalidTag
		return
	}
	return
}

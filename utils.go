package conf

import "strings"

const sep = "_"

func toOldKey(key string, path ...string) string {
	key = strings.ToUpper(key)
	if len(path) == 0 {
		return key
	}
	name := strings.Join(path, sep)
	name = strings.ToUpper(name)
	return name + sep + key
}

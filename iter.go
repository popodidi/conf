package conf

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	sep = "_"
)

func iterFields(ptrValue reflect.Value, prefix string,
	fn func(field reflect.Value, typeField reflect.StructField, key string) error,
) (err error) {
	ptrType := ptrValue.Type()
	if ptrType == nil {
		return ErrNilConfig
	}

	configKind := ptrValue.Kind()
	if configKind != reflect.Ptr {
		return ErrConfigNotPtr
	}

	configValue := ptrValue.Elem()
	configType := configValue.Type()

	if configValue.Kind() != reflect.Struct {
		return ErrConfigNotStruct
	}

	for i := 0; i < configType.NumField(); i++ {
		v := configValue.Field(i)
		v.CanAddr()
		t := configType.Field(i)
		key := t.Name
		if prefix != "" {
			key = fmt.Sprintf("%s%s%s", prefix, sep, key)
		}
		key = strings.ToUpper(key)

		// recursive call for struct fields
		if v.Kind() == reflect.Struct {
			if !v.CanAddr() {
				return ErrCantAddr
			}
			err = iterFields(v.Addr(), key, fn)
			if err != nil {
				return err
			}
			continue
		}
		err = fn(v, t, key)
		if err != nil {
			return err
		}
	}
	return nil
}

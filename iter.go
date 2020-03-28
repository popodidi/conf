package conf

import (
	"reflect"
)

type iterator func(
	field reflect.Value,
	typeField reflect.StructField,
	key string,
	path ...string,
) error

func iterFields(ptrValue reflect.Value, prepath []string,
	fn iterator) (err error) {
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
		path := append(prepath[:0:0], prepath...)
		key := t.Name

		// recursive call for struct fields
		if v.Kind() == reflect.Struct {
			if !v.CanAddr() {
				return ErrCantAddr
			}
			path = append(path, key)
			err = iterFields(v.Addr(), path, fn)
			if err != nil {
				return err
			}
			continue
		}

		// call the iterator fn
		err = fn(v, t, key, path...)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapRecorder(m Map) iterator {
	return func(
		field reflect.Value, typeField reflect.StructField,
		key string, path ...string) (
		err error,
	) {
		return m.MustIn(path...).Set(key, typeField.Type.Name())
	}
}

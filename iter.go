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

	ptrValueKind := ptrValue.Kind()
	if ptrValueKind != reflect.Ptr {
		return ErrConfigNotPtr
	}

	ptrValueElem := ptrValue.Elem()
	ptrValueElemType := ptrValueElem.Type()

	if ptrValueElem.Kind() != reflect.Struct {
		return ErrConfigNotStruct
	}

	for i := 0; i < ptrValueElemType.NumField(); i++ {
		v := ptrValueElem.Field(i)
		t := ptrValueElemType.Field(i)
		path := append(prepath[:0:0], prepath...)
		key := t.Name

		// recursive call for struct fields
		if v.Kind() == reflect.Struct && !isScanner(v.Type()) {
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

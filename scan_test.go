package conf

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type T time.Time

var now = time.Now()

func (t *T) Scan(str string) error {
	*t = T(now)
	return nil
}

type Str struct {
	str string
}

func (s *Str) Scan(str string) error {
	s.str = str
	return nil
}

func TestScan(t *testing.T) {
	{
		var a struct {
			Str string
		}
		require.NoError(t,
			Scan(reflect.ValueOf(&a).Elem().FieldByName("Str"), "hello"))
		require.Equal(t, "hello", a.Str)
		require.NoError(t,
			Scan(reflect.ValueOf(&a).Elem().FieldByName("Str"), ""))
		require.Equal(t, "", a.Str)
	}
	{
		var a struct {
			T *T
		}
		require.NoError(t,
			Scan(reflect.ValueOf(&a).Elem().FieldByName("T"), "hello"))
		require.Equal(t, now, time.Time(*a.T))
	}
	{
		var a struct {
			Str *Str
		}
		require.NoError(t,
			Scan(reflect.ValueOf(&a).Elem().FieldByName("Str"), "hello"))
		require.Equal(t, "hello", a.Str.str)
	}
}

func TestValuers(t *testing.T) {
	cases := [][]interface{}{ // [kind, str, expected_value]
		[]interface{}{reflect.String, "string", "string"},
		[]interface{}{reflect.Int, "1", int(1)},
		[]interface{}{reflect.Int8, "1", int8(1)},
		[]interface{}{reflect.Int16, "1", int16(1)},
		[]interface{}{reflect.Int32, "1", int32(1)},
		[]interface{}{reflect.Int64, "1", int64(1)},
		[]interface{}{reflect.Uint, "1", uint(1)},
		[]interface{}{reflect.Uint8, "1", uint8(1)},
		[]interface{}{reflect.Uint16, "1", uint16(1)},
		[]interface{}{reflect.Uint32, "1", uint32(1)},
		[]interface{}{reflect.Uint64, "1", uint64(1)},
		[]interface{}{reflect.Float32, "1.234", float32(1.234)},
		[]interface{}{reflect.Float64, "1.234", float64(1.234)},
		[]interface{}{reflect.Bool, "true", true},
	}

	for _, c := range cases {
		val, err := ParseValue(c[0].(reflect.Kind), c[1].(string))
		require.NoError(t, err)
		require.Equal(t, c[2], val.Interface())
	}
}

package conf

import (
	"fmt"
)

// Map defines the alias type for nested map. The map reflects the values read
// from sources. A value in the map is either a string or another Map.
type Map map[string]interface{}

// ValidateStrMap checks if the map is a map which contains only string or sub
// Map.
func (m Map) ValidateStrMap() error {
	for k, v := range m {
		if _, ok := v.(string); ok {
			continue
		} else if m, ok := v.(Map); ok {
			err := m.ValidateStrMap()
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid value type of %s. %w", k, ErrValueType)
		}
	}
	return nil
}

// In returns the sub map of m with path. `nil` will be returned if no map is
// found.
func (m Map) In(path ...string) Map {
	if len(path) == 0 {
		return m
	}
	sub, ok := m[path[0]]
	if !ok {
		return nil
	}
	subM, ok := sub.(Map)
	if !ok {
		return nil
	}
	return subM.In(path[1:]...)
}

// MustIn returns the sub map of m with path and creates maps all along the path
// if needed. The key collisions, i.e the existing values, along the path will
// be overridden with new maps.
func (m Map) MustIn(path ...string) Map {
	if len(path) == 0 {
		return m
	}
	var (
		sub  interface{}
		subM Map
		ok   bool
	)
	if sub, ok = m[path[0]]; ok {
		subM, ok = sub.(Map)
	}
	if !ok {
		subM = make(Map)
		m[path[0]] = subM
	}
	return subM.MustIn(path[1:]...)
}

// Set sets the val in the map with path and returns error if the path does not
// exist.
func (m Map) Set(key string, val interface{}) error {
	if m == nil {
		return ErrNilMap
	}
	m[key] = val
	return nil
}

// Get returns the value in path with key and returns nil if nothing found or
// the val is a sub Map.
func (m Map) Get(key string) interface{} {
	if m == nil {
		return nil
	}
	val, ok := m[key]
	if !ok {
		return nil
	}
	if _, isMap := val.(Map); isMap {
		return nil
	}
	return val
}

// Iter iterates all the values in map.
func (m Map) Iter(
	f func(key string, val interface{}, path ...string) (next bool)) {
	m.iter([]string{}, f)
}

func (m Map) iter(prepath []string,
	f func(key string, val interface{}, path ...string) (next bool)) (
	next bool) {
	for k, v := range m {
		if sub, ok := v.(Map); ok {
			subPrepath := append(prepath[:0:0], prepath...)
			subPrepath = append(subPrepath, k)
			next = sub.iter(subPrepath, f)
			if !next {
				return
			}
		} else {
			f(k, v, prepath...)
		}
	}
	return true
}

// Clone returns a shallow copy of the map.
func (m Map) Clone() (clone Map, err error) {
	clone = make(Map)
	m.Iter(func(key string, val interface{}, path ...string) bool {
		err = clone.MustIn(path...).Set(key, val)
		return err == nil
	})
	if err != nil {
		clone = nil
	}
	return
}

// FlattenedClone returns a flattened shallow copy of the map.
func (m Map) FlattenedClone(keyFn func(key string, path ...string) string) (
	clone Map, err error) {
	clone = make(Map)
	m.Iter(func(key string, val interface{}, path ...string) bool {
		err = clone.Set(keyFn(key, path...), val)
		return err == nil
	})
	if err != nil {
		clone = nil
	}
	return
}

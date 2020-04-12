package conf

import "errors"

// Exported Errors.
var (
	// Config errors
	ErrCantAddr        = errors.New("value can't addr")
	ErrCantInterface   = errors.New("value can't interface")
	ErrConfigNotFound  = errors.New("config not found")
	ErrConfigNotLoaded = errors.New("config not loaded")
	ErrConfigNotPtr    = errors.New("config should be a struct pointer")
	ErrConfigNotStruct = errors.New("config should be a struct pointer")
	ErrDuplicateKey    = errors.New("duplicate key")
	ErrInvalidTag      = errors.New("invalid tag")
	ErrInvalidValue    = errors.New("invalid config value")
	ErrNilConfig       = errors.New("nil config")
	ErrUnsupportedType = errors.New("unsupported type")

	// Map errors
	ErrNilMap    = errors.New("nil map")
	ErrValueType = errors.New("map value should be either string or Map")
)

package optional

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type optionalType interface {
	string | float64
}

type Of[T optionalType] struct {
	value T
	valid bool
}

func New[T optionalType](value T, valid bool) Of[T] {
	return Of[T]{
		value: value,
		valid: valid,
	}
}

func NewWithValue[T optionalType](value T) Of[T] {
	return Of[T]{
		value: value,
		valid: true,
	}
}

func (o Of[T]) ValueOrZero() T {
	if o.valid {
		return o.value
	}

	return defaultValue[T]()
}

func (o Of[T]) HasValue() bool {
	return o.valid
}

func (o *Of[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		o.value = defaultValue[T]()
		o.valid = false
		return nil
	}

	if err := json.Unmarshal(data, &o.value); err != nil {
		return err
	}

	o.valid = true

	return nil
}

func (o Of[T]) MarshalJSON() ([]byte, error) {
	if o.valid {
		return json.Marshal(o.value)
	}

	return json.Marshal(nil)
}

// Scan sql.Scannerを満たすためのメソッド
func (o Of[T]) Scan(value any) error {
	switch ((any)(o.value)).(type) {
	case string:
		nullString := sql.NullString{}
		err := nullString.Scan(value)
		if err != nil {
			return err
		}
		o.value = any(nullString.String).(T)
		o.valid = nullString.Valid
	case float64:
		nullFloat64 := sql.NullFloat64{}
		err := nullFloat64.Scan(value)
		if err != nil {
			return err
		}
		o.value = any(nullFloat64.Float64).(T)
		o.valid = nullFloat64.Valid
	default:
		return fmt.Errorf("unsupported type for Of[T].Scan: %T", o.value)
	}

	return nil
}

// Value driver.Valuerを満たすためのメソッド
func (o Of[T]) Value() (driver.Value, error) {
	if !o.valid {
		return nil, nil
	}

	return o.value, nil
}

func defaultValue[T optionalType]() T {
	v := new(T)
	return *v
}

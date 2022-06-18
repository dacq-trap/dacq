package optional

import (
	"bytes"
	"encoding/json"
)

type optionalType interface {
	string | float64
}

type Of[T optionalType] struct {
	Value T
	Valid bool
}

func New[T optionalType](value T, valid bool) Of[T] {
	return Of[T]{
		Value: value,
		Valid: valid,
	}
}

func NewWithValue[T optionalType](value T) Of[T] {
	return Of[T]{
		Value: value,
		Valid: true,
	}
}

func (o Of[T]) ValueOrZero() T {
	if o.Valid {
		return o.Value
	}

	return defaultValue[T]()
}

func (o Of[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		o.Value = defaultValue[T]()
		o.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &o.Value); err != nil {
		return err
	}

	o.Valid = true

	return nil
}

func (o Of[T]) MarshalJSON() ([]byte, error) {
	if o.Valid {
		return json.Marshal(o.Value)
	}

	return json.Marshal(nil)
}

func defaultValue[T optionalType]() T {
	v := new(T)
	return *v
}

package optional

import (
	"bytes"
	"encoding/json"
)

type optionalType interface {
	string | float64
}

type Optional[T optionalType] struct {
	Value T
	Valid bool
}

func New[T optionalType](value T, valid bool) Optional[T] {
	return Optional[T]{
		Value: value,
		Valid: valid,
	}
}

func NewWithValue[T optionalType](value T) Optional[T] {
	return Optional[T]{
		Value: value,
		Valid: true,
	}
}

func (o Optional[T]) ValueOrZero() T {
	if o.Valid {
		return o.Value
	}

	return defaultValue[T]()
}

func (o Optional[T]) UnmarshalJSON(data []byte) error {
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

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.Valid {
		return json.Marshal(o.Value)
	}

	return json.Marshal(nil)
}

func defaultValue[T optionalType]() T {
	v := new(T)
	return *v
}

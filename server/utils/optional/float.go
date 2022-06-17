package optional

import (
	"bytes"
	"database/sql"
	"strconv"
)

type Float struct {
	sql.NullFloat64
}

func FloatFrom(v float64) Float {
	return NewFloat(v, true)
}

func NewFloat(v float64, valid bool) Float {
	return Float{
		NullFloat64: sql.NullFloat64{
			Float64: v,
			Valid:   valid,
		},
	}
}

func (i Float) ValueOrZero() float64 {
	if i.Valid {
		return i.Float64
	}

	return 0
}

func (i *Float) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		i.Float64, i.Valid = 0, false
		return nil
	}

	if err := json.Unmarshal(data, &i.Float64); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

func (i Float) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Float64)
	}

	return json.Marshal(nil)
}

func (i *Float) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		i.Valid = false
		return nil
	}

	var err error
	i.Float64, err = strconv.ParseFloat(string(text), 64)
	i.Valid = err == nil

	return err
}

func (i Float) MarshalText() ([]byte, error) {
	if i.Valid {
		return []byte(strconv.FormatFloat(i.Float64, 'f', -1, 64)), nil
	}

	return []byte{}, nil
}

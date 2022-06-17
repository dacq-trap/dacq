package optional

import (
	"bytes"
	"database/sql"
)

type String struct {
	sql.NullString
}

func StringFrom(v string) String {
	return NewString(v, true)
}

func NewString(v string, valid bool) String {
	return String{
		NullString: sql.NullString{
			String: v,
			Valid:  valid,
		},
	}
}

func (s String) ValueOrZero() string {
	if s.Valid {
		return s.String
	}

	return ""
}

func (s *String) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		s.String, s.Valid = "", false
		return nil
	}

	if err := json.Unmarshal(data, &s.String); err != nil {
		return err
	}

	s.Valid = true
	return nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return json.Marshal(nil)
}

func (s *String) UnmarshalText(text []byte) error {
	s.String = string(text)
	s.Valid = s.String != ""

	return nil
}

func (s String) MarshalText() ([]byte, error) {
	if s.Valid {
		return []byte(s.String), nil
	}

	return []byte{}, nil
}

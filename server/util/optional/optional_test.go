package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dacq-trap/dacq/server/util/optional"
	"github.com/stretchr/testify/assert"
)

func TestOf_MarshalJSON(t *testing.T) {
	type S[T optional.OptionalType] struct {
		Foo optional.Of[T] `json:"foo"`
	}

	stringTests := []struct {
		name     string
		foo      optional.Of[string]
		expected string
	}{
		{
			name:     "ValidString",
			foo:      optional.NewWithValue("bar"),
			expected: `{"foo":"bar"}`,
		},
		{
			name:     "NullString",
			foo:      optional.New("", false),
			expected: `{"foo":null}`,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			s := S[string]{
				Foo: tt.foo,
			}
			got, err := json.Marshal(s)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expected, string(got))
			}
		})
	}

	floatTests := []struct {
		name     string
		foo      optional.Of[float64]
		expected string
	}{
		{
			name:     "ValidFloat",
			foo:      optional.NewWithValue(1.05),
			expected: `{"foo":1.05}`,
		},
		{
			name:     "NullFloat",
			foo:      optional.New(0.0, false),
			expected: `{"foo":null}`,
		},
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			s := S[float64]{
				Foo: tt.foo,
			}
			got, err := json.Marshal(s)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expected, string(got))
			}
		})
	}

}

func TestOf_UnmarshalJSON(t *testing.T) {
	type S[T optional.OptionalType] struct {
		Foo optional.Of[T] `json:"foo"`
	}

	stringTests := []struct {
		name        string
		rawJson     string
		expectedFoo optional.Of[string]
	}{
		{
			name:        "ValidString",
			rawJson:     `{"foo":"bar"}`,
			expectedFoo: optional.NewWithValue("bar"),
		},
		{
			name:        "NullString",
			rawJson:     `{"foo":null}`,
			expectedFoo: optional.New("", false),
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			var s S[string]
			err := json.Unmarshal([]byte(tt.rawJson), &s)
			expected := S[string]{
				Foo: tt.expectedFoo,
			}
			if assert.NoError(t, err) {
				assert.Equal(t, expected, s)
			}
		})
	}

	floatTests := []struct {
		name        string
		rawJson     string
		expectedFoo optional.Of[float64]
	}{
		{
			name:        "ValidFloat",
			rawJson:     `{"foo":1.05}`,
			expectedFoo: optional.NewWithValue(1.05),
		},
		{
			name:        "NullFloat",
			rawJson:     `{"foo":null}`,
			expectedFoo: optional.New(0.0, false),
		},
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			var s S[float64]
			err := json.Unmarshal([]byte(tt.rawJson), &s)
			expected := S[float64]{
				Foo: tt.expectedFoo,
			}
			if assert.NoError(t, err) {
				assert.Equal(t, expected, s)
			}
		})
	}
}

func TestOf_Scan(t *testing.T) {
	stringTests := []struct {
		name     string
		fooInDb  any
		expected optional.Of[string]
	}{
		{
			name:     "ValidString",
			fooInDb:  "bar",
			expected: optional.NewWithValue("bar"),
		},
		{
			name:     "NullString",
			fooInDb:  nil,
			expected: optional.New("", false),
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			rows := sqlmock.NewRows([]string{"foo"}).AddRow(tt.fooInDb)
			mock.ExpectQuery("SELECT").WillReturnRows(rows)

			rs, _ := db.Query("SELECT")
			defer rs.Close()

			for rs.Next() {
				var foo optional.Of[string]
				_ = rs.Scan(&foo)
				assert.Equal(t, tt.expected, foo)
			}
		})
	}

	floatTests := []struct {
		name     string
		fooInDb  any
		expected optional.Of[float64]
	}{
		{
			name:     "ValidFloat",
			fooInDb:  1.05,
			expected: optional.NewWithValue(1.05),
		},
		{
			name:     "NullFloat",
			fooInDb:  nil,
			expected: optional.New(0.0, false),
		},
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			defer db.Close()

			rows := sqlmock.NewRows([]string{"foo"}).AddRow(tt.fooInDb)
			mock.ExpectQuery("SELECT").WillReturnRows(rows)

			rs, _ := db.Query("SELECT")
			defer rs.Close()

			for rs.Next() {
				var foo optional.Of[float64]
				_ = rs.Scan(&foo)
				assert.Equal(t, tt.expected, foo)
			}
		})
	}
}

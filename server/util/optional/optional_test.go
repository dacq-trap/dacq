package optional

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf_MarshalJSON(t *testing.T) {
	type S[T optionalType] struct {
		Foo Of[T] `json:"foo"`
	}

	stringTests := []struct {
		name     string
		foo      Of[string]
		expected string
	}{
		{
			name:     "ValidString",
			foo:      NewWithValue("bar"),
			expected: `{"foo":"bar"}`,
		},
		{
			name:     "NullString",
			foo:      New("", false),
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
		foo      Of[float64]
		expected string
	}{
		{
			name:     "ValidFloat",
			foo:      NewWithValue(1.05),
			expected: `{"foo":1.05}`,
		},
		{
			name:     "NullFloat",
			foo:      New(0.0, false),
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
	type S[T optionalType] struct {
		Foo Of[T] `json:"foo"`
	}

	stringTests := []struct {
		name        string
		rawJson     string
		expectedFoo Of[string]
	}{
		{
			name:        "ValidString",
			rawJson:     `{"foo":"bar"}`,
			expectedFoo: NewWithValue("bar"),
		},
		{
			name:        "NullString",
			rawJson:     `{"foo":null}`,
			expectedFoo: New("", false),
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
		expectedFoo Of[float64]
	}{
		{
			name:        "ValidFloat",
			rawJson:     `{"foo":1.05}`,
			expectedFoo: NewWithValue(1.05),
		},
		{
			name:        "NullFloat",
			rawJson:     `{"foo":null}`,
			expectedFoo: New(0.0, false),
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
		expected Of[string]
	}{
		{
			name:     "ValidString",
			fooInDb:  "bar",
			expected: NewWithValue("bar"),
		},
		{
			name:     "NullString",
			fooInDb:  nil,
			expected: New("", false),
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
				var foo Of[string]
				rs.Scan(&foo)
				assert.Equal(t, tt.expected, foo)
			}
		})
	}

	floatTests := []struct {
		name     string
		fooInDb  any
		expected Of[float64]
	}{
		{
			name:     "ValidFloat",
			fooInDb:  1.05,
			expected: NewWithValue(1.05),
		},
		{
			name:     "NullFloat",
			fooInDb:  nil,
			expected: New(0.0, false),
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
				var foo Of[float64]
				rs.Scan(&foo)
				assert.Equal(t, tt.expected, foo)
			}
		})
	}
}

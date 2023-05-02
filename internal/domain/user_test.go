package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPagination(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		want   Pagination
	}{
		{
			name: "no_params_default_pagination",
			want: Pagination{limit: DefaultPaginationLimit, Offset: 0},
		},
		{
			name:  "limit_only",
			limit: 18,
			want:  Pagination{limit: 18, Offset: 0},
		},
		{
			name:   "offset_only",
			offset: 18,
			want:   Pagination{limit: DefaultPaginationLimit, Offset: 18},
		},
		{
			name:  "above_max_limit",
			limit: MaxPaginationLimit + 1,
			want:  Pagination{limit: MaxPaginationLimit, Offset: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewPagination(tt.limit, tt.offset)
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestFields_AddIfNotNil(t *testing.T) {
	name := "Christopher"

	tests := []struct {
		name     string
		key      Field
		value    *string
		f        Fields
		expected Fields
	}{

		{name: "nil_field", key: "first_name", value: nil, f: Fields{}, expected: Fields{}},
		{name: "non_nil_field", key: "first_name", value: &name, f: Fields{}, expected: Fields{"first_name": name}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.AddIfNotNil(tt.key, tt.value)
			assert.EqualValues(t, tt.expected, tt.f)
		})
	}
}

func Test_valueIfNotEmpty(t *testing.T) {
	name := "Christopher"
	tests := []struct {
		name  string
		value string
		want  *string
	}{
		{name: "empty_string", value: "", want: nil},
		{name: "non_empty_string", value: "Christopher", want: &name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, valueIfNotEmpty(tt.value))
		})
	}
}

func TestNewPagination1(t *testing.T) {
	tests := []struct {
		name   string
		limit  int
		offset int
		want   Pagination
	}{
		{name: "default_pagination", want: Pagination{limit: DefaultPaginationLimit, Offset: 0}},
		{name: "limit_in_range", limit: MaxPaginationLimit, want: Pagination{limit: MaxPaginationLimit, Offset: 0}},
		{name: "limit_above_maximum", limit: MaxPaginationLimit + 1, want: Pagination{limit: MaxPaginationLimit, Offset: 0}},
		{name: "offset", offset: 10, want: Pagination{limit: DefaultPaginationLimit, Offset: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPagination(tt.limit, tt.offset))
		})
	}
}

func Test_max(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "a_greater_than_b", a: 4, b: 5, want: 4},
		{name: "a_less_than_b", a: 5, b: 4, want: 4},
		{name: "a_equal_to_b", a: 10, b: 10, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, min(tt.a, tt.b))
		})
	}
}

func TestNewUser(t *testing.T) {
	tests := []struct {
		name string

		email string

		want    User
		wantErr error
	}{
		{name: "email_is_required", email: "", want: User{}, wantErr: ErrEmailRequired},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser("", "", "", "", tt.email, "")
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

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

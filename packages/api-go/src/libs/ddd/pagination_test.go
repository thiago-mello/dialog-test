package ddd

import "testing"

func TestPaginatedQuery_GetPageSize(t *testing.T) {
	type fields struct {
		PageSize   int32
		LastSeenId string
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "Zero page size should return 15",
			fields: fields{
				PageSize:   0,
				LastSeenId: "",
			},
			want: 15,
		},
		{
			name: "Negative page size should return 15",
			fields: fields{
				PageSize:   -1,
				LastSeenId: "",
			},
			want: 15,
		},
		{
			name: "Positive page size should return the same value",
			fields: fields{
				PageSize:   10,
				LastSeenId: "",
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PaginatedQuery{
				PageSize:   tt.fields.PageSize,
				LastSeenId: tt.fields.LastSeenId,
			}
			if got := p.GetPageSize(); got != tt.want {
				t.Errorf("PaginatedQuery.GetPageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

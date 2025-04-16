package ddd

type PaginatedQuery struct {
	PageSize   int32  `query:"page_size" validate:"omitempty,min=1"`
	LastSeenId string `query:"last_seen_id" validate:"omitempty,min=1,uuid"`
}

func (p PaginatedQuery) GetPageSize() int32 {
	if p.PageSize == 0 {
		return 15
	}

	return p.PageSize
}

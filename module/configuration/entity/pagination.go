package entity

type OffsetPagination struct {
	Limit  uint32 `json:"limit"`
	Offset uint32 `json:"offset"`
	Total  uint32 `json:"total"`
}

type PaginationResponse struct {
	OffsetPagination *OffsetPagination `json:"offset_pagination,omitempty"`
}

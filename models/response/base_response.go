package response

// Pagination Pagination
type Pagination struct {
	Offset int   `json:"offset,omitempty"`
	Page   int   `json:"page,omitempty"`
	Limit  int   `json:"limit,omitempty"`
	Total  int64 `json:"total,omitempty"`
	Pages  int   `json:"pages,omitempty"`
}

// NewPagePagination NewPagination
func NewPagePagination(page, limit int, total int64) *Pagination {
	offset := (page - 1) * limit

	p := &Pagination{
		Page:   page,
		Offset: offset,
		Limit:  limit,
		Total:  total,
	}

	return p
}

// NewOffsetPagination NewPagination
func NewOffsetPagination(offset, limit int, total int64) *Pagination {
	page := offset/limit + 1
	// TODO:
	// pages := total/100 + 1

	p := &Pagination{
		Page:   page,
		Offset: offset,
		Limit:  limit,
		Total:  total,
		Pages:  100,
	}

	return p
}

// BasePageResponse BasePageResponse
type BasePageResponse struct {
	Items      []interface{} `json:"data"`
	Pagination *Pagination   `json:"pagination"`
	// Total      uint64        `json:"total"`
}

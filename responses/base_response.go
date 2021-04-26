package response

// Pagination
type Pagination struct {
	Offset int   `json:"offset,omitempty"`
	Page   int   `json:"page,omitempty"`
	Limit  int   `json:"limit,omitempty"`
	Total  int64 `json:"total,omitempty"`
	Pages  int   `json:"pages,omitempty"`
}

// BasePage
type BasePage struct {
	Items      []interface{} `json:"data"`
	Pagination *Pagination   `json:"pagination"`
}

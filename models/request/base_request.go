package uiamsdk

// BaseRequest BaseRequest
type BaseRequest struct {
	RealmID    string `json:"realm_id"`
	IdentityID string `json:"Identity_id"`
	// Authorization string `json:"authorization"`
}

// BasePageRequest BasePageRequest
type BasePageRequest struct {
	BaseRequest

	Offset uint64 `form:"offset"`
	Limit  uint64 `form:"limit"`
}

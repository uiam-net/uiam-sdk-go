package uiamsdk

// BaseRequest BaseRequest
type BaseRequest struct {
	IdentityID   uint64 `json:"identity_id"`
	IdentityUUID string `json:"identity_uuid"`
	RealmID      uint64 `json:"realm_id"`
	RealmSlug    string `json:"realm_slug"`
	MerchantID   uint64 `json:"-"`
	MerchantSlug string `json:"merchant_slug"`
}

// BasePageRequest BasePageRequest
type BasePageRequest struct {
	BaseRequest

	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

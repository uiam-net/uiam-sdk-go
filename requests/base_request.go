package uiamsdk

// BaseRequest BaseRequest
type BaseRequest struct {
	RealmID      uint64 `json:"realm_id"`
	RealmSlug    string `json:"realm_slug"`
	IdentityID   uint64 `json:"identity_id"`
	IdentityUUID string `json:"identity_uuid"`
	MerchantSlug string `json:"merchant_slug"`
}

// BasePageRequest BasePageRequest
type BasePageRequest struct {
	BaseRequest

	Offset uint64 `form:"offset"`
	Limit  uint64 `form:"limit"`
}

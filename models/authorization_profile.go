package uiamsdk

// ==== Mixin Operations ===== //
// ==== mixin Operations ===== //

// MixinUserProfile messenger user entity
type MixinUserProfile struct {
	UserID     string `json:"user_id"`
	UiamNumber string `json:"uiam_number"`
	FullName   string `json:"full_name,omitempty"`
	AvatarURL  string `json:"avatar_url,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	Phone      string `json:"phone,omitempty"`
}

// MixinUserDataProfile MixinUserDataProfile
type MixinUserDataProfile struct {
	Data  *MixinUserProfile `json:"data"`
	Error *MixinAuthError   `json:"error"`
}

// MixinAuthError MixinAuthError
type MixinAuthError struct {
	Status      int32  `json:"status"`
	Code        int32  `json:"code"`
	Description string `json:"description"`
}

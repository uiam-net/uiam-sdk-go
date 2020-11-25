package uiamsdk

import "time"

// Realm Realm
type Realm struct {
	ID          uint32          `json:"realm_id" `
	Name        string          `json:"name"`
	Slug        string          `json:"slug"`
	Description string          `json:"description"`
	Status      RealmStatusEnum `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// RealmStatusEnum 枚举
type RealmStatusEnum string

const (
	// RealmStatusEnumNormal 正常
	RealmStatusEnumNormal RealmStatusEnum = "normal"
	// RealmStatusEnumDisabled 下架
	RealmStatusEnumDisabled RealmStatusEnum = "disabled"
)

func (e RealmStatusEnum) String() string {
	switch e {
	case RealmStatusEnumNormal:
		return "normal"
	case RealmStatusEnumDisabled:
		return "disabled"
	default:
		return "unkonwn"
		// return fmt.Sprintf("%f", e)
	}
}

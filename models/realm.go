package uiamsdk

import (
	"time"
)

// Realm Realm
type Realm struct {
	ID           uint32          `gorm:"COLUMN:id;PRIMARY_KEY;NOT NULL;TYPE:INT(11)"`
	UUID         string          `gorm:"COLUMN:uuid;NOT NULL;TYPE:VARCHAR(36)"`
	Slug         string          `gorm:"COLUMN:slug;NOT NULL;TYPE:VARCHAR(36)"`
	MerchantUUID string          `gorm:"COLUMN:merchant_uuid;NOT NULL;TYPE:VARCHAR(36)"`
	Name         string          `gorm:"COLUMN:name;NOT NULL;TYPE:VARCHAR(36)"`
	AuthModel    string          `gorm:"COLUMN:auth_model;NOT NULL;TYPE:VARCHAR(255)"`
	Description  string          `gorm:"COLUMN:description;NOT NULL;TYPE:VARCHAR(255)"`
	Remark       string          `gorm:"COLUMN:description;NOT NULL;TYPE:VARCHAR(255)"`
	Status       RealmStatusEnum `gorm:"NOT NULL"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName TableName
func (Realm) TableName() string {
	return "realms"
}

// ================== EUUM ================== //

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
		return ""
	}
}

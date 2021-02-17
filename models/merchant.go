package uiamsdk

import (
	"time"
)

// Merchant Merchant
type Merchant struct {
	ID          uint32             `json:"id" gorm:"COLUMN:id;PRIMARY_KEY;NOT NULL;TYPE:INT(11)"`
	UUID        string             `json:"uuid" gorm:"COLUMN:uuid;NOT NULL;TYPE:VARCHAR(36)"`
	Slug        string             `json:"slug" gorm:"COLUMN:slug;NOT NULL;TYPE:VARCHAR(36)"`
	Name        string             `json:"name" gorm:"COLUMN:name;NOT NULL;TYPE:VARCHAR(36)"`
	Description string             `json:"description" gorm:"COLUMN:description;NOT NULL;TYPE:VARCHAR(255)"`
	Remark      string             `json:"remark" gorm:"COLUMN:description;NOT NULL;TYPE:VARCHAR(255)"`
	Status      MerchantStatusEnum `json:"status" gorm:"NOT NULL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName TableName
func (Merchant) TableName() string {
	return "merchants"
}

// MerchantStatusEnum 枚举
type MerchantStatusEnum string

const (
	// MerchantStatusEnumNormal 正常
	MerchantStatusEnumNormal MerchantStatusEnum = "normal"
	// MerchantStatusEnumDisabled 下架
	MerchantStatusEnumDisabled MerchantStatusEnum = "disabled"
)

func (e MerchantStatusEnum) String() string {
	switch e {
	case MerchantStatusEnumNormal:
		return "normal"
	case MerchantStatusEnumDisabled:
		return "disabled"
	default:
		return ""
	}
}

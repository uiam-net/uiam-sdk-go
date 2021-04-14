package uiamsdk

import (
	"time"
)

// Merchant Merchant
type Merchant struct {
	ID          uint32             `json:"id" gorm:"COLUMN:id;PRIMARY_KEY;NOT NULL;TYPE:INT(11)"`
	UUID        string             `mapstructure:"uuid" json:"uuid" gorm:"COLUMN:uuid;NOT NULL;TYPE:VARCHAR(36)"`
	Type        MerchantTypeEnum   `mapstructure:"type" json:"type,omitempty" gorm:"COLUMN:type;NOT NULL;TYPE:VARCHAR(18)"`
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

// =============== IdentityTypeEnum =============== //

// IdentityTypeEnum 枚举
type MerchantTypeEnum string

const (
	// MerchantTypeEnumPersonal 正常注册用户
	MerchantTypeEnumPersonal MerchantTypeEnum = "personal"
	// MerchantTypeEnumCompany 系统调用使用
	MerchantTypeEnumCompany MerchantTypeEnum = "copmpany"
)

func (e MerchantTypeEnum) String() string {
	switch e {
	case MerchantTypeEnumPersonal:
		return "personal"
	case MerchantTypeEnumCompany:
		return "copmpany"
	default:
		return ""
	}
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

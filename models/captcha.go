package uiamsdk

import (
	"time"
)

// Captcha Captcha
type Captcha struct {
	ID          uint64          `json:"id" gorm:"COLUMN:id; PRIMARY_KEY;NOT NULL; AUTO_INCREMENT"`
	PhoneCode   string          `json:"phone_code" gorm:"COLUMN:phone_code; NOT NULL;"`
	PhoneNumber string          `json:"phone_number" gorm:"COLUMN:phone_number; NOT NULL;"`
	Code        string          `json:"code" gorm:"COLUMN:code; NOT NULL;"`
	Type        CaptchaTypeEnum `json:"type" gorm:"COLUMN:type; NOT NULL;"`
	VerifyCount uint32          `json:"verify_count" gorm:"COLUMN:verify_count; NOT NULL;"`
	ExpiredAt   time.Time       `json:"expired_at" gorm:"COLUMN:expired_at; NOT NULL;"`
	CreatedAt   time.Time       `json:"created_at" gorm:"COLUMN:created_at; NOT NULL;"`
	UpdatedAt   time.Time       `json:"updated_at" gorm:"COLUMN:updated_at; NOT NULL;"`
}

// TableName TableName
func (Captcha) TableName() string {
	return "captchas"
}

// CaptchaTypeEnum CaptchaTypeEnum
type CaptchaTypeEnum string

const (
	// CaptchaTypeEnumLogin offer
	CaptchaTypeEnumLogin CaptchaTypeEnum = "login"
)

func (e CaptchaTypeEnum) String() string {
	switch e {
	case CaptchaTypeEnumLogin:
		return "login"
	default:
		return ""
	}
}

// NewCaptchaTypeEnum 构造函数
func NewCaptchaTypeEnum(captchaCodeTypeEnum string) CaptchaTypeEnum {
	switch captchaCodeTypeEnum {
	case "login":
		return CaptchaTypeEnumLogin
	default:
		return ""
	}
}

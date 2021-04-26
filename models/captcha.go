package uiamsdk

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

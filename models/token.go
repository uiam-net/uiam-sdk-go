package uiamsdk

// SessionTypeEnum 枚举
type SessionTypeEnum string

const (
	SessionTypeEnumToken   SessionTypeEnum = "token"
	SessionTypeEnumCaptcha SessionTypeEnum = "captcha"
	SessionTypeEnumMfa     SessionTypeEnum = "mfa"
)

func (e SessionTypeEnum) String() string {
	switch e {
	case SessionTypeEnumToken:
		return "token"
	case SessionTypeEnumCaptcha:
		return "captcha"
	case SessionTypeEnumMfa:
		return "mfa"
	default:
		return ""
	}
}

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

// LoginTypeEnum 枚举
type LoginTypeEnum string

const (
	LoginTypeEnumUsername LoginTypeEnum = "username"
	LoginTypeEnumPhone    LoginTypeEnum = "phone"
	LoginTypeEnumEmail    LoginTypeEnum = "email"
	LoginTypeEnumSmsCode  LoginTypeEnum = "sms_code"
)

func (e LoginTypeEnum) String() string {
	switch e {
	case LoginTypeEnumUsername:
		return "username"
	case LoginTypeEnumPhone:
		return "phone"
	case LoginTypeEnumEmail:
		return "email"
	case LoginTypeEnumSmsCode:
		return "sms_code"
	default:
		return ""
	}
}

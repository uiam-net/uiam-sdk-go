package uiamsdk

// LoginRequest LoginRequest
type LoginRequest struct {
	BaseRequest

	Type        LoginTypeEnum `json:"type"`
	PhoneCode   string        `json:"phone_code"`
	PhoneNumber string        `json:"phone_number"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
}

// LoginTypeEnum 枚举
type LoginTypeEnum string

const (
	LoginTypeEnumName    LoginTypeEnum = "name"
	LoginTypeEnumPhone   LoginTypeEnum = "phone"
	LoginTypeEnumEmail   LoginTypeEnum = "email"
	LoginTypeEnumSmsCode LoginTypeEnum = "sms_code"
)

func (e LoginTypeEnum) String() string {
	switch e {
	case LoginTypeEnumName:
		return "name"
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

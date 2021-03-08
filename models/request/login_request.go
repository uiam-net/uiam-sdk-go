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
	// LoginTypeEnumID LoginTypeEnumID
	LoginTypeEnumID LoginTypeEnum = "id"
	// LoginTypeEnumPhone LoginTypeEnumPhone
	LoginTypeEnumPhone LoginTypeEnum = "phone"
	// LoginTypeEnumEmail LoginTypeEnumEmail
	LoginTypeEnumEmail LoginTypeEnum = "email"
)

func (e LoginTypeEnum) String() string {
	switch e {
	case LoginTypeEnumID:
		return "id"
	case LoginTypeEnumPhone:
		return "phone"
	case LoginTypeEnumEmail:
		return "email"
	default:
		return ""
	}
}

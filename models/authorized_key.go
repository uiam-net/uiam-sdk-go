package uiamsdk

// =============== AuthKeyTypeEnum =============== //

// AuthkeyTypeEnum 枚举
type AuthkeyTypeEnum string

const (
	AuthkeyTypeEnumLogin AuthkeyTypeEnum = "login"
	AuthkeyTypeEnumApi   AuthkeyTypeEnum = "api"
)

func (e AuthkeyTypeEnum) String() string {
	switch e {
	case AuthkeyTypeEnumLogin:
		return "login"
	case AuthkeyTypeEnumApi:
		return "api"
	default:
		return ""
	}
}

// NewAuthkeyTypeEnum 构造函数
func NewAuthkeyTypeEnum(keyType string) AuthkeyTypeEnum {
	switch keyType {
	case "login":
		return AuthkeyTypeEnumLogin
	case "api":
		return AuthkeyTypeEnumApi
	default:
		return ""
	}
}

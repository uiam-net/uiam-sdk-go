package uiamsdk

// =============== AuthKeyTypeEnum =============== //

// AuthKeyTypeEnum 枚举
type AuthKeyTypeEnum string

const (
	AuthKeyTypeEnumLogin AuthKeyTypeEnum = "login"
	AuthKeyTypeEnumApi   AuthKeyTypeEnum = "api"
)

func (e AuthKeyTypeEnum) String() string {
	switch e {
	case AuthKeyTypeEnumLogin:
		return "login"
	case AuthKeyTypeEnumApi:
		return "api"
	default:
		return ""
	}
}

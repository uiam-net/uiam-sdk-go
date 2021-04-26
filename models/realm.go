package uiamsdk

// ================== EUUM ================== //

// RealmStatusEnum 枚举
type RealmStatusEnum string

const (
	// RealmStatusEnumNormal 正常
	RealmStatusEnumNormal RealmStatusEnum = "normal"
	// RealmStatusEnumDisabled 下架
	RealmStatusEnumDisabled RealmStatusEnum = "disabled"
)

func (e RealmStatusEnum) String() string {
	switch e {
	case RealmStatusEnumNormal:
		return "normal"
	case RealmStatusEnumDisabled:
		return "disabled"
	default:
		return ""
	}
}

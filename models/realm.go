package uiamsdk

// ================== EUUM ================== //

// RealmStatusEnum 枚举
type RealmAuthModelEnum string

const (
	RealmAuthModelEnumRealm      RealmAuthModelEnum = "realm"
	RealmAuthModelEnumUser       RealmAuthModelEnum = "user"
	RealmAuthModelEnumUserDevice RealmAuthModelEnum = "device"
)

func (e RealmAuthModelEnum) String() string {
	switch e {
	case RealmAuthModelEnumRealm:
		return "realm"
	case RealmAuthModelEnumUser:
		return "user"
	case RealmAuthModelEnumUserDevice:
		return "device"
	default:
		return ""
	}
}

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

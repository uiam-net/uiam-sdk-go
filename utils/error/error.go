package error

import "fmt"

// Error Error Interface
type Error interface {
	Error() string
	GetCode() int
	GetMsg() string
}

var (
	// Unrecognized APP
	Unrecognized = AppError{4000, "未知错误"}

	// ReqParamError 客户端错误
	ReqParamError   = AppError{4001, "请求参数错误"}
	ReqForbidden    = AppError{4003, "请求被拒绝"}
	ReqPathNotExist = AppError{4004, "请求路径不存在"}

	// UnknownError 服务端错误
	UnknownError      = AppError{5000, "服务端未知错误"}
	Timeout           = AppError{5001, "服务超时"}
	ServerUnavailable = AppError{5005, "服务当前不可用"}
	// DbQueryError Common Server
	DbQueryError   = AppError{5001, "数据库查询错误"}
	DBSessionError = AppError{5001, "数据库事务错误"}
	// CacheRedisError Redis
	CacheRedisError       = AppError{5010, "REDIS 出错"}
	CacheKeyNotExistError = AppError{5011, "REDIS KEY 不存在"}

	// ========== APP ============== //

	// AppNotExists 业务
	AppNotExists = AppError{50001, "app 不存在"}

	// ========== Identity ============== //

	// AppIdentityError App Identity
	AppIdentityError = AppError{100000, "AppIdentityError"}

	// AuthAccessForbid AuthAccessForbid
	AuthAccessForbid     = AppError{4000, "没有访问权限"}
	AuthExpired          = AppError{4000, "令牌过期"}
	AuthTokenGenError    = AppError{4003, "无法生成签名"}
	AuthTokenInvalid     = AppError{4003, "无效签名 Invalid API token"}
	AuthTokenError       = AppError{4003, "API Token Error"}
	AuthUnknownScheme    = AppError{4003, "Unknown Auth scheme"}
	AuthNotExistError    = AppError{4003, "AuthNotExistError"}
	AuthTokenParseError  = AppError{4003, "AuthTokenParseError"}
	AuthTokenHeaderError = AppError{4003, "AuthTokenHeaderError"} // 不能分两段 scheme value

	// IdentityAlreadyExist IdentityAlreadyExist
	IdentityAlreadyExist         = AppError{50101, "用户已存在"}
	IdentityNotExists            = AppError{52001, "用户不存在"}
	IdentityNamePwdIncorrect     = AppError{52000, "用户名或密码错误"}
	IdentityPhoneIncorrect       = AppError{52002, "手机号码错误"}
	IdentityStatusError          = AppError{52002, "用户不可用"}
	WrongPassword                = AppError{52002, "密码错误"}
	IdentityAlreadyAuthenticated = AppError{52003, "用户已经认证"}
	PhoneNumberNotExist          = AppError{52004, "手机号不存在"}

	// ========== Identity ============== //

	// OAuthMixinAuthInvalid OAuth
	OAuthMixinAuthInvalid  = AppError{5001, "Mixin 无效签名1"}
	OAuthMixinAuth2Invalid = AppError{52001, "Mixin 无效签名2"}
	OAuthMixinAuth3Invalid = AppError{52001, "Mixin 无效签名3"}
	MixinCreateWalletError = AppError{52001, "Mixin 创建钱包出错"}

	// Phone verification
	CaptchaInvalid   = AppError{54001, "验证码非法"}
	CaptchaIncorrect = AppError{54002, "验证码不正确"}
	CaptchaExpired   = AppError{54003, "验证码已过期"}
	CaptchaTimeLimit = AppError{54004, "60s内不能重复获取验证码"}
)

// AppError AppError
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetCode GetCode
func (err AppError) GetCode() int {
	return err.Code
}

// GetMsg GetMsg
func (err AppError) GetMsg() string {
	return err.Message
}

// Error Error
func (err AppError) Error() string {
	return fmt.Sprintf("errorCode: %d, errorMsg %s", err.Code, err.Message)
}

// =========== 构造函数 ================= //
// =========== 构造函数 ================= //

// NewError New returns an error that formats as the given text.
func NewError(text string) Error {
	return &AppError{Unrecognized.Code, text}
}

// NewAppError New returns an error that formats as the given text.
func NewAppError(text string) Error {
	return &AppError{UnknownError.Code, text}
}

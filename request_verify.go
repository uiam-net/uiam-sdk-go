package uiamsdk

// PasswordVerifyRequest PasswordVerifyRequest
type PasswordVerifyRequest struct {
	BaseRequest

	ID       string `json:"id"`
	Password string `json:"password"`
}

// PhoneCodeVerifyRequest PhoneCodeVerifyRequest
type PhoneCodeVerifyRequest struct {
	BasePageRequest

	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

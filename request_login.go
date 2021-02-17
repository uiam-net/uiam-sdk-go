package uiamsdk

// LoginRequest LoginRequest
type LoginRequest struct {
	BaseRequest

	ID          string `json:"id"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

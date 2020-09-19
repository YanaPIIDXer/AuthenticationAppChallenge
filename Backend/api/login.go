package api

// ログイン要求
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// ログイン結果
type LoginResult struct {
	ResultCode int `json:"result_code"`
}

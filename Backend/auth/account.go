package auth

// アカウント
type Account struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

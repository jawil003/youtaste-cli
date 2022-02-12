package models

type CreateProviderLoginRequest struct {
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
}

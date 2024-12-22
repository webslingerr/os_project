package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

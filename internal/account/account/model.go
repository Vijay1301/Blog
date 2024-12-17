package account

type SignUp struct {
	AccountId string `json:"accountId"`
	UserID    string `json:"userId"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

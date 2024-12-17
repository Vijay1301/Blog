package account

import "time"

type SignUp struct {
	AccountId string `json:"accountId"`
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AccountDao struct {
	AccountId string    `bson:"accountId"`
	UserId    string    `bson:"userId"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	CreatedBy string    `bson:"createdBy"`
	UpdatedAt time.Time `bson:"updatedAt"`
	UpdatedBy string    `bson:"updatedBy"`
}

type SignUpRes struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccountId    string `json:"accountId"`
	UserId       string `json:"userId"`
}

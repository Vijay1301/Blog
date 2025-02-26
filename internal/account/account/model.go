package account

import "time"

type SignUp struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountDao struct {
	Id             string    `bson:"id"`
	Email          string    `bson:"email"`
	Username       string    `bson:"username"`
	FullName       string    `bson:"full_name"`
	Password       string    `bson:"password"`
	Bio            string    `bson:"bio"`
	ProfileImage   string    `bson:"profile_image"`
	CoverImage     string    `bson:"cover_image"`
	Website        string    `bson:"website"`
	Location       string    `bson:"location"`
	SocialLinks    []string  `bson:"social_links"`
	FollowerCount  int       `bson:"follower_count"`
	FollowingCount int       `bson:"following_count"`
	JoinedAt       time.Time `bson:"joined_at"`
	LastActive     time.Time `bson:"last_active"`
	IsVerified     bool      `bson:"is_verified"`
	IsPremium      bool      `bson:"is_premium"`
	CreatedAt      time.Time `bson:"created_at"`
	CreatedBy      string    `bson:"created_by"`
	UpdatedAt      time.Time `bson:"updated_at"`
	UpdatedBy      string    `bson:"updated_by"`
}

type SignUpRes struct {
	AccessToken  string `json:"access_Token"`
	RefreshToken string `json:"refresh_Token"`
	Id           string `json:"id"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_Token"`
	RefreshToken string `json:"refresh_Token"`
	Id           string `json:"id"`
}

type Account struct {
	Email          string    `json:"email" bson:"email"`
	Username       string    `json:"username" bson:"username"`
	FullName       string    `json:"full_name" bson:"full_name"`
	Bio            string    `json:"bio" bson:"bio"`
	ProfileImage   string    `json:"profile_image" bson:"profile_image"`
	CoverImage     string    `json:"cover_image" bson:"cover_image"`
	Website        string    `json:"website" bson:"website"`
	Location       string    `json:"location" bson:"location"`
	SocialLinks    []string  `json:"social_links" bson:"social_links"`
	FollowerCount  int       `json:"follower_count" bson:"follower_count"`
	FollowingCount int       `json:"following_count" bson:"following_count"`
	JoinedAt       time.Time `json:"joined_at" bson:"joined_at"`
	LastActive     time.Time `json:"last_active" bson:"last_active"`
	IsVerified     bool      `json:"is_verified" bson:"is_verified"`
	IsPremium      bool      `json:"is_premium" bson:"is_premium"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	CreatedBy      string    `json:"created_by" bson:"created_by"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy      string    `json:"updated_by" bson:"updated_by"`
}

type UpdateAccount struct {
	Email        string   `json:"email"`
	Username     string   `json:"username"`
	FullName     string   `json:"full_name"`
	Bio          string   `json:"bio"`
	ProfileImage string   `json:"profile_image"`
	CoverImage   string   `json:"cover_image"`
	Website      string   `json:"website"`
	Location     string   `json:"location"`
	SocialLinks  []string `json:"social_links"`
}

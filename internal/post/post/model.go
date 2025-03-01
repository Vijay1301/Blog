package post

import (
	"time"
)

type BlogPost struct {
	ID          string                   `json:"id"`
	Title       string                   `json:"title"`
	CoverImage  string                   `json:"cover_image"`
	Content     []map[string]interface{} `json:"content"`
	Tags        []string                 `json:"tags"`
	AuthorID    string                   `json:"author_id"`
	Author      User                     `json:"author"`
	PublishedAt *time.Time               `json:"published_at"`
	ReadTime    int                      `json:"read_time"`
	Visibility  string                   `json:"visibility"`
	Status      int                      `json:"status"`
}

type UpdateBlogPost struct {
	Title       string                   `json:"title"`
	CoverImage  string                   `json:"cover_image"`
	Content     []map[string]interface{} `json:"content"`
	Tags        []string                 `json:"tags"`
	AuthorID    string                   `json:"author_id"`
	Author      User                     `json:"author"`
	PublishedAt *time.Time               `json:"published_at"`
	ReadTime    int                      `json:"read_time"`
	Visibility  string                   `json:"visibility"`
	Status      int                      `json:"status"`
}

type User struct {
	Email        string `json:"email" bson:"email"`
	Username     string `json:"username" bson:"username"`
	FullName     string `json:"full_name" bson:"full_name"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
}

type BlogPostDao struct {
	AccountId   string                   `bson:"account_id"`
	ID          string                   `bson:"id"`
	Title       string                   `bson:"title"`
	CoverImage  string                   `bson:"cover_image"`
	Content     []map[string]interface{} `bson:"content"`
	Tags        []string                 `bson:"tags"`
	AuthorID    string                   `bson:"author_id"`
	Author      User                     `bson:"author"`
	PublishedAt *time.Time               `bson:"published_at"`
	ReadTime    int                      `bson:"read_time"`
	Visibility  string                   `bson:"visibility"`
	Status      int                      `bson:"status"`
	CreatedAt   time.Time                `bson:"created_at"`
	CreatedBy   string                   `bson:"created_by"`
	UpdatedAt   time.Time                `bson:"updated_at"`
	UpdatedBy   string                   `bson:"updated_by"`
}

type GetPost struct {
	ID          string                   `bson:"id" json:"id"`
	Title       string                   `bson:"title" json:"title"`
	CoverImage  string                   `bson:"cover_image" json:"cover_image"`
	Content     []map[string]interface{} `bson:"content" json:"content"`
	Tags        []string                 `bson:"tags" json:"tags"`
	AuthorID    string                   `bson:"author_id" json:"author_id"`
	Author      User                     `bson:"author" json:"author"`
	PublishedAt *time.Time               `bson:"published_at" json:"published_at"`
	ReadTime    int                      `bson:"read_time" json:"read_time"`
	Visibility  string                   `bson:"visibility" json:"visibility"`
	Status      int                      `bson:"status" json:"status"`
	CreatedAt   time.Time                `bson:"created_at" json:"created_at"`
	CreatedBy   string                   `bson:"created_by" json:"created_by"`
	UpdatedAt   time.Time                `bson:"updated_at" json:"updated_at"`
	UpdatedBy   string                   `bson:"updated_by" json:"updated_by"`
}

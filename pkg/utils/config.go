package utils

type Config struct {
	Port     string `json:"port"`
	Database *Mongo `json:"database"`
}

type Mongo struct {
	Url      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
	SkipAuth bool   `json:"skipAuth"`
}

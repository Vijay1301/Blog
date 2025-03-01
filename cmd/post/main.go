package main

import (
	"fmt"
	"os"

	"github.com/blog/poc/pkg/utils"

	apiserver "github.com/blog/poc/internal/post"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	config, errr := utils.PortConfig(env)
	if errr != nil {
		fmt.Println(errr)
		os.Exit(-1)
	}

	server := apiserver.New(env)
	err := server.Init(env)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	server.Serve(config.Port)

}

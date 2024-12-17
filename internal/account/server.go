package account

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/blog/poc/docs"
	"github.com/blog/poc/internal/account/account"
	"github.com/blog/poc/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

const (
	ConfigPath = "pkg/config/"
)

type Server struct {
	server      *fiber.App
	initialized bool
	env         string
}

func New(env string) *Server {
	fiberConfig := fiber.Config{
		AppName: "account-service",
	}

	return &Server{
		initialized: false,
		env:         env,
		server:      fiber.New(fiberConfig),
	}
}

func (s *Server) Init(env string) error {

	//Health Check Endpoint

	filename := env + ".json"
	configFile := filepath.Join(ConfigPath, filename)
	fmt.Println(configFile)
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	var Config *utils.Config
	err = json.Unmarshal(configBytes, &Config)
	if err != nil {
		return err
	}

	s.server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "*",
	}))

	s.server.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Alive")
	})

	s.server.Get("/swagger/*", fiberSwagger.WrapHandler)

	//initializing resources
	databaseClient, err := utils.CreateNewConnection(Config.Database)
	if err != nil {
		return err
	}

	Db := databaseClient.Database("blogs")

	apiV1 := s.server.Group("/api/v1")

	authRouter := apiV1.Group("/account")
	authDao := account.NewDAO(Db)
	authService := account.NewService(authDao)
	authHandler := account.NewHandler(authService)
	authHandler.MountRoutes(authRouter)

	return nil
}

func (s *Server) Serve(port string) {
	s.server.Listen(":" + port)
}

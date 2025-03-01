package post

import (
	"encoding/json"
	"os"
	"path/filepath"

	_ "github.com/blog/poc/docs"
	"github.com/blog/poc/internal/post/post"
	utils "github.com/blog/poc/pkg/utils"
	mongo "github.com/blog/poc/pkg/utils/mongo"
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
		AppName: "Post-service",
	}

	return &Server{
		initialized: false,
		env:         env,
		server:      fiber.New(fiberConfig),
	}
}

func (s *Server) Init(env string) error {

	filename := env + ".json"

	configFile := filepath.Join(ConfigPath, filename)

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

	databaseClient, err := mongo.CreateNewConnection(Config.Database)
	if err != nil {
		return err
	}

	Db := databaseClient.Database("blogs")

	apiV1 := s.server.Group("/api/v1")

	postRouter := apiV1.Group("/blog")
	postDao := post.NewDAO(Db)
	postService := post.NewService(postDao)
	postHandler := post.NewHandler(postService)
	postHandler.Routes(postRouter)

	return nil
}

func (s *Server) Serve(port string) {
	s.server.Listen(":" + port)
}

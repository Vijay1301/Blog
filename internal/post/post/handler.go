package post

import (
	"encoding/json"

	"github.com/blog/poc/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Routes(r fiber.Router) {

	r.Use(middleware.AuthMiddleware)
	{

		r.Post("/post", h.CreatePost)

		r.Put("/post/:id", h.UpadetPost)

		r.Get("/post/:id", h.GetPostById)

		r.Get("/post", h.GetAllPost)

		r.Delete("/post/:id", h.DeletePost)

	}

}

// @Summary	Create Post
// @Description	Create Post
// @Tags Post
// @Accept	json
// @Produce	json
// @Security Bearer
// @Param user body BlogPost true "Create Post Request"
// @Success	200	{string}  string "sucess"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/blog/post [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	var req BlogPost

	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invaild Request",
		})
	}

	err = h.Service.CreatePost(c.Context(), Id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})

}

// @Summary	Update Post
// @Description	Update Post
// @Tags Post
// @Accept	json
// @Produce	json
// @Security Bearer
// @Param user body UpdateBlogPost true "Create Post Request"
// @param id path string true "Post Id"
// @Success	200	{string}  string "sucess"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/blog/post/{id} [put]
func (h *Handler) UpadetPost(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	var req UpdateBlogPost

	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invaild Request",
		})
	}

	PostId := c.Params("id")

	err = h.Service.UpdatePost(c.Context(), Id, req, PostId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}

// @Summary	Get Post By Id
// @Description	Get Post By Id
// @Tags Post
// @Accept	json
// @Produce	json
// @Security Bearer
// @param id path string true "Post Id"
// @Success	200	{string}  string "sucess"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/blog/post/{id} [get]
func (h *Handler) GetPostById(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	PostId := c.Params("id")

	post, err := h.Service.GetPostById(c.Context(), Id, PostId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
	})

}

// @Summary	Get All Post
// @Description	Get All Post
// @Tags Post
// @Accept	json
// @Produce	json
// @Security Bearer
// @Success	200	{string}  string "sucess"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/blog/post [get]
func (h *Handler) GetAllPost(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	post, count, err := h.Service.GetAllPost(c.Context(), Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(post) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"posts": []interface{}{},
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"totalCount": count,
		"posts":      post,
	})

}

// @Summary	Delete Post
// @Description	Delete Post
// @Tags Post
// @Accept	json
// @Produce	json
// @Security Bearer
// @param id path string true "Post Id"
// @Success	200	{string}  string "sucess"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/blog/post/{id} [delete]
func (h *Handler) DeletePost(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	PostId := c.Params("id")

	err := h.Service.DeletePost(c.Context(), Id, PostId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}

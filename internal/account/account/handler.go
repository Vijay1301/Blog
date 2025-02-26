package account

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

	r.Post("/signup", h.signup)
	r.Post("/login", h.Login)

	r.Use(middleware.AuthMiddleware)
	{
		r.Get("/", h.GetAccountById)

		r.Put("/edit", h.UpdateAccount)

	}

}

// @Summary	SignUp
// @Description	SignUp
// @Tags Account
// @Accept	json
// @Produce	json
// @Param user body SignUp true "SignUp Request"
// @Success	200	{object}  SignUpRes
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/account/signup [post]
func (h *Handler) signup(c *fiber.Ctx) error {
	var req SignUp
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invaild Request",
		})
	}

	res, err := h.Service.Signup(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)

}

// @Summary	Login
// @Description	Login
// @Tags Account
// @Accept	json
// @Produce	json
// @Param user body Login true "SignUp Request"
// @Success	200	{object}  LoginResponse
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/account/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var req Login
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invaild Request",
		})
	}

	res, err := h.Service.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(res)

}

// @Summary	Get Account Details
// @Description	Get Account Details
// @Tags Account
// @Accept	json
// @Produce	json
// @Security Bearer
// @Success	200	{object}  Account
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/account [get]
func (h *Handler) GetAccountById(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	account, err := h.Service.GetAccount(c.Context(), Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to fetch data",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})

}

// @Summary	Update Account Details
// @Description	Update Account Details
// @Tags Account
// @Accept	json
// @Produce	json
// @Security Bearer
// @Param user body UpdateAccount true "Update Request"
// @Success	200	{object} string "success"
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/account/edit [put]
func (h *Handler) UpdateAccount(c *fiber.Ctx) error {

	Id, ok := c.Locals("accountId").(string)

	if Id == "" || !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "missing account Id",
		})
	}

	var req UpdateAccount

	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invaild Request",
		})
	}

	err = h.Service.UpdateAccount(c.Context(), Id, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "failed to process",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})

}

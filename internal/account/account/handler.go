package account

import (
	"encoding/json"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, err := h.Service.Signup(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, err := h.Service.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)

}

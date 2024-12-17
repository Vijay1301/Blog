package account

import (
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

// @Summary	SignUp
// @Description	SignUp
// @Tags Account
// @Accept	json
// @Produce	json
// @Param user body SignupRequest true "SignUp Request"
// @Success	200	{object}  SignupResponse
// @Failure	400	{string}  error  "Bad Request"
// @Router	/api/v1/account/signup [post]
func (h *Handler) signup(c *fiber.Ctx) error {
	// var req SignUp
	// err := json.Unmarshal(c.Body(), &req)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	// }

	// res, err := h.Service.Signup(c.Context(), req)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	// }
	// return c.Status(fiber.StatusCreated).JSON(res)
	return nil
}

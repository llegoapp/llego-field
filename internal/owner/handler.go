package owner

import (
	"fields/pkg/apperror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewOwnerHandler(s Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) GetOwner(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return apperror.NewBadRequestError("invalid owner id")
	}
	owner, err := h.service.GetOwner(id)
	if err != nil {
		return err
	}
	return c.JSON(owner)
}

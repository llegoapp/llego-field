package owner

import (
	"fields/pkg/apperror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OwnerHandler struct {
	service OwnerService
}

func NewOwnerHandler(s OwnerService) *OwnerHandler {
	return &OwnerHandler{
		s,
	}
}

func (h *OwnerHandler) GetOwner(c *fiber.Ctx) error {
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

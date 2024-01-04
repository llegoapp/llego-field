package field

import (
	"fields/pkg/apperror"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FieldHandler struct {
	service FieldService
}

func NewFieldHandler(s FieldService) *FieldHandler {
	return &FieldHandler{
		s,
	}
}

func (h *FieldHandler) GetField(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return apperror.NewBadRequestError("invalid field id")
	}
	field, err := h.service.GetField(id)
	if err != nil {
		return err
	}
	return c.JSON(field)
}

func (h *FieldHandler) ListFields(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page number")
	}
	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page size")
	}
	field, count, err := h.service.ListFields(page, pageSize)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"count": count,
		"data":  field,
	})
}

func (h *FieldHandler) ListFieldsByOwnerId(c *fiber.Ctx) error {
	ownerId, err := strconv.Atoi(c.Params("ownerId"))
	if err != nil {
		return apperror.NewBadRequestError("invalid owner id")
	}
	field, err := h.service.ListFieldsByOwnerId(ownerId)
	if err != nil {
		return err
	}
	return c.JSON(field)
}

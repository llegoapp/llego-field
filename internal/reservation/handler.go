package reservation

import (
	"fields/pkg/apperror"
	"fields/pkg/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReservationHandler struct {
	service ReservationService
}

func NewReservationHandler(s ReservationService) *ReservationHandler {
	return &ReservationHandler{
		s,
	}
}

func (h *ReservationHandler) CreateReservation(c *fiber.Ctx) error {
	userId, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return apperror.NewUnauthorizedError("unauthorized")
	}

	var resv Reservation
	if err := c.BodyParser(&resv); err != nil {
		return apperror.NewBadRequestError("invalid request body")
	}
	resv.BookerId = userId.ID

	return c.SendStatus(fiber.StatusCreated)
}

func (h *ReservationHandler) GetReservation(c *fiber.Ctx) error {
	reservationId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return apperror.NewBadRequestError("invalid reservation id")
	}

	resv, err := h.service.GetReservation(reservationId)
	if err != nil {
		return err
	}
	return c.JSON(resv)
}

func (h *ReservationHandler) ListReservation(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page number")
	}
	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page size")
	}
	resv, count, err := h.service.ListReservation(page, pageSize)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"count": count,
		"data":  resv,
	})
}

func (h *ReservationHandler) ListReservationByBookerId(c *fiber.Ctx) error {
	userId, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return apperror.NewUnauthorizedError("unauthorized")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page number")
	}
	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page size")
	}
	resv, count, err := h.service.ListReservationByBookerId(userId.ID, page, pageSize)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"count": count,
		"data":  resv,
	})
}

func (h *ReservationHandler) ListReservationByFieldId(c *fiber.Ctx) error {
	fieldId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return apperror.NewBadRequestError("invalid field id")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page number")
	}
	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return apperror.NewBadRequestError("invalid page size")
	}
	resv, count, err := h.service.ListReservationByFieldId(fieldId, page, pageSize)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"count": count,
		"data":  resv,
	})
}

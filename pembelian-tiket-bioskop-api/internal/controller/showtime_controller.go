package controller

import (
	"pembelian-tiket-bioskop-api/internal/dto"
	"pembelian-tiket-bioskop-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShowtimeController interface {
	CreateShowtime(ctx *fiber.Ctx) error
	GetShowtimes(ctx *fiber.Ctx) error
	GetShowtimeByID(ctx *fiber.Ctx) error
	UpdateShowtime(ctx *fiber.Ctx) error
	DeleteShowtime(ctx *fiber.Ctx) error
}

type ShowtimeControllerImpl struct {
	Log             *logrus.Logger
	ShowtimeService service.ShowtimeService
}

func NewShowtimeController(log *logrus.Logger, svc service.ShowtimeService) ShowtimeController {
	return &ShowtimeControllerImpl{
		Log:             log,
		ShowtimeService: svc,
	}
}

func (c *ShowtimeControllerImpl) CreateShowtime(ctx *fiber.Ctx) error {
	var payload dto.ShowtimeRequest
	if err := ctx.BodyParser(&payload); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	result, err := c.ShowtimeService.CreateShowtime(ctx.UserContext(), &payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.WebResponse[*dto.ShowtimeResponse]{
		Message: "Showtime created successfully",
		Data:    result,
	})
}

func (c *ShowtimeControllerImpl) GetShowtimes(ctx *fiber.Ctx) error {
	result, err := c.ShowtimeService.GetShowtimes(ctx.UserContext())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[[]*dto.ShowtimeResponse]{
		Message: "Showtimes retrieved successfully",
		Data:    result,
	})
}

func (c *ShowtimeControllerImpl) GetShowtimeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Showtime ID is required.")
	}

	result, err := c.ShowtimeService.GetShowtimeByID(ctx.UserContext(), id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[*dto.ShowtimeResponse]{
		Message: "Showtime retrieved successfully",
		Data:    result,
	})
}

func (c *ShowtimeControllerImpl) UpdateShowtime(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Showtime ID is required.")
	}

	var payload dto.ShowtimeRequest
	if err := ctx.BodyParser(&payload); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	result, err := c.ShowtimeService.CreateShowtime(ctx.UserContext(), &payload)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[*dto.ShowtimeResponse]{
		Message: "Showtime updated successfully",
		Data:    result,
	})
}

func (c *ShowtimeControllerImpl) DeleteShowtime(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Showtime ID is required.")
	}

	if err := c.ShowtimeService.DeleteShowtime(ctx.UserContext(), id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[any]{
		Message: "Showtime deleted successfully",
		Data:    nil,
	})
}

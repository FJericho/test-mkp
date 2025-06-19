package controller

import (
	"pembelian-tiket-bioskop-api/internal/dto"
	"pembelian-tiket-bioskop-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type StudioController interface {
	CreateStudio(ctx *fiber.Ctx) error
	GetStudios(ctx *fiber.Ctx) error
	DeleteStudio(ctx *fiber.Ctx) error
}

type StudioControllerImpl struct {
	Log           *logrus.Logger
	StudioService service.StudioService
}

func NewStudioController(log *logrus.Logger, studioService service.StudioService) StudioController {
	return &StudioControllerImpl{
		Log:           log,
		StudioService: studioService,
	}
}

func (c *StudioControllerImpl) CreateStudio(ctx *fiber.Ctx) error {
	var payload dto.StudioRequest

	if err := ctx.BodyParser(&payload); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	result, err := c.StudioService.CreateStudio(ctx.UserContext(), &payload)
	if err != nil {
		c.Log.Warnf("Failed to create studio: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.WebResponse[*dto.StudioResponse]{
		Message: "Studio created successfully",
		Data:    result,
	})
}

func (c *StudioControllerImpl) GetStudios(ctx *fiber.Ctx) error {
	result, err := c.StudioService.GetStudios(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to get studios: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[[]*dto.StudioResponse]{
		Message: "Studios retrieved successfully",
		Data:    result,
	})
}

func (c *StudioControllerImpl) DeleteStudio(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Studio ID is required.")
	}

	err := c.StudioService.DeleteStudio(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete studio: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[any]{
		Message: "Studio deleted successfully",
		Data:    nil,
	})
}

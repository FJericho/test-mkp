package controller

import (
	"pembelian-tiket-bioskop-api/internal/dto"
	"pembelian-tiket-bioskop-api/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FilmController interface {
	CreateFilm(ctx *fiber.Ctx) error
	GetFilms(ctx *fiber.Ctx) error
	DeleteFilm(ctx *fiber.Ctx) error
}

type FilmControllerImpl struct {
	Log         *logrus.Logger
	FilmService service.FilmService
}

func NewFilmController(log *logrus.Logger, filmService service.FilmService) FilmController {
	return &FilmControllerImpl{
		Log:         log,
		FilmService: filmService,
	}
}

func (c *FilmControllerImpl) CreateFilm(ctx *fiber.Ctx) error {
	var payload dto.FilmRequest

	if err := ctx.BodyParser(&payload); err != nil {
		c.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	result, err := c.FilmService.CreateFilm(ctx.UserContext(), &payload)
	if err != nil {
		c.Log.Warnf("Failed to create film: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.WebResponse[*dto.FilmResponse]{
		Message: "Film created successfully",
		Data:    result,
	})
}

func (c *FilmControllerImpl) GetFilms(ctx *fiber.Ctx) error {
	result, err := c.FilmService.GetFilms(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to get films: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[[]*dto.FilmResponse]{
		Message: "Films retrieved successfully",
		Data:    result,
	})
}

func (c *FilmControllerImpl) DeleteFilm(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Film ID is required.")
	}

	err := c.FilmService.DeleteFilm(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete film: %+v", err)
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.WebResponse[any]{
		Message: "Film deleted successfully",
		Data:    nil,
	})
}

package service

import (
	"context"
	"pembelian-tiket-bioskop-api/internal/dto"
	"pembelian-tiket-bioskop-api/internal/helper"
	"pembelian-tiket-bioskop-api/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FilmService interface {
	CreateFilm(ctx context.Context, film *dto.FilmRequest) (*dto.FilmResponse, error)
	GetFilms(ctx context.Context) ([]*dto.FilmResponse, error)
	DeleteFilm(ctx context.Context, id string) error
}

type FilmServiceImpl struct {
	Log            *logrus.Logger
	FilmRepository repository.FilmRepository
	Validate       *validator.Validate
}

func NewFilmService(log *logrus.Logger, repo repository.FilmRepository, validate *validator.Validate) FilmService {
	return &FilmServiceImpl{
		FilmRepository: repo,
		Log:            log,
		Validate:       validate,
	}
}

func (f *FilmServiceImpl) CreateFilm(ctx context.Context, film *dto.FilmRequest) (*dto.FilmResponse, error) {
	err := f.Validate.Struct(film)
	if err != nil {
		f.Log.Warnf("Invalid validation: %+v", err)
		errorResponse := helper.GenerateValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, errorResponse)
	}

	exist, err := f.FilmRepository.CheckFilmIfExist(ctx, film.Title)
	if exist {
		f.Log.Warn("Film already exist")
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Film already exist")
	}

	createdFilm, err := f.FilmRepository.CreateFilm(ctx, film.ToEntity())
	if err != nil {
		f.Log.WithError(err).Error("Failed to create film")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create film")
	}

	return &dto.FilmResponse{
		ID:          createdFilm.ID,
		Title:       createdFilm.Title,
		Genre:       createdFilm.Genre,
		Duration:    createdFilm.Duration,
		Description: createdFilm.Description,
	}, nil
}

func (f *FilmServiceImpl) DeleteFilm(ctx context.Context, id string) error {
	err := f.FilmRepository.DeleteFilm(ctx, id)
	if err != nil {
		f.Log.WithError(err).Error("Failed to delete film")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete film")
	}

	return nil
}

func (f *FilmServiceImpl) GetFilms(ctx context.Context) ([]*dto.FilmResponse, error) {
	res, err := f.FilmRepository.GetFilms(ctx)
	if err != nil {
		f.Log.WithError(err).Error("Failed to fetch films")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve films")
	}

	var responses []*dto.FilmResponse
	for _, film := range res {
		responses = append(responses, &dto.FilmResponse{
			ID:          film.ID,
			Title:       film.Title,
			Genre:       film.Genre,
			Duration:    film.Duration,
			Description: film.Description,
		})
	}

	return responses, nil
}

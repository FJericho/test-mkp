package service

import (
	"context"
	"pembelian-tiket-bioskop-api/internal/dto"
	"pembelian-tiket-bioskop-api/internal/entity"
	"pembelian-tiket-bioskop-api/internal/helper"
	"pembelian-tiket-bioskop-api/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ShowtimeService interface {
	CreateShowtime(ctx context.Context, s *dto.ShowtimeRequest) (*dto.ShowtimeResponse, error)
	GetShowtimes(ctx context.Context) ([]*dto.ShowtimeResponse, error)
	GetShowtimeByID(ctx context.Context, id string) (*dto.ShowtimeResponse, error)
	DeleteShowtime(ctx context.Context, id string) error
}

type ShowtimeServiceImpl struct {
	Log          *logrus.Logger
	ShowtimeRepo repository.ShowtimeRepository
	Validate     *validator.Validate
}

func NewShowtimeService(log *logrus.Logger, repo repository.ShowtimeRepository, validate *validator.Validate) ShowtimeService {
	return &ShowtimeServiceImpl{
		Log:          log,
		ShowtimeRepo: repo,
		Validate:     validate,
	}
}

func (s *ShowtimeServiceImpl) CreateShowtime(ctx context.Context, sh *dto.ShowtimeRequest) (*dto.ShowtimeResponse, error) {
	if err := s.Validate.Struct(sh); err != nil {
		s.Log.Warnf("Validation error: %+v", err)
		errorResponse := helper.GenerateValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, errorResponse)
	}

	created, err := s.ShowtimeRepo.CreateShowtime(ctx, sh.ToEntity())
	if err != nil {
		s.Log.WithError(err).Error("Failed to create showtime")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create showtime")
	}

	return &dto.ShowtimeResponse{
		ID: created.ID,
		Film: dto.FilmResponse{
			ID:          created.Film.ID,
			Title:       created.Film.Title,
			Genre:       created.Film.Genre,
			Description: created.Film.Description,
			Duration:    created.Film.Duration,
		},
		Studio: dto.StudioResponse{
			ID:      created.Studio.ID,
			Name:    created.Studio.Name,
			Address: created.Studio.Address,
		},
		StartTime: created.StartTime,
		EndTime:   created.EndTime,
	}, nil
}

func (s *ShowtimeServiceImpl) GetShowtimes(ctx context.Context) ([]*dto.ShowtimeResponse, error) {
	res, err := s.ShowtimeRepo.GetShowtimes(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to fetch showtimes")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve showtimes")
	}

	var responses []*dto.ShowtimeResponse
	for _, sh := range res {
		responses = append(responses, &dto.ShowtimeResponse{
			ID:        sh.ID,
			StartTime: sh.StartTime,
			EndTime:   sh.EndTime,
			Film: dto.FilmResponse{
				ID:          sh.Film.ID,
				Title:       sh.Film.Title,
				Genre:       sh.Film.Genre,
				Description: sh.Film.Description,
				Duration:    sh.Film.Duration,
			},
			Studio: dto.StudioResponse{
				ID:      sh.Studio.ID,
				Name:    sh.Studio.Name,
				Address: sh.Studio.Address,
			},
		})
	}
	return responses, nil
}

func (s *ShowtimeServiceImpl) GetShowtimeByID(ctx context.Context, id string) (*dto.ShowtimeResponse, error) {
	sh, err := s.ShowtimeRepo.GetShowtimeByID(ctx, id)
	if err != nil {
		s.Log.WithError(err).Error("Showtime not found")
		return nil, fiber.NewError(fiber.StatusNotFound, "Showtime not found")
	}

	return &dto.ShowtimeResponse{
		ID: sh.ID,
		Film: dto.FilmResponse{
			ID:          sh.Film.ID,
			Title:       sh.Film.Title,
			Genre:       sh.Film.Genre,
			Description: sh.Film.Description,
			Duration:    sh.Film.Duration,
		},
		Studio: dto.StudioResponse{
			ID:      sh.Studio.ID,
			Name:    sh.Studio.Name,
			Address: sh.Studio.Address,
		},
		StartTime: sh.StartTime,
		EndTime:   sh.EndTime,
	}, nil
}

func (s *ShowtimeServiceImpl) UpdateShowtime(ctx context.Context, id string, sh *entity.Showtime) (*dto.ShowtimeResponse, error) {
	if err := s.Validate.Struct(sh); err != nil {
		s.Log.Warnf("Validation error: %+v", err)
		errorResponse := helper.GenerateValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, errorResponse)
	}

	updated, err := s.ShowtimeRepo.UpdateShowtime(ctx, id, sh)
	if err != nil {
		s.Log.WithError(err).Error("Failed to update showtime")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update showtime")
	}

	return &dto.ShowtimeResponse{
		ID: updated.ID,
		Film: dto.FilmResponse{
			ID:          updated.Film.ID,
			Title:       updated.Film.Title,
			Genre:       updated.Film.Genre,
			Description: updated.Film.Description,
			Duration:    updated.Film.Duration,
		},
		Studio: dto.StudioResponse{
			ID:      updated.Studio.ID,
			Name:    updated.Studio.Name,
			Address: updated.Studio.Address,
		},
		StartTime: updated.StartTime,
		EndTime:   updated.EndTime,
	}, nil
}

func (s *ShowtimeServiceImpl) DeleteShowtime(ctx context.Context, id string) error {
	if err := s.ShowtimeRepo.DeleteShowtime(ctx, id); err != nil {
		s.Log.WithError(err).Error("Failed to delete showtime")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete showtime")
	}
	return nil
}

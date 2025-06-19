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

type StudioService interface {
	CreateStudio(ctx context.Context, studio *dto.StudioRequest) (*dto.StudioResponse, error)
	GetStudios(ctx context.Context) ([]*dto.StudioResponse, error)
	DeleteStudio(ctx context.Context, id string) error
}

type StudioServiceImpl struct {
	Log              *logrus.Logger
	StudioRepository repository.StudioRepository
	Validate         *validator.Validate
}

func NewStudioService(log *logrus.Logger, repo repository.StudioRepository, validate *validator.Validate) StudioService {
	return &StudioServiceImpl{
		StudioRepository: repo,
		Log:              log,
		Validate:         validate,
	}
}

func (s *StudioServiceImpl) CreateStudio(ctx context.Context, studio *dto.StudioRequest) (*dto.StudioResponse, error) {
	err := s.Validate.Struct(studio)
	if err != nil {
		s.Log.Warnf("Validation failed: %+v", err)
		errorResponse := helper.GenerateValidationErrors(err)
		return nil, fiber.NewError(fiber.StatusBadRequest, errorResponse)
	}

	exist, err := s.StudioRepository.CheckStudioIfExist(ctx, studio.Name)
	if exist {
		s.Log.Warn("Studio already exist")
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, "Studio already exist")
	}

	createdStudio, err := s.StudioRepository.CreateStudio(ctx, studio.ToEntity())
	if err != nil {
		s.Log.WithError(err).Error("Failed to create studio")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create studio")
	}

	return &dto.StudioResponse{
		ID:      createdStudio.ID,
		Name:    createdStudio.Name,
		Address: createdStudio.Address,
	}, nil
}

func (s *StudioServiceImpl) GetStudios(ctx context.Context) ([]*dto.StudioResponse, error) {
	res, err := s.StudioRepository.GetStudios(ctx)
	if err != nil {
		s.Log.WithError(err).Error("Failed to fetch studios")
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve studios")
	}

	var responses []*dto.StudioResponse
	for _, studio := range res {
		responses = append(responses, &dto.StudioResponse{
			ID:      studio.ID,
			Name:    studio.Name,
			Address: studio.Address,
		})
	}

	return responses, nil
}

func (s *StudioServiceImpl) DeleteStudio(ctx context.Context, id string) error {
	err := s.StudioRepository.DeleteStudio(ctx, id)
	if err != nil {
		s.Log.WithError(err).Error("Failed to delete studio")
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete studio")
	}

	return nil
}

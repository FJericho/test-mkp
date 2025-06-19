package repository

import (
	"context"
	"pembelian-tiket-bioskop-api/internal/entity"

	"gorm.io/gorm"
)

type StudioRepository interface {
	CreateStudio(ctx context.Context, a *entity.Studio) (*entity.Studio, error)
	GetStudios(ctx context.Context) ([]*entity.Studio, error)
	DeleteStudio(ctx context.Context, id string) error
	CheckStudioIfExist(ctx context.Context, name string) (bool, error)
}

type StudioRepositoryImpl struct {
	DB *gorm.DB
}

func NewStudioRepository(db *gorm.DB) StudioRepository {
	return &StudioRepositoryImpl{
		DB: db,
	}
}

func (s *StudioRepositoryImpl) CreateStudio(ctx context.Context, a *entity.Studio) (*entity.Studio, error) {
	if err := s.DB.WithContext(ctx).Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (s *StudioRepositoryImpl) DeleteStudio(ctx context.Context, id string) error {
	if err := s.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.Studio{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudioRepositoryImpl) GetStudios(ctx context.Context) ([]*entity.Studio, error) {
	var studios []*entity.Studio
	if err := s.DB.WithContext(ctx).Find(&studios).Error; err != nil {
		return nil, err
	}
	return studios, nil
}

func (s *StudioRepositoryImpl) CheckStudioIfExist(ctx context.Context, name string) (bool, error) {
	var studio entity.Studio

	result := s.DB.WithContext(ctx).Where("name = ?", name).First(&studio)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}


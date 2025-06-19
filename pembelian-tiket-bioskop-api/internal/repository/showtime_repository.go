package repository

import (
	"context"
	"fmt"
	"pembelian-tiket-bioskop-api/internal/entity"

	"gorm.io/gorm"
)

type ShowtimeRepository interface {
	CreateShowtime(ctx context.Context, s *entity.Showtime) (*entity.Showtime, error)
	GetShowtimes(ctx context.Context) ([]*entity.Showtime, error)
	GetShowtimeByID(ctx context.Context, id string) (*entity.Showtime, error)
	UpdateShowtime(ctx context.Context, id string, s *entity.Showtime) (*entity.Showtime, error)
	DeleteShowtime(ctx context.Context, id string) error
}

type ShowtimeRepositoryImpl struct {
	DB *gorm.DB
}

func NewShowtimeRepository(db *gorm.DB) ShowtimeRepository {
	return &ShowtimeRepositoryImpl{
		DB: db,
	}
}

func (r *ShowtimeRepositoryImpl) CreateShowtime(ctx context.Context, s *entity.Showtime) (*entity.Showtime, error) {
	if err := r.DB.WithContext(ctx).Create(s).Error; err != nil {
		return nil, err
	}

	var film entity.Film
	if err := r.DB.WithContext(ctx).Where("id = ?", s.FilmID).First(&film).Error; err != nil {
		return nil, fmt.Errorf("film not found: %w", err)
	}

	var studio entity.Studio
	if err := r.DB.WithContext(ctx).Where("id = ?", s.StudioID).First(&studio).Error; err != nil {
		return nil, fmt.Errorf("studio not found: %w", err)
	}

	s.Film = film
	s.Studio = studio

	return s, nil
}

func (r *ShowtimeRepositoryImpl) GetShowtimes(ctx context.Context) ([]*entity.Showtime, error) {
	var showtimes []*entity.Showtime
	if err := r.DB.WithContext(ctx).Preload("Film").Preload("Studio").Find(&showtimes).Error; err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (r *ShowtimeRepositoryImpl) GetShowtimeByID(ctx context.Context, id string) (*entity.Showtime, error) {
	var showtime entity.Showtime
	if err := r.DB.WithContext(ctx).Preload("Film").Preload("Studio").First(&showtime, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (r *ShowtimeRepositoryImpl) UpdateShowtime(ctx context.Context, id string, s *entity.Showtime) (*entity.Showtime, error) {
	var existing entity.Showtime
	if err := r.DB.WithContext(ctx).First(&existing, "id = ?", id).Error; err != nil {
		return nil, err
	}

	existing.FilmID = s.FilmID
	existing.StudioID = s.StudioID
	existing.StartTime = s.StartTime
	existing.EndTime = s.EndTime

	if err := r.DB.WithContext(ctx).Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

func (r *ShowtimeRepositoryImpl) DeleteShowtime(ctx context.Context, id string) error {
	if err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.Showtime{}).Error; err != nil {
		return err
	}
	return nil
}

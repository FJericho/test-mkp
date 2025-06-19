package repository

import (
	"context"
	"pembelian-tiket-bioskop-api/internal/entity"

	"gorm.io/gorm"
)

type FilmRepository interface {
	CreateFilm(ctx context.Context, a *entity.Film) (*entity.Film, error)
	GetFilms(ctx context.Context) ([]*entity.Film, error)
	DeleteFilm(ctx context.Context, id string) error
	CheckFilmIfExist(ctx context.Context, title string) (bool, error)
}

type FilmRepositoryImpl struct {
	DB *gorm.DB
}

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &FilmRepositoryImpl{
		DB: db,
	}
}

func (f *FilmRepositoryImpl) CreateFilm(ctx context.Context, a *entity.Film) (*entity.Film, error) {
	if err := f.DB.WithContext(ctx).Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (f *FilmRepositoryImpl) DeleteFilm(ctx context.Context, id string) error {
	if err := f.DB.WithContext(ctx).Where("id = ?", id).Delete(&entity.Film{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *FilmRepositoryImpl) GetFilms(ctx context.Context) ([]*entity.Film, error) {
	var films []*entity.Film
	if err := f.DB.WithContext(ctx).Find(&films).Error; err != nil {
		return nil, err
	}
	return films, nil
}

func (f *FilmRepositoryImpl) CheckFilmIfExist(ctx context.Context, title string) (bool, error) {
	var film entity.Film

	result := f.DB.WithContext(ctx).Where("title = ?", title).First(&film)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

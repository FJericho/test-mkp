package dto

import "pembelian-tiket-bioskop-api/internal/entity"

type FilmResponse struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Duration    int    `json:"duration,omitempty"`
	Description string `json:"description,omitempty"`
}

type FilmRequest struct {
	Title       string `json:"title" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	Duration    int    `json:"duration" validate:"required,min=1"`
	Description string `json:"description" validate:"required"`
}

func (r *FilmRequest) ToEntity() *entity.Film {
	return &entity.Film{
		Title:       r.Title,
		Genre:       r.Genre,
		Duration:    r.Duration,
		Description: r.Description,
	}
}

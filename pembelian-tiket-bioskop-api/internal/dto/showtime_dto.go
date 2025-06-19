package dto

import (
	"pembelian-tiket-bioskop-api/internal/entity"
	"time"
)

type ShowtimeResponse struct {
	ID        string        `json:"id,omitempty"`
	StartTime time.Time     `json:"start_time,omitempty"`
	EndTime   time.Time     `json:"end_time,omitempty"`
	Film      FilmResponse  `json:"film"`
	Studio    StudioResponse `json:"studio"`
}

type ShowtimeRequest struct {
	FilmID    string    `json:"film_id" validate:"required"`
	StudioID  string    `json:"studio_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
}

func (r *ShowtimeRequest) ToEntity() *entity.Showtime {
	return &entity.Showtime{
		FilmID:    r.FilmID,
		StudioID:  r.StudioID,
		StartTime: r.StartTime,
		EndTime:   r.EndTime,
	}
}

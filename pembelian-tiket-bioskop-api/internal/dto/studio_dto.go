package dto

import "pembelian-tiket-bioskop-api/internal/entity"

type StudioResponse struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type StudioRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (r *StudioRequest) ToEntity() *entity.Studio {
	return &entity.Studio{
		Name:    r.Name,
		Address: r.Address,
	}
}

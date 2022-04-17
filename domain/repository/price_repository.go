package repository

import (
	"cargo-rest-api/domain/entity"
)

// PriceRepository is an interface.
type PriceRepository interface {
	SavePrice(tour *entity.Price) (*entity.Price, map[string]string, error)
	UpdatePrice(UUID string, tour *entity.Price) (*entity.Price, map[string]string, error)
	DeletePrice(UUID string) error
	GetPrice(UUID string) (*entity.Price, error)
	GetPrices(parameters *Parameters) ([]*entity.Price, *Meta, error)
}

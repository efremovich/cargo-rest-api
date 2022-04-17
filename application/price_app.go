package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type priceApp struct {
	tr repository.PriceRepository
}

// priceApp implement the PriceAppInterface.
var _ PriceAppInterface = &priceApp{}

// PriceAppInterface is an interface.
type PriceAppInterface interface {
	SavePrice(*entity.Price) (*entity.Price, map[string]string, error)
	UpdatePrice(UUID string, price *entity.Price) (*entity.Price, map[string]string, error)
	DeletePrice(UUID string) error
	GetPrices(p *repository.Parameters) ([]*entity.Price, *repository.Meta, error)
	GetPrice(UUID string) (*entity.Price, error)
}

func (t priceApp) SavePrice(price *entity.Price) (*entity.Price, map[string]string, error) {
	return t.tr.SavePrice(price)
}

func (t priceApp) UpdatePrice(UUID string, price *entity.Price) (*entity.Price, map[string]string, error) {
	return t.tr.UpdatePrice(UUID, price)
}

func (t priceApp) DeletePrice(UUID string) error {
	return t.tr.DeletePrice(UUID)
}

func (t priceApp) GetPrices(p *repository.Parameters) ([]*entity.Price, *repository.Meta, error) {
	return t.tr.GetPrices(p)
}

func (t priceApp) GetPrice(UUID string) (*entity.Price, error) {
	return t.tr.GetPrice(UUID)
}

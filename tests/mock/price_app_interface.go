package mock

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

// PriceAppInterface is a mock of application.PriceAppInterface.
type PriceAppInterface struct {
	SavePriceFn   func(*entity.Price) (*entity.Price, map[string]string, error)
	UpdatePriceFn func(string, *entity.Price) (*entity.Price, map[string]string, error)
	DeletePriceFn func(UUID string) error
	GetPricesFn   func(params *repository.Parameters) ([]*entity.Price, *repository.Meta, error)
	GetPriceFn    func(UUID string) (*entity.Price, error)
}

// SavePrice calls the SavePriceFn.
func (u *PriceAppInterface) SavePrice(price *entity.Price) (*entity.Price, map[string]string, error) {
	return u.SavePriceFn(price)
}

// UpdatePrice calls the UpdatePriceFn.
func (u *PriceAppInterface) UpdatePrice(uuid string, price *entity.Price) (*entity.Price, map[string]string, error) {
	return u.UpdatePriceFn(uuid, price)
}

// DeletePrice calls the DeletePriceFn.
func (u *PriceAppInterface) DeletePrice(uuid string) error {
	return u.DeletePriceFn(uuid)
}

// GetPrices calls the GetPricesFn.
func (u *PriceAppInterface) GetPrices(params *repository.Parameters) ([]*entity.Price, *repository.Meta, error) {
	return u.GetPricesFn(params)
}

// GetPrice calls the GetPriceFn.
func (u *PriceAppInterface) GetPrice(uuid string) (*entity.Price, error) {
	return u.GetPriceFn(uuid)
}

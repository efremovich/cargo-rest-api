package application

import (
	"cargo-rest-api/domain/entity"
	"cargo-rest-api/domain/repository"
)

type tripApp struct {
	tr repository.TripRepository
}

// tripApp implement the TripAppInterface.
var _ TripAppInterface = &tripApp{}

// TripAppInterface is an interface.
type TripAppInterface interface {
	SaveTrip(*entity.Trip) (*entity.Trip, map[string]string, error)
	UpdateTrip(
		UUID string,
		trip *entity.Trip,
	) (*entity.Trip, map[string]string, error)
	DeleteTrip(UUID string) error
	GetTrips(p *repository.Parameters) ([]*entity.Trip, *repository.Meta, error)
	GetTrip(UUID string) (*entity.Trip, error)
}

func (t tripApp) SaveTrip(
	trip *entity.Trip,
) (*entity.Trip, map[string]string, error) {
	return t.tr.SaveTrip(trip)
}

func (t tripApp) UpdateTrip(
	UUID string,
	trip *entity.Trip,
) (*entity.Trip, map[string]string, error) {
	return t.tr.UpdateTrip(UUID, trip)
}

func (t tripApp) DeleteTrip(UUID string) error {
	return t.tr.DeleteTrip(UUID)
}

func (t tripApp) GetTrips(
	p *repository.Parameters,
) ([]*entity.Trip, *repository.Meta, error) {
	return t.tr.GetTrips(p)
}

func (t tripApp) GetTrip(UUID string) (*entity.Trip, error) {
	return t.tr.GetTrip(UUID)
}

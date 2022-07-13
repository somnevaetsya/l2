package usecase

import (
	"l2/develop/dev11/app/models"
	"l2/develop/dev11/app/repository"
)

type Usecase struct {
	repository *repository.Repository
}

func MakeUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{repository: repo}
}

func (usecase *Usecase) CreateEvent(event models.Event) (models.Event, error) {
	return usecase.repository.CreateEvent(event)
}

func (usecase *Usecase) UpdateEvent(event models.Event) (models.Event, error) {
	return usecase.repository.UpdateEvent(event)
}

func (usecase *Usecase) DeleteEvent(event models.Event) error {
	return usecase.repository.DeleteEvent(event)
}

func (usecase *Usecase) GetEvents(event models.Event) ([]models.Event, error) {
	return usecase.repository.GetEvents(event)
}

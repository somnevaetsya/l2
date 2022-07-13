package repository

import (
	"errors"
	"l2/develop/dev11/app/models"
	"sync"
)

type Repository struct {
	sync.Mutex
	values map[uint]models.Event
}

func MakeRepository() *Repository {
	data := make(map[uint]models.Event)
	return &Repository{values: data}
}

func (repo *Repository) CreateEvent(event models.Event) (models.Event, error) {
	repo.Lock()
	defer repo.Unlock()
	event.Id = uint(len(repo.values) + 1)
	repo.values[event.Id] = event
	return event, nil
}

func (repo *Repository) UpdateEvent(event models.Event) (models.Event, error) {
	repo.Lock()
	defer repo.Unlock()
	if _, ok := repo.values[event.Id]; ok {
		repo.values[event.Id] = event
		return event, nil
	} else {
		return models.Event{}, errors.New("event not found")
	}
}

func (repo *Repository) DeleteEvent(event models.Event) error {
	repo.Lock()
	defer repo.Unlock()
	if _, ok := repo.values[event.Id]; ok {
		delete(repo.values, event.Id)
		return nil
	} else {
		return errors.New("event not found")
	}
}

func (repo *Repository) GetEvents(event models.Event) ([]models.Event, error) {
	repo.Lock()
	defer repo.Unlock()
	var result []models.Event
	for _, item := range repo.values {
		if item.DateFrom.Before(event.DateTo) && item.DateTo.After(event.DateFrom) {
			result = append(result, item)
		}
	}
	return result, nil
}

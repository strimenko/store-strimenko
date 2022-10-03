package service

import (
	"github.com/strimenko/store-strimenko/models"
	"github.com/strimenko/store-strimenko/pkg/repository"
)

type HelmetService struct {
	repo repository.Helmet
}

func NewHelmetService(repo repository.Helmet) *HelmetService {
	return &HelmetService{repo: repo}
}

func (s *HelmetService) CreateHelmet(models models.Helmet) (int, error) {
	return s.repo.Create(models)
}

func (s *HelmetService) GetAll() ([]models.Helmet, error) {
	return s.repo.GetAll()
}

func (s *HelmetService) GetById(itemId int) (models.Helmet, error) {
	return s.repo.GetById(itemId)
}

func (s *HelmetService) Update(itemId int, input models.UpdateItem) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(itemId, input)
}

func (s *HelmetService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}

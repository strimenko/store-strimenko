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

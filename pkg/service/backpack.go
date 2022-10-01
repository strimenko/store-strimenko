package service

import (
	"github.com/strimenko/store-strimenko/models"
	"github.com/strimenko/store-strimenko/pkg/repository"
)

type BackpackService struct {
	repo repository.Backpack
}

func NewBackpackService(repo repository.Backpack) *BackpackService {
	return &BackpackService{repo: repo}
}

func (s *BackpackService) CreateBackpack(models models.Backpack) (int, error) {
	return s.repo.Create(models)
}

func (s *BackpackService) GetAll() ([]models.Backpack, error) {
	return s.repo.GetAll()
}

func (s *BackpackService) GetById(itemId int) (models.Backpack, error) {
	return s.repo.GetById(itemId)
}

package service

import (
	"github.com/strimenko/store-strimenko/models"
	"github.com/strimenko/store-strimenko/pkg/repository"
)

type BodyarmorService struct {
	repo repository.Bodyarmor
}

func NewBodyarmorService(repo repository.Bodyarmor) *BodyarmorService {
	return &BodyarmorService{repo: repo}
}

func (s *BodyarmorService) CreateBodyarmor(models models.Bodyarmor) (int, error) {
	return s.repo.Create(models)
}

func (s *BodyarmorService) GetAll() ([]models.Bodyarmor, error) {
	return s.repo.GetAll()
}

func (s *BodyarmorService) GetById(itemId int) (models.Bodyarmor, error) {
	return s.repo.GetById(itemId)
}

func (s *BodyarmorService) Update(itemId int, input models.UpdateItem) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(itemId, input)
}

func (s *BodyarmorService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}

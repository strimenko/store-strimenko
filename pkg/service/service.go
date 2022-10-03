package service

import (
	"github.com/strimenko/store-strimenko/models"
	"github.com/strimenko/store-strimenko/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Backpack interface {
	CreateBackpack(item models.Backpack) (int, error)
	GetAll() ([]models.Backpack, error)
	GetById(itemId int) (models.Backpack, error)
	Update(itemId int, input models.UpdateItem) error
	Delete(itemId int) error
}

type Bodyarmor interface {
	CreateBodyarmor(item models.Bodyarmor) (int, error)
	GetAll() ([]models.Bodyarmor, error)
	GetById(itemId int) (models.Bodyarmor, error)
	Update(itemId int, input models.UpdateItem) error
	Delete(itemId int) error
}

type Helmet interface {
	CreateHelmet(item models.Helmet) (int, error)
	GetAll() ([]models.Helmet, error)
	GetById(itemId int) (models.Helmet, error)
	Update(itemId int, input models.UpdateItem) error
	Delete(itemId int) error
}

type Service struct {
	Authorization
	Backpack
	Bodyarmor
	Helmet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Backpack:      NewBackpackService(repos.Backpack),
		Bodyarmor:     NewBodyarmorService(repos.Bodyarmor),
		Helmet:        NewHelmetService(repos.Helmet),
	}
}

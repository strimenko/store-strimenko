package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/strimenko/store-strimenko/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, passwrod string) (models.User, error)
}

type Backpack interface {
	Create(models models.Backpack) (int, error)
	GetAll() ([]models.Backpack, error)
	GetById(itemId int) (models.Backpack, error)
}

type Bodyarmor interface {
	Create(models models.Bodyarmor) (int, error)
	GetAll() ([]models.Bodyarmor, error)
	GetById(itemId int) (models.Bodyarmor, error)
}

type Helmet interface {
	Create(models models.Helmet) (int, error)
	GetAll() ([]models.Helmet, error)
	GetById(itemId int) (models.Helmet, error)
}

type Repository struct {
	Authorization
	Backpack
	Bodyarmor
	Helmet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Backpack:      NewBackpackPostgres(db),
		Bodyarmor:     NewBodyarmorPostgres(db),
		Helmet:        NewHelmetPostgres(db),
	}
}

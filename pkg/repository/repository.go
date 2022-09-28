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
}

type Bodyarmor interface {
}

type Helmet interface {
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
	}
}

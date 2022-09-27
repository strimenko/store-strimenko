package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}

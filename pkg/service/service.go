package service

import "github.com/strimenko/store-strimenko/pkg/repository"

type Authorization interface {
}

type Backpack interface {
}

type Bodyarmor interface {
}

type Helmet interface {
}

type Service struct {
	Authorization
	Backpack
	Bodyarmor
	Helmet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}

package service

import (
	"cmd/main.go/configs"
	"cmd/main.go/repository"
)

type Service struct {
	r      repository.Repository
	config *configs.Config
}

func NewService(r repository.Repository, config *configs.Config) *Service {
	return &Service{
		r,
		config,
	}
}

package service

import (
	"github.com/SvetlanaGrin/AvitoTest/pkg/repository"
)

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

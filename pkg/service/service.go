package service

import (
	"github.com/SvetlanaGrin/AvitoTest/pkg/repository"
)

type Service struct {
	Authorization
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Balance:       NewBalanceService(repos),
	}
}

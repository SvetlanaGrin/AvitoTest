package service

import (
	"github.com/SvetlanaGrin/AvitoTest"
	"github.com/SvetlanaGrin/AvitoTest/pkg/repository"
)

type BalanceService struct {
	repo repository.Balance
}

type Balance interface {
	GetById(userid int) (AvitoTest.UserBalance, error)
	//UpdatePlus(userid int, input AvitoTest.UserBalance) error
	//UpdateMinus(userid int, input AvitoTest.UserBalance) error
}

func NewBalanceService(repo repository.Balance) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetById(userid int) (AvitoTest.UserBalance, error) {
	return s.repo.GetById(userid)
}

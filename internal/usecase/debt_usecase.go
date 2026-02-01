package usecase

import (
	"newapp/internal/entity"
	"newapp/internal/repository"
)

type DebtUsecase struct {
	repo *repository.DebtRepository
}

func NewDebtUsecase(r *repository.DebtRepository) *DebtUsecase {
	return &DebtUsecase{
		repo: r,
	}
}

func (u *DebtUsecase) GetAllDebts() ([]entity.Debt, error) {
	return u.repo.FindAll()
}

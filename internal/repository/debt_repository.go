package repository

import (
	"database/sql"
	"newapp/internal/entity"
)

type DebtRepository struct {
	DB *sql.DB
}

func NewDebtRepository(db *sql.DB) *DebtRepository {
	return &DebtRepository{
		DB: db,
	}
}
type DebtRepository interface {
	FindAll() ([]entity.Debt, error)
}


func (r *DebtRepository) FindAll() ([]entity.Debt, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, amount, status, note, created_at, updated_at
		FROM debts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debts []entity.Debt
	for rows.Next() {
		var d entity.Debt
		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Amount,
			&d.Status,
			&d.Note,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		debts = append(debts, d)
	}

	return debts, nil
}

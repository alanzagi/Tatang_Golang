package repository

import (
	"database/sql"
	"newapp/internal/entity"
)

type debtSQLiteRepository struct {
	db *sql.DB
}

// constructor
func NewDebtSQLiteRepository(db *sql.DB) DebtRepository {
	return &debtSQLiteRepository{
		db: db,
	}
}

func (r *debtSQLiteRepository) FindAll() ([]entity.Debt, error) {
	rows, err := r.db.Query(`
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
		if err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Amount,
			&d.Status,
			&d.Note,
			&d.CreatedAt,
			&d.UpdatedAt,
		); err != nil {
			return nil, err
		}
		debts = append(debts, d)
	}

	return debts, nil
}

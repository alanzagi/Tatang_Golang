package main

import (
	"database/sql"

	_ "modernc.org/sqlite"

	"newapp/internal/handler"
	"newapp/internal/repository"
	"newapp/internal/router"
	"newapp/internal/usecase"
)

func main() {
	db, err := sql.Open("sqlite", "hutang.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewDebtSQLiteRepository(db)
usecase := usecase.NewDebtUsecase(repo)
handler := handler.NewDebtHandler(usecase)

	

	r := router.SetupRouter(debtHandler)
	r.Run(":8080")
}

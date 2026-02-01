package main

import (
	"newapp/internal/handler"
	"newapp/internal/router"
)

func main() {
	debtHandler := &handler.DebtHandler{}

	r := router.SetupRouter(debtHandler)

	r.Run(":8080")
}

package router

import (
	"newapp/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(debtHandler *handler.DebtHandler) *gin.Engine {
	r := gin.Default()

	// WAJIB: load view
	r.LoadHTMLGlob("web/templates/*.html")

	// static file (css, js)
	r.Static("/static", "./web/static")

	// route
	r.GET("/debts", debtHandler.Index)

	return r
}

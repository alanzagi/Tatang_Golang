package router

import (
	"newapp/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(debtHandler *handler.DebtHandler) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("web/templates/*.html")
	r.Static("/static", "./web/static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/debts")
	})

	r.GET("/debts", debtHandler.Index)

	return r
}

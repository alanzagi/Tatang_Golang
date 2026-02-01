package handler

import "github.com/gin-gonic/gin"

type DebtHandler struct{}

func (h *DebtHandler) Index(c *gin.Context) {
	c.HTML(200, "debts.html", gin.H{
		"title": "Daftar Hutang",
	})
}

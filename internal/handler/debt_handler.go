package handler

import (
	"net/http"

	"newapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

type DebtHandler struct {
	usecase usecase.DebtUsecase
}

func NewDebtHandler(uc usecase.DebtUsecase) *DebtHandler {
	return &DebtHandler{usecase: uc}
}

func (h *DebtHandler) Index(c *gin.Context) {
	debts, err := h.usecase.GetAllDebts()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "debts.html", gin.H{
		"title": "Daftar Hutang",
		"debts": debts,
	})
}

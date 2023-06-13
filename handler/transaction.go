package handler

import (
	"letsfunding/helper"
	"letsfunding/transaction"
	"letsfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	transactionHandler struct {
		service transaction.Service
	}
)

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign Transaction", http.StatusOK, "Success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}

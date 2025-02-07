package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) GetTransactionsByCampaignID(c *gin.Context) {
	var input transaction.GetTransactionByCampaignIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormat := transaction.FormatCampaignTransactions(transactions)

	response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transactionsFormat)
	c.JSON(http.StatusOK, response)
	return
}

func (h *TransactionHandler) GetTransactionsByUserID(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactionsFormat := transaction.FormatUserTransactions(transactions)

	response := helper.APIResponse("User's transactions", http.StatusOK, "success", transactionsFormat)
	c.JSON(http.StatusOK, response)
	return
}

// input jumlah funding dari user
// handler tangkap input lalu di-mapping ke input struct
// panggil service buat transaksi, manggil sistem midtrans
// panggil repository, create new transaction data

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/service"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type transactionRequest struct {
	AccountID       uint            `json:"account_id"`
	OperationTypeID int             `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
}

type transactionResponse struct {
	TransactionID   uint            `json:"transaction_id"`
	AccountID       uint            `json:"account_id"`
	OperationTypeID int             `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
}

type TransactionHandler struct {
	TransactionService *service.Service
	Logger             *zap.Logger
}

func NewTransactionHandler(service *service.Service, logger *zap.Logger) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: service,
		Logger:             logger,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	context := ctx.Request.Context()

	var req transactionRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var input model.Transaction

	if err := copier.Copy(&input, &req); err != nil {
		h.Logger.Error("Failed to copy request data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Logger.Info("Start creating transaction")
	transaction, err := h.TransactionService.CreateTransaction(context, &input)

	if err != nil {
		h.Logger.Error("Failed to create transaction", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	var response transactionResponse

	if err := copier.Copy(&response, transaction); err != nil {
		h.Logger.Error("Failed to copy transaction data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Logger.Info("Transaction created successfully")
	ctx.JSON(http.StatusCreated, gin.H{"transaction": response})

}

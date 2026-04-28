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

// transactionRequest represents the expected JSON structure for creating a transaction.
// Contains an AccountID, OperationTypeID, and Amount field which are required for transaction creation.
type TransactionRequest struct {
	AccountID       uint            `json:"account_id"`
	OperationTypeID int             `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
}

// transactionResponse represents the JSON structure for transaction responses.
// Contains a TransactionID, AccountID, OperationTypeID, and Amount field to be
// returned in the HTTP response when creating a transaction.
type TransactionResponse struct {
	TransactionID   uint            `json:"transaction_id"`
	AccountID       uint            `json:"account_id"`
	OperationTypeID int             `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
}

// TransactionHandler handles HTTP requests related to transactions.
// Contains a reference to the TransactionService for business logic and a Logger for logging.
type TransactionHandler struct {
	TransactionService *service.Service
	Logger             *zap.Logger
}

// NewTransactionHandler creates a new instance of TransactionHandler
// with the provided service and logger.
func NewTransactionHandler(service *service.Service, logger *zap.Logger) *TransactionHandler {
	return &TransactionHandler{
		TransactionService: service,
		Logger:             logger,
	}
}

// CreateTransaction godoc
//
// @Summary Create a new transaction
// @Description Create a new transaction with the provided account ID, operation type ID, and amount
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body TransactionRequest true "Transaction creation payload"
// @Success 201 {object} TransactionResponse "Transaction created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(ctx *gin.Context) {
	context := ctx.Request.Context()

	var req TransactionRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		h.Logger.Error("failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var input model.Transaction

	if err := copier.Copy(&input, &req); err != nil {
		h.Logger.Error("failed to copy request data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	h.Logger.Info("Start creating transaction")
	transaction, err := h.TransactionService.CreateTransaction(context, &input)

	if err != nil {
		h.Logger.Error("failed to create transaction", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	var response TransactionResponse

	if err := copier.Copy(&response, transaction); err != nil {
		h.Logger.Error("failed to copy transaction data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	h.Logger.Info("transaction created successfully")
	ctx.JSON(http.StatusCreated, response)

}

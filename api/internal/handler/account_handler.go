package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/momoyo-droid/pismo/api/internal/model"
	"github.com/momoyo-droid/pismo/api/internal/service"
	"go.uber.org/zap"
)

// AccountHandler handles HTTP requests related to accounts.
// Contains a reference to the AccountService for business logic and a Logger for logging.
type AccountHandler struct {
	AccountService *service.Service
	Logger         *zap.Logger
}

// NewAccountHandler creates a new instance of AccountHandler with the provided service and logger.
func NewAccountHandler(service *service.Service, logger *zap.Logger) *AccountHandler {
	return &AccountHandler{
		AccountService: service,
		Logger:         logger,
	}
}

// accountRequest represents the expected JSON structure for creating an account.
// Contains a DocumentNumber field which is required for account creation.
type AccountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

// accountResponse represents the JSON structure for account responses.
// COntains an ID and DocumentNumber field to be returned in the HTTP response when fetching or creating an account.
type AccountResponse struct {
	ID             uint   `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

// CreateAccount godoc
//
// @Summary Create a new account
// @Description Create a new account with the provided document number
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body AccountRequest true "Account creation payload"
// @Success 201 {object} AccountResponse "Account created successfully"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(ctx *gin.Context) {
	context := ctx.Request.Context()

	var req AccountRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		h.Logger.Error("failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var input model.Account
	if err := copier.Copy(&input, &req); err != nil {
		h.Logger.Error("failed to copy request data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	h.Logger.Info("start creating account")
	account, err := h.AccountService.CreateAccount(context, &input)
	if err != nil {
		h.Logger.Error("failed to create account", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	var response AccountResponse

	if err := copier.Copy(&response, account); err != nil {
		h.Logger.Error("failed to copy account data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	h.Logger.Info("account created successfully")
	ctx.JSON(http.StatusCreated, response)
}

// GetAccountByID godoc
//
// @Summary Get account by ID
// @Description Retrieve account details by its ID
// @Tags accounts
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {object} AccountResponse "Account fetched successfully"
// @Failure 400 {object} map[string]string "Invalid account ID"
// @Failure 404 {object} map[string]string "Account not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /accounts/{id} [get]
func (h *AccountHandler) GetAccountByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	accountID := ctx.Param("id")

	id, err := strconv.ParseUint(accountID, 10, 64)
	if err != nil {
		h.Logger.Error("invalid account ID", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid account ID"})
		return
	}

	h.Logger.Info("start fetching account")
	account, err := h.AccountService.GetAccountByID(context, id)

	if err != nil {
		h.Logger.Error("failed to fetch account", zap.Error(err))
		ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	var response AccountResponse

	if err := copier.Copy(&response, account); err != nil {
		h.Logger.Error("failed to copy account data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	h.Logger.Info("account fetched successfully")
	ctx.JSON(http.StatusOK, response)
}

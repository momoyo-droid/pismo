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

type AccountHandler struct {
	AccountService *service.Service
	Logger         *zap.Logger
}

func NewAccountHandler(service *service.Service, logger *zap.Logger) *AccountHandler {
	return &AccountHandler{
		AccountService: service,
		Logger:         logger,
	}
}

type accountRequest struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}

type accountResponse struct {
	ID             uint   `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (h *AccountHandler) CreateAccount(ctx *gin.Context) {
	context := ctx.Request.Context()

	var req accountRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var input model.Account
	if err := copier.Copy(&input, &req); err != nil {
		h.Logger.Error("Failed to copy request data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Logger.Info("Start creating account")
	account, err := h.AccountService.CreateAccount(context, &input)
	var response accountResponse

	if err := copier.Copy(&response, account); err != nil {
		h.Logger.Error("Failed to copy account data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err != nil {
		h.Logger.Error("Failed to create account", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	h.Logger.Info("Account created successfully")
	ctx.JSON(http.StatusCreated, gin.H{"account": response})
}

func (h *AccountHandler) GetAccountByID(ctx *gin.Context) {
	context := ctx.Request.Context()

	accountID := ctx.Param("id")

	id, err := strconv.ParseUint(accountID, 10, 64)
	if err != nil {
		h.Logger.Error("Invalid account ID", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	h.Logger.Info("Start fetching account", zap.String("account_id", accountID))
	account, err := h.AccountService.GetAccountByID(context, id)

	if err != nil {
		h.Logger.Error("Failed to fetch account", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch account"})
		return
	}

	var response accountResponse

	if err := copier.Copy(&response, account); err != nil {
		h.Logger.Error("Failed to copy account data", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	h.Logger.Info("Account fetched successfully", zap.String("account_id", accountID))
	ctx.JSON(http.StatusOK, gin.H{"account": response})
}

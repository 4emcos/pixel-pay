package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pixel-pay/database"
	"pixel-pay/internal/services"
	"pixel-pay/internal/types"
)

func PostTransaction(context *gin.Context, db database.Pgx) {
	transaction := &types.TransactionRequest{}

	if err := context.ShouldBindJSON(transaction); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := services.NewTransaction(db, *transaction)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "transaction success"})
	return
}

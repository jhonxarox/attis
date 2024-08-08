package routes

import (
	"backend-assignment/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeTransactionRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/transactions", func(c *gin.Context) {
		var input struct {
			AccountID uint    `json:"accountId"`
			Amount    float64 `json:"amount"`
			Type      string  `json:"type"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var account models.Account
		if err := db.Where("id = ?", input.AccountID).First(&account).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
			return
		}

		transaction := models.Transaction{AccountID: input.AccountID, Amount: input.Amount, Type: input.Type}

		// Process the transaction asynchronously
		go processTransaction(transaction, db)

		c.JSON(http.StatusAccepted, gin.H{"message": "Transaction is being processed", "transaction": transaction})
	})

	r.GET("/transactions/:accountId", func(c *gin.Context) {
		accountId := c.Param("accountId")
		var transactions []models.Transaction

		if err := db.Where("account_id = ?", accountId).Find(&transactions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	})
}

func processTransaction(transaction models.Transaction, db *gorm.DB) {
	// Simulate long-running process
	time.Sleep(30 * time.Second)

	var account models.Account
	if err := db.Where("id = ?", transaction.AccountID).First(&account).Error; err != nil {
		// Handle error
		return
	}

	if transaction.Type == "send" {
		account.Balance -= transaction.Amount
	} else if transaction.Type == "withdraw" {
		account.Balance += transaction.Amount
	}

	db.Save(&account)
	db.Create(&transaction)
}

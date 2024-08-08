package routes

import (
	"backend-assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeAccountRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/accounts", func(c *gin.Context) {
		var input struct {
			UserID  uint    `json:"userId"`
			Type    string  `json:"type"`
			Balance float64 `json:"balance"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		account := models.Account{UserID: input.UserID, Type: input.Type, Balance: input.Balance}

		if err := db.Create(&account).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"account": account})
	})

	r.GET("/accounts/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		var accounts []models.Account

		if err := db.Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"accounts": accounts})
	})
}

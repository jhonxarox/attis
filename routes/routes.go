package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	InitializeAuthRoutes(r, db)
	InitializeAccountRoutes(r, db)
	InitializeTransactionRoutes(r, db)
}

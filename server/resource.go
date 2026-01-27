package server

import (
	"TODO-MIS/common/middware"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGinEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middware.AuthMiddleware())
	return r
}

func NewLogger() (*zap.Logger, error) {
	if os.Getenv("APP_ENV") == "prod" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}

func NewDB() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		panic("MYSQL_DSN is empty")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

package server

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGinEngine() *gin.Engine {
	return gin.New()
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

	var db *gorm.DB
	var err error

	// Retry connecting to the database a few times to wait for MySQL to be ready
	const maxRetries = 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic(fmt.Errorf("connect fail after retries: %v, DSN: %s", err, dsn))
	}

	return db
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"jaketrent.com/roboretro/messages"
	"log"
	"os"
)

func hasDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Db unable to connect", err)
	}
	defer db.Close()

	db.AutoMigrate(&messages.Message{})

	router := gin.Default()
	router.Use(hasDatabase(db))
	messages.Mount(router)
	router.Run()
}

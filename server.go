package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"io/ioutil"
	"jaketrent.com/roboretro/messages"
	"log"
	"net/http"
	"os"
)

func hasDatabase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func RequestLogger(c *gin.Context) {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	rdr1 := ioutil.NopCloser((bytes.NewBuffer(buf)))
	rdr2 := ioutil.NopCloser((bytes.NewBuffer(buf)))

	fmt.Println(func(reader io.Reader) string {
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		s := buf.String()
		return s
	}(rdr1))

	c.Request.Body = rdr2
	c.Next()
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
	router.Use(gin.Logger())
	router.Use(RequestLogger)
	router.Use(gin.Recovery())
	router.Use(hasDatabase(db))
	messages.Mount(router)
	router.GET("/", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/ui") })
	router.Static("/ui", "./client/build")
	router.Run()
}

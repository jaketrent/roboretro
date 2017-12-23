package messages

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"time"
)

type ok struct {
	Data []*Message `json:"data"`
}
type clienterr struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}
type bad struct {
	Errors []*clienterr `json:"errors"`
}

func createMessage(userId string, userName string, text string) *Message {
	email, err := getUserEmail(userId)
	if err != nil {
		email = "Unknown"
	}
	return &Message{
		Name:  userName,
		Email: email,
		Date:  time.Now().Unix(),
		Text:  text,
	}
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	slackToken := os.Getenv("SLACK_VERIFICATION_TOKEN")

	var msg *Message
	var err error
	token := c.PostForm("token")
	userId := c.PostForm("user_id")
	userName := c.PostForm("user_name")
	text := c.PostForm("text")

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		fmt.Println("message create req error", err)
		return
	}

	if token != slackToken {
		c.String(http.StatusUnauthorized, "Invalid credentials")
		fmt.Println("invalid slack token", token)
		return
	}

	msg = createMessage(userId, userName, text)
	msg, err = insert(db, msg)

	if err != nil {
		c.String(http.StatusInternalServerError, "Create error")
		fmt.Println("message create db error", err)
		return
	}

	c.String(http.StatusOK, "Thanks %s, posted!", userName)
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	messages, err := findAll(db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []*clienterr{{Title: "List error", Status: http.StatusInternalServerError}},
		})
	}

	c.JSON(http.StatusOK, ok{Data: messages})
}

func Mount(router *gin.Engine) {
	router.POST("/api/v1/messages", create)
	router.GET("/api/v1/messages", list)
}

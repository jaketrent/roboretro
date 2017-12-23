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

func createMessage(userName string, text string) *Message {
	return &Message{
		Name:  userName,
		Email: "TODO",
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
	userName := c.PostForm("user_name")
	text := c.PostForm("text")

	if err != nil {
		c.JSON(http.StatusBadRequest, bad{
			Errors: []*clienterr{{"Bad request", http.StatusBadRequest}},
		})
		fmt.Println("message create req error", err)
		return
	}

	if token != slackToken {
		c.JSON(http.StatusUnauthorized, bad{
			Errors: []*clienterr{{"Invalid credentials", http.StatusUnauthorized}},
		})
		fmt.Println("invalid slack token", token)
		return
	}

	msg = createMessage(userName, text)
	msg, err = insert(db, msg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, bad{
			Errors: []*clienterr{{Title: "Create error", Status: http.StatusInternalServerError}},
		})
		fmt.Println("message create db error", err)
		return
	}

	c.JSON(http.StatusOK, struct {
		Data *Message `json:"data"`
	}{msg})
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

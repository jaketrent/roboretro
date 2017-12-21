package messages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type message struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
	Text  string `json:"text"`
}

type ok struct {
	Data []*message `json:"data"`
}

func list(c *gin.Context) {
	messages := []*message{{"jake", "trent.jake@gmail.com", "23 Dec 2017", "wow"}}
	c.JSON(http.StatusOK, ok{Data: messages})
}

func Mount(router *gin.Engine) {
	router.GET("/api/v1/messages", list)
}

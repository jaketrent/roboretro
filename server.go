package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jaketrent.com/roboretro/messages"
)

func main() {
	fmt.Println("wow")

	router := gin.Default()
	messages.Mount(router)
	router.Run()
}

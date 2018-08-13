package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someGet",middleware1,middleware2,handler)
	router.Run(":8888")
}

func handler(c *gin.Context)  {
	log.Printf("exec handler")
	c.JSON(http.StatusOK,gin.H{
		"status":"get",
	})
}

func middleware1(c *gin.Context) {
	log.Printf("exec middleware1")
	c.Next()
}
func middleware2(c *gin.Context) {
	log.Printf("arrive at middleware2")
	c.Next()
	log.Printf("exec middleware2")
}
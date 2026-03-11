package main

import (
	"net/http"
  "github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
  router := gin.Default()
  
  router.GET("/health", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status": "healthy",
	})
  })

  return router
}


func StartServer(router *gin.Engine) {
	
}
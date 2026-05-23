package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("DevDeploy Starting...")

	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"pong",
		})
	}
	r.Run(":8000")
}


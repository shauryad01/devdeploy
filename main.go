package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shauryad01/devdeploy/db"
)

func main() {
	godotenv.Load()
	fmt.Println("DevDeploy Starting...")

	conn, err := db.Connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	log.Println("database connected:", conn.Stats())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8000")
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/shauryad01/devdeploy/db"
)

type CreateDeploymentRequest struct {
	RepoURL string `json:"repo_url"`
}

type Deployment struct {
	ID        int       `db:"id"`
	RepoURL   string    `db:"repo_url"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	godotenv.Load()
	fmt.Println("DevDeploy Starting...")

	conn, err := db.Connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	log.Println("database connected:", conn.Stats())

	jobs := make(chan int, 100)
	go worker(jobs, conn)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/deployments", func(c *gin.Context) {
		var req CreateDeploymentRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request",
			})
			return
		}

		if req.RepoURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "repo_url is required",
			})
			return
		}

		var d Deployment

		query := "INSERT INTO deployments (repo_url) VALUES ($1) RETURNING *"
		row := conn.QueryRowx(query, req.RepoURL)
		err := row.StructScan(&d)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create Deployment",
			})
			return
		}
		jobs <- d.ID
		c.JSON(http.StatusCreated, d)
	})

	r.GET("/deployments/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid or empty ID"})
			return
		}

		var d Deployment

		query := "SELECT * FROM deployments WHERE id = $1"
		err = conn.Get(&d, query, id)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "deployment not found"})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch deployment"})
			return
		}
		c.JSON(http.StatusOK, d)

	})
	r.Run(":8000")
}

func worker(jobs chan int, db *sqlx.DB) {
	for id := range jobs {
		fmt.Printf("processing %d", id)
	}
}

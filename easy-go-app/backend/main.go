package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"izi-go.com/database"
	"izi-go.com/task"
)

func main() {
	r := gin.Default()

	// Init database
	db, err := database.New(&database.Config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()
	// Run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Init resources
	task.Init(db, r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

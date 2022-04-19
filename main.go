package main

import (
	"log"
	"path"
	"path/filepath"

	"github.com/codypotter/lofi-blog/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./dist/lofi-blog/index.html")
		} else {
			c.File("./dist/lofi-blog/" + path.Join(dir, file))
		}
	})

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	db.Connect()

	if err := r.Run(); err != nil {
		log.Fatal("Server failed to start")
	}
}

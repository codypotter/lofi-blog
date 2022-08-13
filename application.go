package main

import (
	"log"
	"path"
	"path/filepath"

	"github.com/codypotter/lofi-blog/controller"
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
		posts := api.Group("/posts")
		{
			posts.GET("/", controller.GetAllPosts)
			posts.PUT("/refresh", controller.RefreshPosts)
			posts.GET("/featured", controller.GetFeaturedPost)
			posts.GET("/:id", controller.GetPostById)
			posts.PUT("/", controller.UpvotePost)
		}
	}

	db.Connect()

	if err := r.Run(":5000"); err != nil {
		log.Fatal("Server failed to start")
	}
}

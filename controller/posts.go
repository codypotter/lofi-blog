package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/codypotter/lofi-blog/db"
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
		log.Printf("get all posts failed to parse querystring for page: %v\n", err)
	}
	posts, err := db.GetAllPosts(c, page)
	if err != nil {
		log.Printf("error getting all posts: %v\n", err)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, struct {
		Page  int       `json:"page"`
		More  bool      `json:"more"`
		Posts []db.Post `json:"posts"`
	}{
		page,
		len(posts) >= 10,
		posts,
	})
}

func GetFeaturedPost(c *gin.Context) {
	featuredPost, err := db.GetMostRecentPost(c)
	if err != nil {
		log.Printf("error getting featured post: %v\n", err)
		if errors.Is(err, db.ErrNotFound) {
			c.AbortWithError(404, err)
			return
		}
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, featuredPost)
}

func GetPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("received non numeric id in get post by id")
		c.AbortWithStatus(400)
		return
	}
	post, err := db.GetPostById(c, id)
	if err != nil {
		log.Printf("error getting post by id: %v\n", err)
		if errors.Is(err, db.ErrNotFound) {
			c.AbortWithError(404, err)
			return
		}
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, post)
}

func DropAndReloadPosts(c *gin.Context) {
	err := db.DropPosts(c)
	if err != nil {
		log.Printf("error dropping posts: %v\n", err)
		c.AbortWithError(500, err)
		return
	}

	err = db.ReloadPosts(c)
	if err != nil {
		log.Printf("error reloading posts: %v\n", err)
		c.AbortWithError(500, err)
		return
	}
	c.Status(http.StatusNoContent)
}

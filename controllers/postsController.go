package controllers

import (
	"github.com/gin-gonic/gin"
	"www.github.com/teaampee/jsonCRUD/inits"
	"www.github.com/teaampee/jsonCRUD/models"
)

func PostsCreate(c *gin.Context) {
	//getting data of req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := inits.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	inits.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	//getting data off url
	id := c.Param("id")

	// get post
	var post models.Post
	inits.DB.First(&post, id)

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//getting id off url
	id := c.Param("id")
	// get data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// find post to be updated.
	var post models.Post
	inits.DB.First(&post, id)
	// update it

	result := inits.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	if result.Error != nil {
		c.Status(400)
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

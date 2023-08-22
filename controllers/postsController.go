package controllers

import (
	"errors"

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
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "false data entered",
		})
		return
	}

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := inits.DB.Create(&post) // pass pointer of data to Create
	if result.Error != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "couldn't create posts",
		})

		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	result := inits.DB.Find(&posts)
	if result.Error != nil {
		c.Status(500)
		c.JSON(500, gin.H{
			"message": "couldn't find posts",
			"error":   result.Error,
		})
		errors.New("no posts on server?")

	}
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	//getting data off url
	id := c.Param("id")

	// get post
	var post models.Post
	result := inits.DB.First(&post, id)
	if result.Error != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "couldn't find post",
			"error":   result.Error,
		})
		return
	}

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
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "false data entered",
		})
		return
	}

	// find post to be updated.
	var post models.Post
	result := inits.DB.First(&post, id)
	if result.Error != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "couldn't find post",
			"error":   result.Error,
		})
		return
	}
	// update it

	result = inits.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	if result.Error != nil {
		c.Status(400)
		c.JSON(400, gin.H{
			"message": "couldn't update the post",
			"error":   result.Error,
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gjustin40/CRUD/initalizers"
	"github.com/gjustin40/CRUD/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initalizers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initalizers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"post": posts,
	})
	// Response with them
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initalizers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
	// Response with them
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Get the posts
	var post models.Post
	initalizers.DB.First(&post, id)

	initalizers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the posts
	initalizers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}

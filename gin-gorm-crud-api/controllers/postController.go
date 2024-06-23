package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/initializers"
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/models"
)

// Create post
func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Find many posts
func PostsFind(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// Find one post
func PostFind(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Update post
func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Body:  body.Body,
		Title: body.Title,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Delete post
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}

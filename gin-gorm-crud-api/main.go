package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/controllers"
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/post", controllers.PostCreate)

	r.GET("/findPosts", controllers.PostsFind)
	r.GET("/findPost/:id", controllers.PostFind)

	r.PUT("/updatePost/:id", controllers.PostUpdate)

	r.DELETE("/deletePost/:id", controllers.DeletePost)

	r.Run()
}

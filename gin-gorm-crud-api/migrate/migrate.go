package main

import (
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/initializers"
	"github.com/thedidev/golang-learn/gin-gorm-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

package main

import (
	"github.com/debg48/gin_crud/initializers"
	"github.com/debg48/gin_crud/models"
)

func init() {
	initializers.LoadEnvVaraibles()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

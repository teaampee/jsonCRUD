package main

import (
	"www.github.com/teaampee/jsonCRUD/inits"
	"www.github.com/teaampee/jsonCRUD/models"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConnetToDB()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}

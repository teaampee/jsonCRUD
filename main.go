package main

import (
	"github.com/gin-gonic/gin"
	"www.github.com/teaampee/jsonCRUD/controllers"
	"www.github.com/teaampee/jsonCRUD/inits"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConnetToDB()

}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.POST("/posts/:id", controllers.PostsUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.Run() // listen and serve on 0.0.0.0:8080
}

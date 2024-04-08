package routes

import (
	"github.com/gin-gonic/gin"
	"goapi.railway.app/api/controllers"
)

func RouterSetup() *gin.Engine{
	route := gin.Default() 

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(
			200, gin.H{
				"message": "welcome to api",
		})
	})

	route.POST("/book", controllers.CreateBook)
	route.GET("/book", controllers.GetAll)
	route.GET("/book/:id", controllers.GetOne)
	route.PUT("/book/:id", controllers.UpdateBook)
	route.PATCH("/book/:id", controllers.PatchBook)
	route.DELETE("/book/:id", controllers.DeleteBook)


	return route
}
package main

import (
	"github.com/gin-gonic/gin"
	"goapi.railway.app/api/controllers"
)

func (app *application) routes() *gin.Engine {
	// router := httprouter.New()

	// Define the available routes
	// router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// return router
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

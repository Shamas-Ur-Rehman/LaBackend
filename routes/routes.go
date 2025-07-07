package routes

import (
	"Laorgaincs/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	group := r.Group("/products")
	{
		group.POST("/", controllers.CreateProduct)
		group.GET("/", controllers.GetAllProducts)
		group.GET("/:id", controllers.GetProduct)
		group.PUT("/:id", controllers.UpdateProduct)
		group.DELETE("/:id", controllers.DeleteProduct)
	}
}

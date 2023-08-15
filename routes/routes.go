package routes

import (
	"go-freecodecamp/go-bookstore/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.POST("/signup", controllers.SignUp())
	routes.POST("/login", controllers.Login())
	routes.POST("/admin/add-product", controllers.AddProduct())
	routes.GET("/users/product-view", controllers.ViewProduct())
	routes.GET("/users/search", controllers.SearchProduct())
}
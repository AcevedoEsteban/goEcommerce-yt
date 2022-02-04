package routes

import (
	"github.com/AcevedoEsteban/goEcommerce-yt/controllers"
	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.Post("/users/signup", controllers.SignUp())
	incomingRoutes.Post("/users/login", controllers.Login())
	incomingRoutes.Post("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}

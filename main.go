package main

import (
	"log"
	"os"

	"github.com/AcevedoEsteban/goEcommerce-yt/controllers"
	"github.com/AcevedoEsteban/goEcommerce-yt/database"
	"github.com/AcevedoEsteban/goEcommerce-yt/middleware"
	"github.com/AcevedoEsteban/goEcommerce-yt/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.Client, "Products"), database.UserData(database.Client, "Users")

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouters(routes)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}

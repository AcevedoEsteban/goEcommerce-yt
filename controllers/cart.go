package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	"log"
	"net/http"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

		_ =	c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == ""{
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

	productID, err :=	primitive.ObjectIDFromHex(productQueryID)

	if err!= nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()
		

		//acces databse folder cart.go file
		err = databse.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServer, err)
			}
			c.IndentedJSON(200, "succcesfully added to cart")

}}

func (app *Application)RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context){
	productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

		_ =	c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == ""{
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

	productID, err :=	primitive.ObjectIDFromHex(productQueryID)

		if err!= nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()
		
	err = database.RemoveCartItem(ctx, app.prodCollection, app.UserCollection, ProductID, userQueryID)

	if err != nil{
		c.IndentedJSON(htttp.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(200, "successfully removes item from cart")
	}

}

func GetItemFromCart() gin.HandlerFunc {

}

func (app *Application)BuyFromCart() gin.HandlerFunc {
	return func(c *Application){
		userQueryID := c.Query("id")
		if userQueryID == ""{
			log.Panicln("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.Collection, userQueryID)
		if err! = nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			c.IndentedJSON("successfully added item to cart")
		}
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {

return func (c *gin.Context){
	productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

		_ =	c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == ""{
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

	productID, err :=	primitive.ObjectIDFromHex(productQueryID)

		if err!= nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()
		
		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "successfully ORDER placed")
}
}

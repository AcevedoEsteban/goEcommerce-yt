package controllers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"errors"
	"log"
	"net/http"
	"context"
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

	protductID, err :=	primitive.ObjectIDFromHex(productQueryID)

	if err!= nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
		}

		var ctx, cancel = context.WithTimeOut(context.Background())
}

func RemoveItem() gin.HandlerFunc {

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}

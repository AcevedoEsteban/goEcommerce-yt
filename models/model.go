package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_name      *string            `json:"first_name" validate:"required,min=2,max=30`
	Last_name       *string            `json:"last_name" validate:"required,min=2,max=30`
	Password        *string            `json:"password"  validate:"required,min=6"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_token   *string            `json:"refresh_token"`
	Create_At       time.Time          `json:"created_at"`
	Update_At       time.Time          `json:"updated_at"`
	User_ID         *string            `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart`
	Address_details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"address" bson:"address"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint16            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"streat_name" bson:"streat_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list" bson:"order_list"`
	Order_Status   time.Time          `json:"order_at" bson:"order_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}

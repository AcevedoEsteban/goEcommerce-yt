package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct { 

	ID					primitive.ObjectID
	First_name 			*string
	Last_name			*string
	Email				*string
	Phone 				*string
	Token				*string
	Refresh_token		*string
	Create_At			*string
	Update_At			
	User_ID
	UserCart
	Address_details
	Order_Status

}

type Product struct{
Product_ID
Product_Name
Price
Rating
Image
}

type ProductUser struct{
Product_ID
Product_Name
Price
Rating
Image
}

type Address struct{
	Address_ID
	House 
	Street
	City
	Pincode
}

type Order struct{
	Order_ID
	Order_Cart
	Order_Status
	Price
	Discount
	Payment_Method
}

type Payment struct{
	Digital bool
	COD		bool
}
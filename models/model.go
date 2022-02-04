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
	Create_At			time.Time
	Update_At			time.Time
	User_ID				*string
	UserCart			[]ProductUser
	Address_details		[]Address
	Order_Status		[]Order

}

type Product struct{
Product_ID			primitive.ObjectID
Product_Name		*string
Price				*uint16
Rating				*uint8
Image				*string
}

type ProductUser struct{
Product_ID			primitive.ObjectID
Product_Name		*string
Price				int
Rating				*uint
Image				*string
}

type Address struct{
	Address_ID			primitive.ObjectID
	House 				*string
	Street				*string
	City				*string
	Pincode				*string
}

type Order struct{
	Order_ID			primitive.ObjectID
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
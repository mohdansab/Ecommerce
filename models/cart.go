package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	ProductId   primitive.ObjectID `bson:"productId" json:"productId"`
	Quantity    int                `bson:"qty" json:"qty"`
	UserId      primitive.ObjectID `bson:"userid" json:"userid"`
	ProductName string             `bson:"productName" json:"productName"`
}

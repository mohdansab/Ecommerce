package models

type Product struct {
	Name  string `bson:"name" json:"name"`
	Price int64  `bson:"price" json:"price"`
}

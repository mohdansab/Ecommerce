package models

type User struct {
	FirstName string `bson:"FirstName" json:"FirstName"`
	LastName  string `bson:"LastName" json:"LastName"`
	PhoneNo   int    `bson:"number" json:"number"`
	Email     string `bson:"email" json:"email"`
	Password  int    `bson:"password" json:"password"`
}

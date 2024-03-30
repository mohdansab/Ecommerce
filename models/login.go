package models

type Login struct {
	Email    string `bson:"email" json:"email"`
	Password int    `bson:"password" json:"password"`
}

package models

type Product struct {
	Name        string `bson:"name" json:"name"`
	Price       int64  `bson:"price" json:"price"`
	Description string `bson:"description" json:"description"`
	Quantity    int    `bson:"qty" json:"qty"`
	Brand       string `bson:"brand" json:"brand"`
	Image1      string `bson:"image1" json:"image1"`
	Image2      string `bson:"image2" json:"image2"`
	Image3      string `bson:"image3" json:"image3"`
}

package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohdansab/db"
	"github.com/mohdansab/models"
	"gopkg.in/mgo.v2/bson"
)

func AddCart(c *gin.Context) {
	var cart models.Cart
	err := c.BindJSON(&cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	var product models.Product
	err = db.Product.FindOne(c, bson.M{"name": cart.ProductName}).Decode(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	var cart1 models.Cart
	err = db.Cart.FindOne(c, bson.M{"productName": cart.ProductName}).Decode(&cart1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "product not updated"})
		return
	}

	if cart1.ProductName == cart.ProductName {
		filter := bson.M{"productName": cart.ProductName}
		update := bson.M{"$set": bson.M{"quantity": cart.Quantity + cart1.Quantity}}
		_, err = db.Cart.UpdateOne(c, filter, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "product not updated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "product updated"})
		return

	}
	if product.Quantity <= cart.Quantity {
		c.JSON(http.StatusOK, gin.H{"message": " number of product available in the data base"})
		return
	}

	_, err = db.Cart.InsertOne(c, cart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})

}

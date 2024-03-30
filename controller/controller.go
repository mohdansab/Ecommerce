package controller

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/mohdansab/db"
	"github.com/mohdansab/middleware"
	"github.com/mohdansab/models"
	"gopkg.in/mgo.v2/bson"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"posts": "hello world"})

}

func Create(c *gin.Context) {
	var Product models.Product
	err := c.BindJSON(&Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	var Product2 models.Product
	err = db.Product.FindOne(c, bson.M{"name": Product.Name}).Decode(&Product2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	if Product2.Name == Product.Name {
		c.JSON(http.StatusBadRequest, gin.H{"message": "product already existed"})
		return
	}
	_, err = db.Product.InsertOne(c, Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"Product": Product})

}

func Update(c *gin.Context) {
	var Product models.Product
	err := c.BindJSON(&Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	filter := bson.M{"name": Product.Name}
	update := bson.M{"$set": bson.M{"price": Product.Price}}

	var Product3 models.Product
	err = db.Product.FindOne(c, bson.M{"price": Product.Price}).Decode(&Product3)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	if Product3.Price == Product.Price {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Given product price is already existed .try another price"})
		return
	}

	_, err = db.Product.UpdateOne(c, filter, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfull"})

}

func Delete(c *gin.Context) {
	var Product models.Product
	err := c.BindJSON(&Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	filter := bson.M{"name": Product.Name}

	var Product4 models.Product
	err = db.Product.FindOne(c, bson.M{"name": Product.Name}).Decode(&Product4)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not found"})
		return
	}

	_, err = db.Product.DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})

}

func ViewAll(c *gin.Context) {
	var Product []models.Product
	cur, err := db.Product.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	err = cur.All(c, &Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product": Product})

}

func SignUp(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	if user.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "firstName required"})
		return
	}
	if user.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lastName required"})
		return
	}

	if user.PhoneNo == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid phone number"})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email required"})
		return
	}
	if !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid email"})
		return
	}

	if user.Password == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password required"})
		return
	}

	_, err = db.User.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "signUp successfull"})
}

// email validation
func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	return re.MatchString(email)
}

func Login(c *gin.Context) {
	var login models.Login
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}
	var user2 models.User
	err = db.User.FindOne(c, bson.M{"password": login.Password}).Decode(&user2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect password"})
		return
	}
	if login.Password != user2.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect"})
		return
	}
	tokenString, err := middleware.GenerateJWT(login.Email, 10)
	c.SetCookie("UserAuth", tokenString, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": tokenString})

}

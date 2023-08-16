package controllers

import (
	"context"
	"fmt"
	"go-ecommerce/database"
	"go-ecommerce/models"
	tokenGenerator "go-ecommerce/tokens"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Validate                            = validator.New()
	UserCollection    *mongo.Collection = database.UserData(database.Client, "Users")
	ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(context, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is registered"})
			return
		}

		count, err = UserCollection.CountDocuments(context, bson.M{"phone": user.Phone})
		defer cancel()

		if err != nil {
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "phone number is registered"})
			return
		}

		password := HashPassword(*user.Password)
		now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.Password = &password
		user.Created_At = now
		user.Updated_At = now
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		token, refreshToken, _ := tokenGenerator.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)

		user.Token = &token
		user.Refresh_Token = &refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, err = UserCollection.InsertOne(context, user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
			return
		}
		defer cancel()

		ctx.JSON(http.StatusCreated, "successfully signed in!")
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		var foundUser models.User
		err := UserCollection.FindOne(context, bson.M{"email": user.Email}).Decode(foundUser)
		defer cancel()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "login credentials incorrect"})
			return
		}

		isPasswordValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()

		if !isPasswordValid {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		token, refreshToken, _ := tokenGenerator.TokenGenerator(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, foundUser.User_ID)
		defer cancel()

		tokenGenerator.UpdateAllTokens(token, refreshToken, foundUser.User_ID)

		ctx.JSON(http.StatusFound, foundUser)
	}
}

func AddProduct() gin.HandlerFunc {

}

func ViewProduct() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

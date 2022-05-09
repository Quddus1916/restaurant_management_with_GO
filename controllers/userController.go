package controller

import(
	"github.com/gin-gonic/gin"
	"context"
	"io/ioutil"
	"time"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-restaurent-management/database"
	"golang-restaurent-management/models"
	"go.mongodb.org/mongo-driver/bson"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		 ctx, cancel := context.WithTimeout(context.Background(),100*time.Second)
		result , err := userCollection.Find(context.TODO(),bson.M{})
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching all users"})

		}
         var allUsers[] bson.M
		 err = result.All(ctx,&allUsers)
		 if err!=nil{
			 log.Fatal(err)
		 } 
		 c.JSON(http.StatusOK , allUsers) 

	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		body := c.Request.Body
		user_id , _ := ioutil.ReadAll(body)
		var user models.User
		err:= userCollection.FindOne(ctx, bson.M{"User_id":user_id}).Decode(&user)
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while fetching user data"})
		}	
	}
}



func SignUp() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func LogIn() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}


func Hashpassword(password string) string{
	var a string
	a=""
     return a
}

func verifypassword(userpassword string , providedpassword string) (bool, string){
	var a string
	a=""
	var b bool
	b= false
	 return b,a
}
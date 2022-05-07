package controller

import(
	"github.com/gin-gonic/gin"
	"context"
	"io/ioutil"
	"time"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-restaurent-management/database"
	"golang-restaurent-management/models"
	"go.mongodb.org/mongo-driver/bson"
)
var foodCollection *mongo.Collection = database.OpenCollection(database.Client,"food")

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		body := c.Request.Body
		user_id , _ := ioutil.ReadAll(body)
		var user models.User
		err:= foodCollection.FindOne(ctx, bson.M{"User_id":user_id}).Decode(&user)
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

}

func verifypassword(userpassword string , providedpassword string) (bool, string){
	 
}
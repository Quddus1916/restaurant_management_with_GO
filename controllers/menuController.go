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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/go-playground/validator"
)

var menucollection *mongo.Collection = database.OpenCollection(database.Client ,"menu")
var validate *validator.Validate

func GetMenus() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx , cancel := context.WithTimeout(context.Background(),100*time.Second)

        result,err := menucollection.Find(context.TODO(),bson.M{})
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to fetch menus from database"})
		     return
		}

		var allMenus[] bson.M
		err = result.All(ctx,&allMenus)

		if err!= nil{
			log.Fatal(err)
		}

        c.JSON(http.StatusOK,allMenus)


	}
}

func GetMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx , cancel := context.WithTimeout(context.Background(),100*time.Second)
		body:= c.Request.Body
		menu_Id,_:= ioutil.ReadAll(body)
		var fetchedUser models.Menu
		err:= menucollection.FindOne(ctx,bson.M{"Menu_Id":menu_Id}).Decode(&fetchedUser)
		defer cancel()
		if err!= nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"fetching user failed"})
			return
		}

		c.JSON(http.StatusOK,fetchedUser)
		
	}
}

func CreateMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		var newMenu models.Menu
		ctx , cancel := context.WithTimeout(context.Background(),100*time.Second)
		err:= c.BindJSON(&newMenu)
		if err!= nil{
			c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
			return
		}
		validationerr:= validate.Struct(newMenu)
		if validationerr!= nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error": validationerr.Error()})
			return
		}
		
		newMenu.ID = primitive.NewObjectID()
		newMenu.Created_At,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		result,inserterr:= menucollection.InsertOne(ctx,newMenu)
		
		defer cancel()
        if inserterr!= nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error": "failed to insert"})
			return
		}

		c.JSON(http.StatusOK,result)

		
	}
}

func UpdateMenu() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}



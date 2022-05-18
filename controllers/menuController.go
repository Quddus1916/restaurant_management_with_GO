package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-restaurent-management/database"
	"golang-restaurent-management/models"
	"log"
	"net/http"
	"time"
)

var menucollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var validate *validator.Validate

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		result, err := menucollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error() + " failed to fetch menus from database"})
			return
		}

		var allMenus []bson.M
		err = result.All(ctx, &allMenus)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allMenus)

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
		menu_Id := c.Param("menu_id")
		var fetchedMenu models.Menu
		err := menucollection.FindOne(ctx, bson.M{"menu_id": menu_Id}).Decode(&fetchedMenu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fetching user failed"})
			return
		}

		c.JSON(http.StatusOK, fetchedMenu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newMenu models.Menu
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		err := c.BindJSON(&newMenu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// validationerr:= validate.Struct(newMenu)
		// if validationerr!= nil{
		// 	c.JSON(http.StatusInternalServerError,gin.H{"error": validationerr.Error()})
		// 	return
		// }

		newMenu.ID = primitive.NewObjectID()
		newMenu.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		result, inserterr := menucollection.InsertOne(ctx, newMenu)

		defer cancel()
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert"})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var updatedMenu models.Menu
		err := c.BindJSON(&updatedMenu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		filter := bson.M{"menu_Id": updatedMenu.Menu_Id}
		updatedMenu.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		menucollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updatedMenu},
			},
			&opt,
		)
		defer cancel()
		c.JSON(http.StatusOK, updatedMenu)

	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		menu_Id := c.Param("menu_id")

		dltuser, err := menucollection.DeleteOne(ctx, bson.M{"menu_id": menu_Id})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "fetching user failed"})
			return
		}

		c.JSON(http.StatusOK, dltuser)

	}
}

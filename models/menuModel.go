package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct{
	ID    			primitive.ObjectID 			`bson:"_id"`
	Name             string  					`json:"name" `
	Category		 string	                    `json:"category"`
	Created_At        time.Time                 `json:"created_at"`
	Menu_Id    		  string 			        `json:"menu_id"`
}
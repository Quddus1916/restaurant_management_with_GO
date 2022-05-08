package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct{
	ID    			primitive.ObjectID 			`bson:"_id"`
	Name             string  					`json:"name" validate:"required,min=3,max=100"`
	Category		 string	                    `json:"category" validate:"required,eq=MAJOR|eq=MINOR|eq=MEDIUM"`
	Created_At        time.Time                 `json:"created_at"`
	Menu_Id    		  string 			        `json:"menu_id"`
}
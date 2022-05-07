package models


import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID    			primitive.ObjectID 			`bson:"_id"`
	Name            string                      `json:"name" validate:"required,min=2,max=30"`
	Email           string                      `json:"email" validate:"required,min=2,max=30"`
	Phone           string                      `json:"phone" validate:"required,min=2,max=30"`
	Token           string                      `json:"token"`
	Refresh_token   string                      `json:"refresh_token"`
	Created_at      time.Time                   `json:"created_at"`
	User_id         string                      `json:"user_id"`
}
package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurent-management/controllers"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.signUp())
	incomingRoutes.POST("/users/login", controller.LogIn())
}
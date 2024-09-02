package routes

import (
	"blog-api/controllers"
	"blog-api/pkg/dto/user_dto"
	"blog-api/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userController *controllers.UserController){
	r.POST("/signup", middlewares.ValidateDTO(&user_dto.SignUpDTO{}), userController.CreateUser)
	r.POST("/login",middlewares.ValidateDTO(&user_dto.LoginDTO{}), userController.LoginUser)
	r.GET("/all", middlewares.JWTMiddleware(), userController.GetAllUsers)
}
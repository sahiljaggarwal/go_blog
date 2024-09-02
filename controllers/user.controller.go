package controllers

import (
	"blog-api/models"
	"blog-api/pkg/dto/user_dto"
	"blog-api/pkg/utils"
	"log"

	// "blog-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) CreateUser (c *gin.Context){
	var input user_dto.SignUpDTO

	dto, exists := c.Get("dto")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve DTO from context"})
		c.Abort()
		return
	}

	if dto, ok := dto.(*user_dto.SignUpDTO); ok {
		input = *dto
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid DTO type"})
		c.Abort()
		return
	}
	var existingUser models.User

	if err := uc.DB.Where("username = ? OR email = ?",input.Username, input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"User already exists"})
		return
	}

	user := models.User {
		Username: input.Username,
		Email: input.Email,
		Password: input.Password,
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"User created successfully", "data": user_dto.ToPublicResponse(&user)})
}

func (uc *UserController) GetAllUsers(c *gin.Context){
	var users []models.User

	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"users not found",
			"data":nil,
		})
		return
	}

	var userResponses []user_dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user_dto.ToPrivateResponse(&user))
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"users data",
		"data":  userResponses,
	})

}

func (uc *UserController) LoginUser (c *gin.Context){

	var input user_dto.LoginDTO
	dto, exists :=  c.Get("dto")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Failed to retrive DTO from context",
		})
		c.Abort()
		return
	}
	if dto, ok := dto.(*user_dto.LoginDTO); ok {
		input = *dto
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Invalid DTO type",
		})
		c.Abort()
		return
	}

	var user models.User

	if err := uc.DB.Where("email = ? ", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message":"user not found with this email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid Email or Password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Username)
	if err != nil {
		log.Println("token error ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Could not generate access token", 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Login successfully",
		"token": token,
	})
}
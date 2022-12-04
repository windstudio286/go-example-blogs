package controller

import (
	"blogs/api/service"
	"blogs/models"
	"blogs/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.UserRegister
	if err := ctx.ShouldBind(&user); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}

	hashPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create user")
		return
	}
	utils.SuccessJSON(ctx, http.StatusOK, "Successfully Created user")
}

func (u *UserController) LoginUser(ctx *gin.Context) {
	var user models.UserLogin
	var hmacSampleSecret []byte

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}

	dbUser, err := u.service.LoginUser(user)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		utils.ErrorJSON(ctx, http.StatusBadRequest, "Failed to get token")
		return
	}

	response := &utils.Response{
		Success: true,
		Message: "Token generated sucessfully",
		Data:    tokenString,
	}
	ctx.JSON(http.StatusOK, response)
}

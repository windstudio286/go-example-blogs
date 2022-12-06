package v1

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

// CreateUser godoc
// @Summary      APIv2 Create a user
// @Description  Create a user by providing email and password
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Param Body body models.UserRegister true "The body to create a user"
// @Success      200  {object}  utils.HTTPSucess
// @Failure      400  {object}  utils.HTTPError
// @Router       /auth/register [post]
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

// CreateUser godoc
// @Summary      APIv2 Login
// @Description  Create a user by providing email and password
// @Tags         Auth API
// @Accept       json
// @Produce      json
// @Param Body body models.UserLogin true "The body to create a user"
// @Success      200  {object}  utils.Response{data=string}
// @Failure      400  {object}  utils.HTTPError
// @Router       /auth/login [post]
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

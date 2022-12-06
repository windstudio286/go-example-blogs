package v2

import (
	v2Controller "blogs/api/controller/v2"
	"blogs/infrastructure"
)

type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller v2Controller.UserController
}

func NewUserRoute(
	controller v2Controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u UserRoute) Setup() {
	v1 := u.Handler.Gin.Group("/v2")
	{
		user := v1.Group("/auth")
		{
			user.POST("/register", u.Controller.CreateUser)
			user.POST("/login", u.Controller.LoginUser)
		}
	}
}

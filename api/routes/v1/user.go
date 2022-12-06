package v1

import (
	v1Controller "blogs/api/controller/v1"
	"blogs/infrastructure"
)

type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller v1Controller.UserController
}

func NewUserRoute(
	controller v1Controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u UserRoute) Setup() {
	v1 := u.Handler.Gin.Group("/v1")
	{
		user := v1.Group("/auth")
		{
			user.POST("/register", u.Controller.CreateUser)
			user.POST("/login", u.Controller.LoginUser)
		}
	}
}

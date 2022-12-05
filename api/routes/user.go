package routes

import (
	"blogs/api/controller"
	"blogs/infrastructure"
)

type UserRoute struct {
	Handler    infrastructure.GinRouter
	Controller controller.UserController
}

func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,
) UserRoute {
	return UserRoute{
		Handler:    handler,
		Controller: controller,
	}
}

func (u UserRoute) Setup() {
	v1 := u.Handler.Gin.Group("/api/v1")
	{
		user := v1.Group("/auth")
		{
			user.POST("/register", u.Controller.CreateUser)
			user.POST("/login", u.Controller.LoginUser)
		}
	}
}

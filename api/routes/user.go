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
	user := u.Handler.Gin.Group("/auth")
	{
		user.POST("/register", u.Controller.CreateUser)
		user.POST("/login", u.Controller.LoginUser)
	}
}

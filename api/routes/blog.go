package routes

import (
	"blogs/api/controller"
	"blogs/infrastructure"
)

type PostRoute struct {
	Controller controller.PostController
	Handle     infrastructure.GinRouter
}

func NewPostRouter(controller controller.PostController, handler infrastructure.GinRouter) PostRoute {
	return PostRoute{
		Controller: controller,
		Handle:     handler,
	}
}
func (p PostRoute) Setup() {
	post := p.Handle.Gin.Group("/posts")
	{
		post.GET("/", p.Controller.GetPosts)
		post.POST("/", p.Controller.AddPost)
		post.GET("/:id", p.Controller.GetPost)
		post.DELETE("/:id", p.Controller.DeletePost)
		post.PUT("/:id", p.Controller.UpdatePost)
	}
}

package routes

import (
	"blogs/api/controller"

	"blogs/infrastructure"
)

type PostRoute struct {
	Controller controller.PostController
	Handler    infrastructure.GinRouter
}

func NewPostRouter(controller controller.PostController, handler infrastructure.GinRouter) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}
func (p PostRoute) Setup() {
	v1 := p.Handler.Gin.Group("/api/v1")
	{
		post := v1.Group("/posts")
		{
			post.GET("/", p.Controller.GetPosts)
			post.POST("/", p.Controller.AddPost)
			post.GET("/:id", p.Controller.GetPost)
			post.DELETE("/:id", p.Controller.DeletePost)
			post.PUT("/:id", p.Controller.UpdatePost)
		}
	}
}

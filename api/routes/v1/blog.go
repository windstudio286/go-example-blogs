package v1

import (
	v1Controller "blogs/api/controller/v1"

	"blogs/infrastructure"
)

type PostRoute struct {
	Controller v1Controller.PostController
	Handler    infrastructure.GinRouter
}

func NewPostRouter(controller v1Controller.PostController, handler infrastructure.GinRouter) PostRoute {
	return PostRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// @title Blog API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email ttcong194@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func (p PostRoute) Setup() {
	v1 := p.Handler.Gin.Group("/v1")
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

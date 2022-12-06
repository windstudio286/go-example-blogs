package main

import (
	v1Controller "blogs/api/controller/v1"
	v2Controller "blogs/api/controller/v2"
	"blogs/api/repository"
	v1Routes "blogs/api/routes/v1"
	v2Routes "blogs/api/routes/v2"
	"blogs/api/service"
	_ "blogs/docs"
	"blogs/infrastructure"
	"blogs/models"
	"fmt"
	"unsafe"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()

	db := infrastructure.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	v1PostController := v1Controller.NewPostController(postService)

	v1PostRoute := v1Routes.NewPostRouter(v1PostController, router)
	v1PostRoute.Setup()

	v2PostController := v2Controller.NewPostController(postService)

	v2PostRoute := v2Routes.NewPostRouter(v2PostController, router)
	v2PostRoute.Setup()

	// add config for user
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserSerivce(userRepository)
	v1UserController := v1Controller.NewUserController(userService)
	v1UserRouter := v1Routes.NewUserRoute(v1UserController, router)
	v1UserRouter.Setup()

	v2UserController := v2Controller.NewUserController(userService)
	v2UserRouter := v2Routes.NewUserRoute(v2UserController, router)
	v2UserRouter.Setup()

	//Dòng này sẽ config và đồng bộ với DB từ model
	db.DB.AutoMigrate(&models.Post{}, &models.User{})
	router.Gin.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("v1")))
	router.Gin.GET("/swagger/v2/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("v2")))
	router.Gin.Run(":8000")

}

func main1() {
	m := make(map[int]int)
	addKey(m, 1, 1)
	addKey(m, 2, 2)
	addKey(m, 3, 3)
	for k, v := range m {
		fmt.Println(k, "value is", v)
	}
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8 (linux/amd64)
}

func fn(m map[int]int) {
	m = make(map[int]int)
}

func addKey(m map[int]int, key int, value int) {
	m[key] = value
}

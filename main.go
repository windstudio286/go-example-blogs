package main

import (
	"blogs/api/controller"
	"blogs/api/repository"
	"blogs/api/routes"
	"blogs/api/service"
	"blogs/infrastructure"
	"blogs/models"
	"fmt"
	"unsafe"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()

	db := infrastructure.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	postRoute := routes.NewPostRouter(postController, router)
	postRoute.Setup()

	// add config for user
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserSerivce(userRepository)
	userController := controller.NewUserController(userService)
	userRouter := routes.NewUserRoute(userController, router)
	userRouter.Setup()

	//Dòng này sẽ config và đồng bộ với DB từ model
	db.DB.AutoMigrate(&models.Post{}, &models.User{})

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

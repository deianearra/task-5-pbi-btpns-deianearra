// package router

// import (
//     "github.com/gin-gonic/gin"
//     "path_to/controllers" // Sesuaikan dengan path ke package controllers
// )

// func SetupRouter() *gin.Engine {
//     r := gin.Default()

//     userController := controllers.NewUserController() // Buat instance controller User
//     photoController := controllers.NewPhotoController() // Buat instance controller Photo

//     userRoutes := r.Group("/users")
//     {
//         userRoutes.POST("/register", userController.Register)
//         userRoutes.POST("/login", userController.Login)
//         userRoutes.PUT("/:userId", userController.UpdateUser)
//         userRoutes.DELETE("/:userId", userController.DeleteUser)
//     }

//     photoRoutes := r.Group("/photos")
//     {
//         photoRoutes.POST("/", photoController.CreatePhoto)
//         photoRoutes.GET("/", photoController.GetPhotos)
//         photoRoutes.PUT("/:photoId", photoController.UpdatePhoto)
//         photoRoutes.DELETE("/:photoId", photoController.DeletePhoto)
//     }

//     return r
// }

package handlers

import (
	"task-5-pbi-btpns-deianearra/internal/controllers"
	"task-5-pbi-btpns-deianearra/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	// User Endpoints
	user := r.Group("/users")
	{
		user.POST("/register", controllers.RegisterUser)
		user.POST("/login", controllers.LoginUser)
		user.PUT("/:userId", controllers.UpdateUser)
		user.DELETE("/:userId", controllers.DeleteUser)
	}

	// Photo Endpoints
	photos := r.Group("/photos")
	{
		photos.POST("/", controllers.CreatePhoto)
		photos.GET("/", controllers.GetAllPhotos)
		photos.PUT("/:photoId", controllers.UpdatePhoto)
		photos.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r
}

package routes

import (
	"ina-gin-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	// OAuth 2.0 task routes
	taskGroup := r.Group("/tasks")
	// taskGroup.Use(ginserver.HandleTokenVerify())
	{
		taskGroup.GET("/", controllers.GetTasks)
		taskGroup.POST("/", controllers.CreateTask)
		taskGroup.GET("/:id", controllers.GetTaskByID)
		taskGroup.PUT("/:id", controllers.UpdateTask)
		taskGroup.DELETE("/:id", controllers.DeleteTask)
	}

	return r
}

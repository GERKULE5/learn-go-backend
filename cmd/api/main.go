// main.go
package main

import (
	_ "crud-gin/docs"
	"crud-gin/internal/handler"
	"crud-gin/internal/repository"
	"crud-gin/internal/service"
	"crud-gin/pkg/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CRUD API with Gin
// @version         1.0
// @description     A simple CRUD API built with Gin framework in Go.
// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	fmt.Println("Hello world")

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}

	err = r.Run(":8000")

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}

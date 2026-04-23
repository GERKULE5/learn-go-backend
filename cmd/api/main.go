// main.go
package main

import (
	"crud-gin/internal/handler"
	"crud-gin/internal/repository"
	"crud-gin/internal/service"
	"crud-gin/pkg/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	err = database.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	users := r.Group("/users")
	{
		users.POST("", userHandler.Create)
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	err = r.Run(":8000")

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}

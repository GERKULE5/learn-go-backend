// main.go
package main

import (
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

	_, err = database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	err = r.Run(":8000")

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}

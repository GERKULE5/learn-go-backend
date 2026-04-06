package database

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	DBName string
	SSLMode string
}


func ConnectDB(*gorm.DB)
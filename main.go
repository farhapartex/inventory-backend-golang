package main

import (
	"github.com/goupp-backend/server"
	"github.com/goupp-backend/config"
	"github.com/goupp-backend/model"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	// "net/http"
)

func loadEnv(){
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	fmt.Println("Loaded .env file")
}

func loadDatabase() {
    config.Connect()
	fmt.Println("Connected to database")
    config.Database.AutoMigrate(&model.User{})

	fmt.Println("--- Database migrated ---")
}

func main(){
	loadEnv()
	loadDatabase()
	server.Init()
}
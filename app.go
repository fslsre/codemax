package main

import (
	"codemax/config/router"
	"codemax/middleware"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := router.NewRouter()
	router.Use(middleware.Middleware)
	log.Fatal(http.ListenAndServe(":8888", router))
}

package main

import (
	"gigmile/pkg/controller"
	"gigmile/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("couldn't load env vars: %v", err)
	}

	db := database.Init()
	dbInstance := database.NewInstance(db)

	Newhttp := controller.New(dbInstance)

	Newhttp.Routes(router)
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port += "8000"
	}
	log.Printf("Server listening on port %s\n", port)
	err := router.Run(port)
	if err != nil {
		log.Fatalf("server failed to listen on port %s", port)
	}
}

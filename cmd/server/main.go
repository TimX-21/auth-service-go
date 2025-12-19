package main

import (
	"log"

	"github.com/TimX-21/auth-service-go/internal/config"
	"github.com/TimX-21/auth-service-go/internal/middlewares"
	"github.com/TimX-21/auth-service-go/pkg"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.InitZapSugaredLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
		return
	}
	
	defer config.Log.Sync()

	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
		return
	}

	router := gin.New()

	router.Use(gin.Recovery())
	router.ContextWithFallback = true
	router.Use(middlewares.LoggerMiddleware())

	router.Run(":8080")

	_ = db.Close()
}

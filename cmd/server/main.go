package main

import (
	"log"

	"github.com/TimX-21/auth-service-go/pkg"
)

func main() {
	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	log.Println("DB connected successfully")

	_ = db.Close()
}

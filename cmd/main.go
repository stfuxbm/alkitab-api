package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stfuxbm/alkitab-api/config"
	"github.com/stfuxbm/alkitab-api/internal/database"
	"github.com/stfuxbm/alkitab-api/internal/routes"
)

func main() {
	// Load environment variable dari file .env
	config.LoadEnv()

	// Koneksi ke MongoDB
	database.MongoConnect()

	// Setup semua route
	routes.SetupRoutes() // 👉 disarankan pakai nama fungsi yang deskriptif

	// Jalankan server di port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

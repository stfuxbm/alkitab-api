package routes

import (
	"net/http"

	"github.com/stfuxbm/alkitab-api/internal/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/api/v1/testament/new", handlers.AddBibleVerse)
}

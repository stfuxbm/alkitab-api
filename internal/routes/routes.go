package routes

import (
	"net/http"

	"github.com/stfuxbm/alkitab-api/internal/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/api/v1/testament/new", handlers.AddBibleVerse)
	http.HandleFunc("/api/v1/testament/new/get-all-bible", handlers.GetAllBibleChapters)
}

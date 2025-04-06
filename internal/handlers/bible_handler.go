package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/stfuxbm/alkitab-api/internal/database"
	"github.com/stfuxbm/alkitab-api/internal/models"
)

// Fungsi pembantu
func EncoderJsonData(w http.ResponseWriter, statuscode int, res models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(res)
}

// Handler POST ayat Alkitab
func AddBibleVerse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		EncoderJsonData(w, http.StatusMethodNotAllowed, models.Response{
			Success: false,
			Message: models.MsgMethodNotAllowed,
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	var chapter models.BookChapter
	err := json.NewDecoder(r.Body).Decode(&chapter)
	if err != nil {
		EncoderJsonData(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: models.MsgInvalidJSON,
			Code:    http.StatusBadRequest,
		})
		return
	}

	chapter.CreatedAt = time.Now()
	chapter.UpdatedAt = time.Now()

	collection := database.DB.Collection("new-testament")
	_, err = collection.InsertOne(r.Context(), chapter)
	if err != nil {
		EncoderJsonData(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: models.MsgInternalServerError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	EncoderJsonData(w, http.StatusCreated, models.Response{
		Success: true,
		Message: models.MsgVerseAdded,
		Code:    http.StatusCreated,
	})
}

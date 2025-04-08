package handlers

import (
	"context"
	"net/http"

	"github.com/stfuxbm/alkitab-api/internal/database"
	"github.com/stfuxbm/alkitab-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllBibleChapters mengambil semua data pasal dari MongoDB
func GetAllBibleChapters(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		EncoderJsonData(w, http.StatusMethodNotAllowed, models.Response{
			Success: false,
			Message: models.MsgMethodNotAllowed,
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	collection := database.DB.Collection("new-testament")

	// Ambil semua dokumen
	cursor, err := collection.Find(context.Background(), bson.M{}, options.Find())
	if err != nil {
		EncoderJsonData(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal mengambil data dari database.",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	defer cursor.Close(context.Background())

	var chapters []models.BookChapter
	if err := cursor.All(context.Background(), &chapters); err != nil {
		EncoderJsonData(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal membaca hasil data.",
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Respons sukses
	EncoderJsonData(w, http.StatusOK, models.Response{
		Success: true,
		Data:    chapters,
		Message: "Data semua pasal berhasil diambil.",
		Code:    http.StatusOK,
	})
}

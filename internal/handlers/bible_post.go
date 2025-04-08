package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/stfuxbm/alkitab-api/internal/database"
	"github.com/stfuxbm/alkitab-api/internal/models"
)

// EncoderJsonData mengirimkan respons JSON
func EncoderJsonData(w http.ResponseWriter, statusCode int, res models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

// AddBibleVerse menambahkan satu pasal Alkitab ke MongoDB
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

	// Decode JSON body ke struct
	if err := json.NewDecoder(r.Body).Decode(&chapter); err != nil {
		EncoderJsonData(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: models.MsgInvalidJSON,
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Validasi field wajib
	if chapter.Title == "" || chapter.Chapter == "" {
		EncoderJsonData(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Field 'title' dan 'chapter' wajib diisi.",
			Code:    http.StatusBadRequest,
		})
		return
	}

	if len(chapter.Verses) == 0 {
		EncoderJsonData(w, http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Field 'verses' tidak boleh kosong.",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// Validasi metadata per bahasa
	metadatas := map[string]models.MetaDetail{
		"latin":     chapter.Metadata.Latin,
		"indonesia": chapter.Metadata.Indonesia,
		"toraja":    chapter.Metadata.Toraja,
	}

	for lang, meta := range metadatas {
		if meta.Language == "" || meta.Version == "" || meta.Publisher == "" {
			EncoderJsonData(w, http.StatusBadRequest, models.Response{
				Success: false,
				Message: "Metadata untuk bahasa '" + lang + "' tidak lengkap.",
				Code:    http.StatusBadRequest,
			})
			return
		}
	}

	// Ambil koleksi MongoDB
	collection := database.DB.Collection("new-testament")

	// Validasi duplikat berdasarkan title dan chapter
	filter := map[string]interface{}{
		"title":   chapter.Title,
		"chapter": chapter.Chapter,
	}
	count, err := collection.CountDocuments(r.Context(), filter)
	if err != nil {
		EncoderJsonData(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal mengecek data duplikat.",
			Code:    http.StatusInternalServerError,
		})
		return
	}
	if count > 0 {
		EncoderJsonData(w, http.StatusConflict, models.Response{
			Success: false,
			Message: "Data untuk kitab dan pasal ini sudah ada.",
			Code:    http.StatusConflict,
		})
		return
	}

	// Set waktu
	chapter.CreatedAt = time.Now()
	chapter.UpdatedAt = time.Now()

	// Simpan ke MongoDB
	_, err = collection.InsertOne(r.Context(), chapter)
	if err != nil {
		EncoderJsonData(w, http.StatusInternalServerError, models.Response{
			Success: false,
			Message: models.MsgInternalServerError,
			Code:    http.StatusInternalServerError,
		})
		return
	}

	// Respons sukses
	EncoderJsonData(w, http.StatusCreated, models.Response{
		Success: true,
		Message: models.MsgVerseAdded,
		Code:    http.StatusCreated,
	})
}

package models

// Struct standar respons API
type Response struct {
	Success bool        `json:"success"` // Status berhasil atau tidak
	Data    interface{} `json:"data"`    // Data dinamis (bisa struct, slice, string, dll)
	Message string      `json:"message"` // Pesan (sukses/gagal)
	Code    int         `json:"code"`    // Kode HTTP (200, 400, dll)
}

const (
	MsgMethodNotAllowed    = "Metode tidak diizinkan"
	MsgInvalidJSON         = "Format JSON tidak valid"
	MsgVerseAdded          = "Ayat berhasil ditambahkan"
	MsgVerseAlreadyExist   = "Ayat sudah ada dalam database"
	MsgVerseNotFound       = "Ayat tidak ditemukan"
	MsgVerseList           = "Daftar ayat berhasil diambil"
	MsgVerseDeleted        = "Ayat berhasil dihapus"
	MsgVerseUpdated        = "Ayat berhasil diperbarui"
	MsgInternalServerError = "Terjadi kesalahan pada server"
)

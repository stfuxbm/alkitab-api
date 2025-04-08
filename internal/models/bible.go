package models

import "time"

// VerseContent sekarang menyimpan beberapa bahasa sebagai string
// Misalnya: Latin, Indonesia, Toraja

type VerseContent struct {
	Latin     string `json:"latin"`
	Indonesia string `json:"indonesia"`
	Toraja    string `json:"toraja"`
}

// Struct untuk menyimpan 1 ayat lengkap
// Contoh: Matius 1:1
type Verse struct {
	VerseNumber int          `json:"verse_number"`
	Content     VerseContent `json:"content"`
}

// MetadataBible digunakan untuk menyimpan informasi sumber untuk seluruh pasal
// Ini membantu mendokumentasikan versi dan penerbit masing-masing bahasa
type MetadataBible struct {
	Latin     MetaDetail `json:"latin"`
	Indonesia MetaDetail `json:"indonesia"`
	Toraja    MetaDetail `json:"toraja"`
}

type MetaDetail struct {
	Language  string `json:"language"`
	Version   string `json:"version"`
	Publisher string `json:"publisher"`
}

// Struct untuk menyimpan satu pasal (chapter)
// Contoh: Matius 1
// Berisi kumpulan ayat dan metadata keseluruhan
type BookChapter struct {
	Title      string        `json:"title"`       // Nama kitab, misal "Matius"
	BookNumber string        `json:"book_number"` // Urutan kitab dalam Alkitab
	Chapter    string        `json:"chapter"`     // Nomor pasal
	Verses     []Verse       `json:"verses"`      // Kumpulan ayat
	Metadata   MetadataBible `json:"metadata"`    // Metadata untuk masing-masing bahasa
	CreatedAt  time.Time     `json:"created_at,omitempty"`
	UpdatedAt  time.Time     `json:"updated_at,omitempty"`
}

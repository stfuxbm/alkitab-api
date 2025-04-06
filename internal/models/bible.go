package models

import "time"

// Struct untuk menyimpan isi ayat dalam berbagai bahasa
type VerseContent struct {
	Latin     string `json:"Latin"`
	Indonesia string `json:"Indonesia"`
	Toraja    string `json:"Toraja"`
}

// Struct untuk menyimpan 1 ayat lengkap
type Verse struct {
	VerseNumber int          `json:"verse_number"`
	Content     VerseContent `json:"content"`
}

// Struct untuk menyimpan satu pasal (chapter)
type BookChapter struct {
	Title      string    `json:"title"`
	BookNumber string    `json:"book_number"`
	Chapter    string    `json:"chapter"`
	Verses     []Verse   `json:"verses"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

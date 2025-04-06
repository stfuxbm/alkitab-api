package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/stfuxbm/alkitab-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database // Variabel global untuk menyimpan koneksi database

func MongoConnect() {
	// Ambil URI dan nama database dari file .env melalui config package
	uri := config.GetMongoURI()
	dbName := config.GetMongoDBName()

	// Buat context dengan timeout 10 detik untuk koneksi ke MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Pastikan context dibatalkan setelah selesai (menghindari memory leak)

	// Inisialisasi koneksi ke MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		// Jika koneksi gagal, hentikan aplikasi dengan error message
		log.Fatal(" Gagal koneksi ke MongoDB:", err)
		return
	}

	// Cek apakah koneksi berhasil (ping ke server MongoDB)
	err = client.Ping(ctx, nil)
	if err != nil {
		// Jika tidak bisa ping, hentikan aplikasi juga
		log.Fatal("Tidak bisa terhubung ke MongoDB:", err)
	}

	// Simpan koneksi database ke variabel global DB
	DB = client.Database(dbName)

	// Tampilkan pesan sukses ke terminal
	fmt.Println("Berhasil terhubung ke MongoDB")
}

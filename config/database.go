package config

import (
	"log"

	// Sesuaikan "clothify-go" dengan nama module yang kamu buat saat 'go mod init'
	"clothify-go/models" 
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB adalah variabel global penampung koneksi agar bisa dipanggil dari file lain
var DB *gorm.DB

func ConnectDatabase() {
	// DSN (Data Source Name): Alamat dan kunci masuk ke PostgreSQL kamu.
	// Nanti di tahap lanjut, nilai ini sebaiknya diambil dari file .env
	dsn := "host=localhost user=postgres password=postgres dbname=clothify_db port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	
	// Membuka koneksi ke database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal akan menghentikan program seutuhnya jika database gagal konek
		log.Fatal("Gagal terhubung ke database! Error: ", err) 
	}

	log.Println("Berhasil terhubung ke database PostgreSQL!")

	// --- FASE MIGRASI OTOMATIS (Pengganti 'npx prisma db push') ---
	// GORM akan membaca struct yang kita buat dan menyulapnya jadi tabel beneran di database
	err = database.AutoMigrate(
		&models.User{}, 
		&models.Password{}, 
		&models.Product{},
	)
	
	if err != nil {
		log.Fatal("Gagal melakukan migrasi tabel! Error: ", err)
	}

	log.Println("Tabel database berhasil di-migrate!")

	// Simpan koneksi yang berhasil ke variabel global DB
	DB = database
}
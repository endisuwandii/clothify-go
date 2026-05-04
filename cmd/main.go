package main

import (
	"clothify-go/handlers"
	"fmt"

	"github.com/gin-gonic/gin"

	// Ganti "clothify-go" dengan nama module kamu di go.mod kalau berbeda
	"clothify-go/config"
)

func main() {
	// 1. Panggil fungsi untuk menyambungkan ke database & auto-migrate tabel
	config.ConnectDatabase()

	// 2. Siapkan router Gin (Framework Web-nya)
	r := gin.Default()

	// 3. Buat rute uji coba (seperti endpoint test di Hono)
	r.GET("/products", handlers.GetProducts)
	

	// 4. Nyalakan server di port 8080
	fmt.Println("🚀 Server berjalan di http://localhost:8080")
	r.Run(":8080")
}




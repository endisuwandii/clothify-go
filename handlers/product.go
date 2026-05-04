package handlers

import (
	"clothify-go/config"
	"clothify-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	// 1. Siapkan wadah (slice/array) untuk menampung daftar produk
	var products []models.Product

	// 2. Perintahkan GORM untuk mengambil SEMUA data dari tabel products
	// Query aslinya: SELECT * FROM products;
	result := config.DB.Find(&products)

	// 3. Cek apakah ada error saat mengambil data
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data produk"})
		return
	}

	// 4. Kirim data produk dalam format JSON ke Frontend
	c.JSON(http.StatusOK, products)
}
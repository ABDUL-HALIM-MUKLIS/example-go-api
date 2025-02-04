package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "net/http"
    "go-api/internal/example"
    "go-api/middleware"
    "go-api/database"
    "github.com/joho/godotenv"
)

func main() {
    // Muat variabel lingkungan dari file .env
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Inisialisasi koneksi ke database
    database.InitDB()

    r := gin.Default()

    // Terapkan middleware global untuk menangani kesalahan
    r.Use(middleware.ErrorHandler)

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Project API menggunakan GO dengan Gin Framework", 
        })
    })

    example.SetupRouter(r.Group("/api/example"))

    r.Run("0.0.0.0:8080")
}

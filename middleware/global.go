package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// ErrorHandler adalah middleware untuk menangani kesalahan
func ErrorHandler(c *gin.Context) {
    c.Next() // Lanjutkan ke handler berikutnya

    // Periksa apakah ada kesalahan yang disimpan dalam konteks
    if _, exists := c.Get("error"); exists {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  false,
            "message": "Terjadi kesalahan pada server, silakan coba beberapa saat lagi",
        })
    }
}
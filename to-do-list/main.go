package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

func generateQRCode(c *gin.Context) {
	text := "http://localhost:8080/message" // QR-код сканерлегенде осы сілтеме ашылады
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "QR код генерациясында қате"})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Writer.Write(png)
}

func showMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Мен сені жақсы көремін! ❤️"})
}

func main() {
	r := gin.Default()
	r.GET("/generate", generateQRCode)
	r.GET("/message", showMessage)

	r.Run(":8080")
}

// cmd/frontend/main.go
package main

import (
	"frontend/internal/frontend"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin-Router erstellen
	router := gin.Default()

	// Statische Dateien (z.B. HTML, CSS, JS) servieren
	router.Static("/static", "./static")

	// Endpunkt für die Startseite
	router.GET("/", frontend.HandleHomePage)

	// Starte den Server auf Port 8080
	router.Run(":8080")
}

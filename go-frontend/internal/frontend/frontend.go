// internal/frontend/frontend.go
package frontend

import (
	"github.com/gin-gonic/gin"
)

// MyData ist ein Beispiel für eine Struktur, die im HTML-Rendering verwendet wird.
type MyData struct {
	Title string
}

// HandleHomePage ist für den Endpunkt der Startseite verantwortlich.
func HandleHomePage(c *gin.Context) {
	data := &MyData{
		Title: "Meine Webseite",
	}

	c.HTML(200, "index.html", gin.H{
		"data": data,
	})
}

// internal/frontend/frontend.go
package frontend

import (
	"github.com/gin-gonic/gin"
)

// HandleHomePage ist für den Endpunkt der Startseite verantwortlich.
func HandleHomePage(c *gin.Context) {
	if c == nil {
		// Handle den Fall, wenn der Context nil ist
		// z.B.: Rückgabe einer Fehlermeldung oder Beenden der Funktion
		return
	}

	data := struct {
		Title string
	}{
		Title: "Meine Webseite",
	}

	c.HTML(200, "./static/index.html", gin.H{
		"data": data,
	})
}

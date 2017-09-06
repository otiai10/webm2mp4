package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
)

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w, true).JSON(http.StatusOK, map[string]interface{}{
		"name":    "webm2mp4",
		"version": version,
		"message": "Hello!",
	})
}

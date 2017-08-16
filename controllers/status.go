package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/webm2mp4/config"
)

// Status ...
func Status(w http.ResponseWriter, r *http.Request) {
	marmoset.Render(w, true).JSON(http.StatusOK, map[string]interface{}{
		"name":    config.AppName(),
		"version": config.Version(),
		"message": "Hello!",
	})
}

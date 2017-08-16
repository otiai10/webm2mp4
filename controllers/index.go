package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/webm2mp4/config"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w)

	render.HTML("index", marmoset.P{
		"AppName": config.AppName(),
	})
}

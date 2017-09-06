package controllers

import (
	"net/http"

	"github.com/otiai10/marmoset"
)

const version = "0.0.1"

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w)

	render.HTML("index", marmoset.P{
		"AppName": "webm2mp4",
	})
}

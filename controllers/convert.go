package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/otiai10/marmoset"
	"github.com/otiai10/webm2mp4/config"

	"github.com/otiai10/goavcodec/v0/goavcodec"
)

// Convert ...
func Convert(w http.ResponseWriter, r *http.Request) {

	render := marmoset.Render(w)

	f, h, err := r.FormFile("file")
	if err != nil {
		render.JSON(http.StatusBadRequest, marmoset.P{"message": err.Error()})
		return
	}

	source, err := ioutil.TempFile("", config.AppName()+"_"+h.Filename+"_")
	if err != nil {
		render.JSON(http.StatusBadRequest, marmoset.P{"message": err.Error()})
		return
	}
	defer func() {
		source.Close()
		os.Remove(source.Name())
	}()

	if _, err = io.Copy(source, f); err != nil {
		render.JSON(http.StatusBadRequest, marmoset.P{"message": err.Error()})
		return
	}

	destpath := source.Name() + ".mp4"

	client, err := goavcodec.NewClient()
	if err != nil {
		render.JSON(http.StatusInternalServerError, marmoset.P{"message": err.Error()})
		return
	}

	if err = client.Convert(source.Name(), destpath); err != nil {
		render.JSON(http.StatusInternalServerError, marmoset.P{"message": err.Error()})
		return
	}

	resultmp4, err := os.Open(destpath)
	if err != nil {
		render.JSON(http.StatusInternalServerError, marmoset.P{"message": err.Error()})
		return
	}
	defer func() {
		resultmp4.Close()
		os.Remove(resultmp4.Name())
	}()

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeFile(w, r, resultmp4.Name())
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/otiai10/marmoset"

	"github.com/otiai10/webm2mp4/config"
	"github.com/otiai10/webm2mp4/controllers"
)

var logger *log.Logger

const pkg = "github.com/otiai10/webm2mp4"

func main() {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		panic("Env var `GOPATH` is not set!")
	}
	pkgpath := filepath.Join(gopath, "src", pkg)

	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", config.AppName()), 0)

	marmoset.LoadViews(filepath.Join(pkgpath, "views"))

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)

	r.GET("/", controllers.Index)
	r.Static("/assets", filepath.Join(pkgpath, "assets"))

	logger.Printf("listening on port %s", config.Port())
	err := http.ListenAndServe(config.Port(), r)
	logger.Println(err)
}

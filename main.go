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

	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", config.AppName()), 0)

	marmoset.LoadViews(viewpath())

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)
	r.POST("/upload", controllers.Convert)

	r.GET("/", controllers.Index)
	r.Static("/assets", assetpath())

	logger.Printf("listening on port %s", config.Port())
	err := http.ListenAndServe(config.Port(), r)
	logger.Println(err)
}

func cwd() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return filepath.Join(gopath, "src", "github.com/otiai10/webm2mp4")
	}
	if cwd, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	} else {
		return cwd
	}
}

func viewpath() string {
	return filepath.Join(cwd(), "views")
}

func assetpath() string {
	return filepath.Join(cwd(), "assets")
}

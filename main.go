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

	if cwd, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	} else {
		marmoset.LoadViews(filepath.Join(cwd, "views"))
	}

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)

	r.GET("/", controllers.Index)
	r.Static("/assets", "./assets")

	logger.Printf("listening on port %s", config.Port())
	err := http.ListenAndServe(config.Port(), r)
	logger.Println(err)
}

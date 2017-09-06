package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/otiai10/marmoset"

	"github.com/otiai10/webm2mp4/config"
	"github.com/otiai10/webm2mp4/controllers"
)

var logger *log.Logger

func main() {

	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", config.AppName()), 0)

	marmoset.LoadViews("./app/views")

	r := marmoset.NewRouter()

	// API
	r.GET("/status", controllers.Status)
	r.POST("/upload", controllers.Convert)

	r.GET("/", controllers.Index)
	r.Static("/assets", "./app/assets")

	logger.Printf("listening on port %s", config.Port())
	err := http.ListenAndServe(config.Port(), r)
	logger.Println(err)
}

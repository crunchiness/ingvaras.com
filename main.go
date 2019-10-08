package main

import (
	"github.com/crunchiness/ingvaras.com/ingvaras"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/artwork/:artist/:album", ingvaras.ArtworkHandler)
	router.GET("/artwork/:artist/:album/:raw", ingvaras.ArtworkHandler)
	router.GET("/tts/:lang/*query", ingvaras.TtsHandler)
	router.NotFound = http.HandlerFunc(notFoundHandler)
	http.Handle("/", router)
	appengine.Main()
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, "static/index.htm")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/404.htm")
}

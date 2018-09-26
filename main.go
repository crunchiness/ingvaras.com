package main

import (
	"github.com/crunchiness/ingvaras.com/ingvaras"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/apis", apisHandler)
	router.GET("/artwork/:artist/:album", ingvaras.ArtworkHandler)
	router.GET("/artwork/:artist/:album/:raw", ingvaras.ArtworkHandler)
	router.GET("/tts/:lang/*query", ingvaras.TtsHandler)
	router.NotFound = http.HandlerFunc(notFoundHandler)
	http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, "static/index.htm")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/404.htm")
}

func apisHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, "static/apis.htm")
}

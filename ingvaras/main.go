package ingvaras

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	router := httprouter.New()
	router.GET("/", handler)
	router.GET("/apis", apisHandler)
	router.GET("/artwork/:artist/:album", artworkHandler)
	router.GET("/artwork/:artist/:album/:raw", artworkHandler)
	router.GET("/tts/:lang/*query", ttsHandler)
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

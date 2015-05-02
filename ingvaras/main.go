package ingvaras

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func init() {
    router := httprouter.New()
    router.GET("/", handler)
    router.GET("/artwork/:artist/:album", artworkHandler)
    router.GET("/artwork/:artist/:album/:redir", artworkHandler)
    router.GET("/tts/:lang/*query", ttsHandler)
    router.NotFound = http.HandlerFunc(notFoundHandler)
    http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    http.ServeFile(w, r, "static/index.html")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/404.html")
}

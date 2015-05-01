package ingvaras

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func init() {
    router := httprouter.New()
    router.GET("/", handler)
    router.GET("/artwork/:artist/:album", artworkHandler)
    router.GET("/artwork/:artist/:album/:redir", artworkHandler)
    router.GET("/tts/:lang/*query", ttsHandler)
    http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprint(w, "Hello, world!")
}

package ingvaras

import (
	"net/url"
    "fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
    "strings"
)

func makeUrl2(lang, query string) string {
    const UrlBase = "http://ingvaras.appspot.com/tts?tl=%s&q=%s"
	query = url.QueryEscape(query)
	link := fmt.Sprintf(UrlBase, lang, query)
	return link
}

func ttsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lang := ps.ByName("lang")
	query := strings.Replace(ps.ByName("query")[1:], "+", " ", -1)
	link := makeUrl2(lang, query)
    http.Redirect(w, r, link, 301)
}

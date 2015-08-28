package ingvaras

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func makeUrl(artist, album string) string {
	const UrlBase = "http://ws.audioscrobbler.com/2.0/?method=album.getinfo&format=json&api_key=%s&artist=%s&album=%s"
	artist = url.QueryEscape(artist)
	album = url.QueryEscape(album)
	link := fmt.Sprintf(UrlBase, lastFmApiKey, artist, album)
	return link
}

type AlbumJson struct {
	Album Album `json:"album"`
}

type Album struct {
	Name      string  `json:"name"`
	Artist    string  `json:"artist"`
	ImageList []Image `json:"image"`
}

type Image struct {
	Url  string `json:"#text"`
	Size string `json:"size"`
}

func artworkHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	context := appengine.NewContext(r)
	client := urlfetch.Client(context)
	artist := strings.Replace(ps.ByName("artist"), "+", " ", -1)
	album := strings.Replace(ps.ByName("album"), "+", " ", -1)
	redir := ps.ByName("raw") == "raw"
	link := makeUrl(artist, album)
	resp, err := client.Get(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var albumJson AlbumJson
	err = json.Unmarshal(body, &albumJson)
	bigImage := Image{}
	for _, element := range albumJson.Album.ImageList {
		if element.Size == "mega" {
			bigImage = element
		}
	}

	if redir {
		http.Redirect(w, r, bigImage.Url, 301)
	}
	template := `<html>
    <head>
        <title>%s - %s (size: %s)</title>
    </head>
    <body style="margin: 0; height: 100%%">
		<a href="%s"><img src="%s" style="max-height: 100%%; width: auto"/></a>
    </body>
    </html>`
	fmt.Fprintf(w, template, albumJson.Album.Artist, albumJson.Album.Name, bigImage.Size, bigImage.Url, bigImage.Url)
}

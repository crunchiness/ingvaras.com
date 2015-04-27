package albumcover

import (
    "fmt"
    "net/http"
    "net/url"
    "appengine"
    "appengine/urlfetch"
    "io/ioutil"
    "encoding/json"
    "strings"
    "github.com/julienschmidt/httprouter"
)



func makeUrl(artist, album string) string {
    const UrlBase = "http://ws.audioscrobbler.com/2.0/?method=album.getinfo&format=json&api_key=%s&artist=%s&album=%s"
    const ApiKey = "a1233d94d65c3e7622da67d9835ed173"
    artist = url.QueryEscape(strings.Replace(artist, "+", " ", -1))
    album = url.QueryEscape(strings.Replace(album, "+", " ", -1))
    link := fmt.Sprintf(UrlBase, ApiKey, artist, album)
    return link
}

func init() {
    router := httprouter.New()
    router.GET("/", handler)
    router.GET("/ac/:artist/:album", handler2)
    router.GET("/ac/:artist/:album/:redir", handler2)
    http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprint(w, "Hello, world!")
}

type AlbumJson struct {
    Album Album `json:"album"`
}

type Album struct {
    Name string `json:"name"`
    Artist  string `json:"artist"`
    ImageList []Image `json:"image"`
}

type Image struct {
    Url string `json:"#text"`
    Size string `json:"size"`
}

func handler2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    context := appengine.NewContext(r)
    client := urlfetch.Client(context)
    artist := ps.ByName("artist")
    album := ps.ByName("album")
    redir := ps.ByName("redir") == "redir"
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
    bigImage := albumJson.Album.ImageList[len(albumJson.Album.ImageList)-1]
    if redir {
        http.Redirect(w, r, bigImage.Url, 301)
    }
    template := `<html>
    <head>
        <title>%s - %s (size: %s)</title>
    </head>
    <body>
        <img src="%s">
    </body>
    </html>`
    fmt.Fprintf(w, template, albumJson.Album.Artist, albumJson.Album.Name, bigImage.Size, bigImage.Url)
}

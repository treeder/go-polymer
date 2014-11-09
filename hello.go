package hello

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func init() {

	// Had to do this because returns svg as text/xml when running on AppEngine: http://goo.gl/hwZSp2
	mime.AddExtensionType(".svg", "image/svg+xml")

	r := mux.NewRouter()
	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/posts", Posts)
	r.HandleFunc("/{rest:.*}", handler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("path:", r.URL.Path)
	http.ServeFile(w, r, "static/"+r.URL.Path)
}

type Post struct {
	Uid      int    `json:"uid"`
	Text     string `json:"text"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Favorite bool   `json:"favorite"`
}

func Posts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	// you'd use a real database here
	file, err := ioutil.ReadFile("posts.json")
	if err != nil {
		log.Println("Error reading posts.json:", err)
		panic(err)
	}
	fmt.Printf("file: %s\n", string(file))
	err = json.Unmarshal(file, &posts)
	if err != nil {
		log.Println("Error unmarshalling posts.json:", err)
	}

	bs, err := json.Marshal(posts)
	if err != nil {
		ReturnError(w, err)
		return
	}
	fmt.Fprint(w, string(bs))
}

func ReturnError(w http.ResponseWriter, err error) {
	fmt.Fprint(w, "{\"error\": \"%v\"}", err)
}

package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func middleware(w http.ResponseWriter, r *http.Request) (newW http.ResponseWriter, newR *http.Request) {
  // use this if your site doesnt need anything else than HTTP GET
  if r.Method != "GET" {
    http.Redirect(w, r, "https://twitter.com/YOUR_TWITTER_HANDLER", 302)
    return
  }
  //w.Header().Set("X-Hacker", "Hi.")
  return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
  r.HandleFunc("/p/{page}", pageHandler)
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", r)
}
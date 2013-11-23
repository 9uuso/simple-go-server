package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  middleware(w, r)
  parseLayout(w, r, "index", "Homepage")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
    middleware(w, r)
    vars := mux.Vars(r)
    filename := vars["page"]
    parseLayout(w, r, filename, strings.Title(filename))
}

func error500(w http.ResponseWriter, r *http.Request, err error) {
  http.Error(w, "Server went nuts. Please try again later.", 500)
  log.Fatal(err)
}
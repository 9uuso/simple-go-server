package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Created int    `json:"created"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := post{"my blog post", "<b>content</b> is html rendered", 1600000}
	middleware(w, r)
	HTML(w, r, "index", p)
}

func homeHandlerJSON(w http.ResponseWriter, r *http.Request) {
	p := post{"my blog post", "content is not html rendered", 1600000}
	middleware(w, r)
	JSON(w, r, p)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	p := post{"my blog post", "content is not html rendered", 1600000}
	middleware(w, r)
	vars := mux.Vars(r)
	filename := vars["page"]
	HTML(w, r, filename, p)
}

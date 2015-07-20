package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var helpers = template.FuncMap{
	// unescape unescapes HTML of s.
	// Used in templates such as "/post/display.tmpl"
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
}

var templates = *template.Must(template.New("layout").Funcs(helpers).ParseGlob("./templates/*.tmpl"))

type content struct {
	Page     template.HTML
	Data     post
	LoggedIn bool
}

func parseTemplate(filename string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	for _, t := range templates.Templates() {
		if t.Name() == filename+".tmpl" {
			err := t.Funcs(helpers).ExecuteTemplate(&buf, t.Name(), data)
			if err != nil {
				return nil, err
			}
			return buf.Bytes(), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("template %s not found", filename))
}

func JSON(w http.ResponseWriter, r *http.Request, data post) {
	d, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	i, err := w.Write(d)
	if err != nil {
		log.Println("JSON: writing of", r, "dropped after", i, "bytes with error", err)
	}
}

func HTML(w http.ResponseWriter, r *http.Request, filename string, data post) {
	var buf bytes.Buffer
	page, err := parseTemplate(filename, data)
	if err != nil {
		log.Println("parseTemplate:", err)
		if err.Error() == fmt.Sprintf("template %s not found", filename) {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(&buf, "layout", content{template.HTML(page), data, false})
	if err != nil {
		log.Println("layout execute:", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	i, err := w.Write(buf.Bytes())
	if err != nil {
		log.Println("HTML: writing of", r, "dropped after", i, "bytes with error", err)
	}
}

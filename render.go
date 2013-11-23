package main

import (
  "bytes"
  "html/template"
  "net/http"
)

type content struct {
  Title   string
  Header  template.HTML
  Content template.HTML
  Footer  template.HTML
}

func parseTemplate(file string, data interface{}) (out []byte, error error) {
  var buf bytes.Buffer
  t, err := template.ParseFiles(file)
  if err != nil {
    return nil, err
  }
  err = t.Execute(&buf, data)
  if err != nil {
    return nil, err
  }
  return buf.Bytes(), nil
}

func parseLayout(w http.ResponseWriter, r *http.Request, filename string, title string) {
  t, err := template.New("layout").ParseFiles("./templates/layout.html")
  header, err := parseTemplate("./templates/header.html", nil)
  page, err := parseTemplate("./templates/"+filename+".html", nil)
  if page == nil {
    page, err = parseTemplate("./templates/404.html", nil)
  }
  footer, err := parseTemplate("./templates/footer.html", nil)
  if err != nil {
    error500(w, r, err)
  }
  err = t.ExecuteTemplate(w, "layout", content{Title: title, Header: template.HTML(header), Content: template.HTML(page), Footer: template.HTML(footer)})
  if err != nil {
    error500(w, r, err)
  }
}
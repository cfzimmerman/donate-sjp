package main 

import (
	"fmt"
	"html/template"
	"net/http"
)

type Routes struct {}

func (r Routes) Version(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "version 0.0.1\n")
}

func (r Routes) Index(w http.ResponseWriter, req *http.Request) {
  page, err := template.ParseFiles("pages/index.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  page.Execute(w, nil)
}


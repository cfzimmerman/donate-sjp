package main

import (
	"net/http"
)

func main() {
  r := Routes{};
  http.HandleFunc("/", r.Index)
  http.HandleFunc("/version", r.Version)

  fs := http.FileServer(http.Dir("pages/build"))
  http.Handle("/build/", http.StripPrefix("/build/", fs))

  err := http.ListenAndServe(":8090", nil)
  if err != nil {
    panic(err)
  }
}

package main

import (
	"net/http"
)

func main() {
  r := Routes{};
  http.HandleFunc("/", r.Index)
  http.HandleFunc("/version", r.Version)

  err := http.ListenAndServe(":8090", nil)
  if err != nil {
    panic(err)
  }
}

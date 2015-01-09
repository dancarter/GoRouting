package main

import "net/http"
import "fmt"

func NewMux() *http.ServeMux {
  mux := http.NewServeMux()
  users := func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Users Index")
  }
  mux.HandleFunc("/users", users)
  return mux
}

func main() {
  http.ListenAndServe(":3000", NewMux())
}

package main

import ("fmt"
        "net/http")

func index_handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Whoa, Go is neat!")
}

func about_handler(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Great web design by Brett")
}

func main() {
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/about/", about_handler)
  http.ListenAndServe(":8000", nil)
}

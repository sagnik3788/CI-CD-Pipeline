package main

import (
 "fmt"
 "net/http"
)

func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "index.html")
 })

 port := 8000
 fmt.Printf("Server listening on :%d...\n", port)
 err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
 if err != nil {
  fmt.Println(err)
 }
}


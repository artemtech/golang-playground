package main

import (
  "os"
  "net/http"
  "fmt"
  "log"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
  current_hostname := os.Getenv("HOSTNAME")
  fmt.Fprintf(w, "Hello gopher! You're accessing from %s\n", current_hostname)
}


func main() {
  port := ":"+os.Getenv("PORT")
  if port == ":" {
    port = ":8000"
  }

  http.HandleFunc("/", mainHandler)
  fmt.Printf("Starting server at port %s\n", port)
  if err := http.ListenAndServe(port, nil); err != nil {
    log.Fatal(err)
  }
}

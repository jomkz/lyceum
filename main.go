package main

import "net/http"

func main() {
  println("listening for connections...")
  http.HandleFunc("/", index)
  http.ListenAndServe(":4778", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
  println("handle index...")
  w.Write([]byte("lyceum"))
}

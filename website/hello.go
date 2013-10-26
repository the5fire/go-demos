package main

import (
    "io"
    "log"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello, World!")
}

func main() {
    http.HandleFunc("/", indexHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}

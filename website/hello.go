package main

import (
    "log"
    "net/http"
    "html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w,  err.Error(), http.StatusInternalServerError)
        return
    }
    locals := make(map[string]interface{})
    locals["langName"] = "Golang"
    tmpl.Execute(w, locals)
    return
}

func main() {
    http.HandleFunc("/", indexHandler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}

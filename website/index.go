package main

import (
    "log"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "net/http"
    "html/template"
)

const (
    username = "root"
    userpwd = ""
    dbname = "go_demos"
)

// User is user
type User struct {
    ID string
    Name string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    db, err := getDB(username, userpwd, dbname)
    if err != nil {
        http.Error(w,  err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w,  err.Error(), http.StatusInternalServerError)
        return
    }
    rows, err := db.Query("SELECT id, name FROM t1")
    if err != nil {
        http.Error(w,  err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var id, name string
    locals := make(map[string]interface{})
    users := []User{}

    for rows.Next() {
        err = rows.Scan(&id, &name) //扫每一行，并把字段的值赋到id和newstitle
        if err == nil {
            fmt.Println(id, name)
            users = append(users, User{id, name})
        }
    }
    locals["users"] = users
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

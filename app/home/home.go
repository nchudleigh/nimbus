package home

import (
    "fmt"
    "net/http"
    "html/template"
    "log"
)

func Root(w http.ResponseWriter, r *http.Request) {
    t, err:=template.ParseFiles("templates/index.html")
    fmt.Println("Welcome home.")
    t.Execute(w,nil)
    if err != nil {
        log.Fatal("Template error",err)
    }
}
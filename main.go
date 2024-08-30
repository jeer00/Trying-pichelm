package main

import (
    "html/template"
    "net/http"
    "os"
)

type PageData struct {
    Title   string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("index.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        secretMessage := os.Getenv("SECRET_MESSAGE")
        data := PageData{
            Title:   "Hidden message",
            Message: secretMessage,
        }
        tmpl.Execute(w, data)
    })

    http.ListenAndServe(":8080", nil)
}
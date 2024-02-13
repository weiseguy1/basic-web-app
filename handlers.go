package main

import (
    "fmt"
    "html/template"
    "net/http"
)
func Home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
    err := parsedTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template:", err)
        return
    }
}

package main

import (
    "fmt"
    "net/http"
)
// To simplify code it is better to break it up into separate files.
// This code so far now has a main.go file and a handlers.go file to better organize it.
const portNumber = ":8080"

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

    _ = http.ListenAndServe(portNumber, nil)
}

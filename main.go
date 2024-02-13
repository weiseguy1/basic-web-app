package main

import (
    "fmt"
    "net/http"
)
/* 
To simplify code it is better to break it up into separate files.
For better readability, this code so far now has:
main.go - The main function of the web app
handlers.go - Place handers here for different web pages
render.go - Renders the webpages from the templates folder
*/
const portNumber = ":8080"

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

    _ = http.ListenAndServe(portNumber, nil)
}

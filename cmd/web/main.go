package main

/*
How this app compiles is it finds the import from handlers,
then from handlers.go pulls in the dependency of render.go
*/

import (
    "fmt"
    "github.com/weiseguy1/basic-web-app/pkg/handlers"
    "net/http"
)
/* 
To simplify code it is better to break it up into separate files.
For better readability, this code so far now has:
cmd
  - web
    - main.go - The main function of the web app
pkg
  - handers
    - handlers.go - Place handers here for different web pages
  - render
    - render.go - Renders the webpages from the templates folder
*/
const portNumber = ":8080"

func main() {
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

    _ = http.ListenAndServe(portNumber, nil)
}

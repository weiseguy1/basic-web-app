package main

import (
    "fmt"
    "net/http"
)

func main() {
    /* 
    HandleFunc takes, 
    - the path name (URL), 
    - an inline function that contains
      + the ResponseWriter which writes responses to the user
      + and a pointer to the Request function with takes in the request the user's browser sends
        processes the request to send back a response
    */
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        n, err := fmt.Fprintf(w, "Hello, World!") //Fprintf returns and int and an error type
        if err != nil { // if there is an error, print the error to console
            fmt.Println(err)
        }
        fmt.Println(fmt.Sprintf("Bytes written: %d", n))
    })

    _ = http.ListenAndServe(":8080", nil)
}

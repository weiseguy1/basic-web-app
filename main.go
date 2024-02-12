package main

import (
    "fmt"
    "errors"
    "net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
    sum := AddValues(2, 2)
    _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func AddValues(x, y int) int {
    return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
    f, err := divideValues(100.0, 0.0)
    if err != nil {
        fmt.Fprintf(w, "Cannot divide by 0")
        // This return halts the function, stopping it from running the print statement below
        return
    }

    fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
    if y <= 0 {
        err := errors.New("cannot divide by zero")
        return 0, err
    }
    result := x / y
    return result, nil
}

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about", About)
    http.HandleFunc("/divide", Divide)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

    _ = http.ListenAndServe(portNumber, nil)
}

package main

/*
How this app compiles is it finds the import from handlers,
then from handlers.go pulls in the dependency of render.go
*/

import (
    "fmt"
    "github.com/weiseguy1/basic-web-app/pkg/config"
    "github.com/weiseguy1/basic-web-app/pkg/handlers"
    "github.com/weiseguy1/basic-web-app/pkg/render"
    "github.com/alexedwards/scs/v2"
    "time"
    "log"
    "net/http"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

    // change to true when in production
    app.InProduction = false

    session = scs.New()
    session.Lifetime = 24 * time.Hour
    session.Cookie.Persist = true
    session.Cookie.SameSite = http.SameSiteLaxMode
    session.Cookie.Secure = app.InProduction

    app.Session = session

    tc, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("cannot create template cache")
    }

    app.TemplateCache = tc
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)
    
    render.NewTemplates(&app)

    fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

    srv := &http.Server {
        Addr: portNumber,
        Handler: routes(&app),
    }

    err = srv.ListenAndServe()
    log.Fatal(err)
}

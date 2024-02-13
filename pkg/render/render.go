package render 

import (
    "fmt"
    "log"
    "html/template"
    "net/http"
)

// RenderTemplate renders templates
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
    err := parsedTemplate.Execute(w, nil)
    if err != nil {
        fmt.Println("error parsing template:", err)
        return
    }
}

// this function stores the parsed template into a variable, which can then be used instead of reading from disk
var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmplKey string) {
    var tmpl *template.Template    
    var err error

    // check to see if we already have the template in our cache
    _, inMap := templateCache[tmplKey]
    if !inMap {
        // need to create the template
        log.Println("creating template and adding to cache")
        err = createTemplateCache(tmplKey)
        if err != nil {
            log.Println(err)
        }
    } else {
        // we have the template in the cache
        log.Println("using cached template")
    }

    tmpl = templateCache[tmplKey]

    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Println(err)
    }
}

func createTemplateCache(t string) error {
    templates := []string {
        fmt.Sprintf("./templates/%s", t),
        "./templates/base.layout.tmpl",
    }

    // parse the template
    tmpl, err := template.ParseFiles(templates...)
    if err != nil {
        return err
    }

    // add template to cache (map)
    templateCache[t] = tmpl

    return nil
}

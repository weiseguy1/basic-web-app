package handlers 

import (
    "github.com/weiseguy1/basic-web-app/pkg/config"
    "github.com/weiseguy1/basic-web-app/pkg/render"
    "github.com/weiseguy1/basic-web-app/pkg/models"
    "net/http"
)


// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
    App *config.AppConfig
}

// NewRepo creates a repository
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository {
        App: a,
    }
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
    Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, again."
    render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
        StringMap: stringMap,
    })
}

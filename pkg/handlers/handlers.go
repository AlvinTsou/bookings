package handlers

import (
	"fmt"
	"net/http"

	"github.com/AlvinTsou/WebDev/pkg/config"
	"github.com/AlvinTsou/WebDev/pkg/models"
	"github.com/AlvinTsou/WebDev/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is a function that handles the home route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println("RemoteIP: ", remoteIP)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is a function that handles the about route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic and send data to the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println("Get from Session RemoteIP: ", remoteIP)
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

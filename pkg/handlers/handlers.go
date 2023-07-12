package handlers

import (
	"net/http"

	"github.com/javiervallejoco/bookings/pkg/config"
	"github.com/javiervallejoco/bookings/pkg/models"
	"github.com/javiervallejoco/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", ip)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler function
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test": "Hello again!",
	}

	ip := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = ip

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

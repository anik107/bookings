package handlers

import (
	"net/http"

	"github.com/anik107/bookings/pkg/config"
	"github.com/anik107/bookings/pkg/models"
	"github.com/anik107/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, I'm from the string map"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: StringMap,
	})
}

package routes

import (
	"net/http"

	c "github.com/PyMarcus/go_site/models"
	rr "github.com/PyMarcus/go_site/renders"
)

func (repo *Repository) Index(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	rr.RenderTemplate(w, "index.html", &c.TemplateData{})
}

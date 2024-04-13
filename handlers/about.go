package routes

import (
	"net/http"

	c "github.com/PyMarcus/go_site/models"
	rr "github.com/PyMarcus/go_site/renders"
)

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	personalInfo := make(map[string]string)
	personalInfo["email"] = "Email@email.com.br"

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	personalInfo["remote_ip"] = remoteIP

	rr.RenderTemplate(w, "about.html", &c.TemplateData{
		StringMap: personalInfo,
	})
}

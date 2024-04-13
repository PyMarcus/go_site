package routes

import (
	"net/http"

	c "github.com/PyMarcus/go_site/models"
	rr "github.com/PyMarcus/go_site/renders"
)

func (repo *Repository) RoomSuite(w http.ResponseWriter, r *http.Request) {
	rr.RenderTemplate(w, "suite.html", &c.TemplateData{})
}

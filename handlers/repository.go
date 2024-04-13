package routes

import "github.com/PyMarcus/go_site/config"

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func CreateNewRepository(config *config.AppConfig) *Repository {
	return &Repository{
		App: config,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

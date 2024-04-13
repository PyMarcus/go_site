package renders

import (
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/PyMarcus/go_site/models"
)

var (
	cache = make(map[string]*template.Template)
	mutex sync.Mutex
)

func RenderTemplate(w http.ResponseWriter, temp string, data *models.TemplateData) {
	mutex.Lock()
	defer mutex.Unlock()

	templ, ok := cache[temp]

	if !ok {
		parser, err := template.ParseFiles("./templates/"+temp, "./templates/base.html")
		if err != nil {
			log.Fatal("Fail to find template " + temp)
		}
		templ = parser
		cache[temp] = parser
	}

	err := templ.Execute(w, data)
	if err != nil {
		log.Panic("Error to execute template " + temp + " " + err.Error())
	}
}

package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ishanshre/Booking-App/pkg/config"
	"github.com/ishanshre/Booking-App/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// this function returns the defualt data to every templates
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// if UseCache is true then redner template from cache
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get a request template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Println("Could not get the template from the template cache")
	}

	// create a new buffer to store templates and data to pass to template
	buff := new(bytes.Buffer)

	// add default data to all templates
	td = AddDefaultData(td)

	// applied the parsed templates and data to the buffer
	if err := t.Execute(buff, td); err != nil {
		log.Println(err)
	}

	// render the template using buffer.WriteTo
	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// path patterns for layout and pages templates
	pathLayoutPattern := filepath.Join("templates", "layout", "*.layout.tmpl")
	pathPagePattern := filepath.Join("templates", "pages", "*.page.tmpl")
	myCache := map[string]*template.Template{}

	// filepath.Glob return a slice with all *.page.tmpl in template/pages folder
	pages, err := filepath.Glob(pathPagePattern)
	if err != nil {
		return myCache, err
	}
	// loop through the pages and add base templates
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, fmt.Errorf("error in parsing template: %s", err)
		}

		// find the base layout file and
		matches, err := filepath.Glob(pathLayoutPattern)
		if err != nil {
			return myCache, err
		}
		// if found
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(pathLayoutPattern)
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

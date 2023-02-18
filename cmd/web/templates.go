package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/lsjoeberg/snippetbox/internal/models"
)

// templateData acts as a holding structure for any dynamic data that
// we need to pass to HTML templates.
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{CurrentYear: time.Now().Year()}
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	cache := make(map[string]*template.Template, len(pages))

	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template into a template set.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse partials templates.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

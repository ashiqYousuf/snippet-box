package main

import (
	"path/filepath"
	"text/template"

	"github.com/ashiqYousuf/snippet-box/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// ?Use the filepath.Glob() function to get a slice of all filepaths that
	// ?match the pattern "./ui/html/pages/*.tmpl".

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// ?Extract the file name (like 'home.tmpl') from the full filepath
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// ?Call ParseGlob() *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// ?Call ParseFiles() *on this template set* to add the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

package main

/*
Alex Edwards : Let's Go

Chapter 5 Dynamic HTML templates

Section 5.5 Common Dynamic Data
*/

import (
	"html/template"
	"path/filepath"

	"alexedwards/5_5/snippetbox/internal/models"
)

// templateData encapsulates dynamic HTML template data.
type templateData struct {
	CurrentYear int // Common dynamic data
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	// A slice of all matching HTML application filepaths.
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one-by-one.
	for _, page := range pages {

		name := filepath.Base(page)

		// Parse the base template file into a template set.
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call `ParseGlob()`` *on this template set* to add any partials.
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Add `page` to template-set.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add template-set to cache map, key = page name.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}

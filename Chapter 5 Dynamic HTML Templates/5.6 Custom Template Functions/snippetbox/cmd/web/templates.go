package main

/*
Alex Edwards : Let's Go

Chapter 5 Dynamic HTML templates

Section 5.6 Custom Template Functions
*/

import (
	"html/template"
	"path/filepath"
	"time"

	"alexedwards/5_6/snippetbox/internal/models"
)

// templateData encapsulates dynamic HTML template data.
type templateData struct {
	CurrentYear int // Common dynamic data.
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// Create a humanDate function which returns a nicely formatted string
// representation of a time.Time object.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup between the names of our
// custom template functions and the functions themselves.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache...
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

		// The template.FuncMap must be registered with the template set before you
		// call the ParseFiles() method. This means we have to use template.New() to
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
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

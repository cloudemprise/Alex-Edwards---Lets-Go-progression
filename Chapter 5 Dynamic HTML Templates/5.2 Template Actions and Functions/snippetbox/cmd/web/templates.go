package main

/*
Alex Edwards : Let's Go

Chapter 5 Dynamic HTML templates

Section 5.2 Template Actions and Functions
*/

import (
	"alexedwards/5_2/snippetbox/internal/models"
)

// Include a Snippets field in the templateData struct.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

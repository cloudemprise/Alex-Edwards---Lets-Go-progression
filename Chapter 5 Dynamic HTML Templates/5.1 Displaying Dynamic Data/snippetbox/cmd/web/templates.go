package main

/*
Alex Edwards : Let's Go

Chapter 5 Dynamic HTML templates

Section 5.1 Displaying Dynamic Data
*/

import (
	"alexedwards/5_1/snippetbox/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the build progresses.
type templateData struct {
	Snippet *models.Snippet
}

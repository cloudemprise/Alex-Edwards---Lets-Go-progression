package models

/*
Alex Edwards : Let's Go

Chapter 4 Database-driven responses

Section 4.7 Executing SQL statements
*/

import (
	"errors"
)

// Ensure no direct dependency between underlying datastore reliant on
// specific datastore errors.
var ErrNoRecord = errors.New("models: no matching record found")

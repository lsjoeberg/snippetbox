package main

import "github.com/lsjoeberg/snippetbox/internal/models"

// templateData acts as a holding structure for any dynamic data that
// we need to pass to HTML templates.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
